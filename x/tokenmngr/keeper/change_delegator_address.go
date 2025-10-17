package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

// This file implements delegator address migration following Cosmos SDK patterns for reward handling.
//
// The key insight from Cosmos SDK's Undelegate function is that it uses distribution hooks:
// 1. BeforeDelegationSharesModified - withdraws all existing rewards to the delegator
// 2. AfterDelegationModified - initializes new distribution tracking
//
// Our implementation mimics this pattern:
// 1. handleDistributionRewards: Withdraws all rewards (like BeforeDelegationSharesModified)
// 2. changeDelegations: Migrates delegations and initializes new tracking (like AfterDelegationModified)
// 3. changeUnbondingDelegations: Migrates unbonding delegations

func (k Keeper) ChangeDelegatorAddress(goCtx context.Context, oldAddress, newAddress sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Validate addresses
	if oldAddress.Empty() || newAddress.Empty() {
		return errorsmod.Wrap(stakingtypes.ErrEmptyDelegatorAddr, "addresses cannot be empty")
	}

	if oldAddress.Equals(newAddress) {
		return errorsmod.Wrap(stakingtypes.ErrBadDelegatorAddr, "old and new addresses are the same")
	}

	// Check if new address already has delegations
	existingDelegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, newAddress, 1)
	if err != nil {
		return errorsmod.Wrap(err, "failed to check existing delegations for new address")
	}
	if len(existingDelegations) > 0 {
		// If new address has delegations, we need to fix its distribution state if broken
		k.Logger().Warn("Target address already has delegations, will attempt to fix distribution state",
			"target_address", newAddress.String(),
			"delegation_count", len(existingDelegations))

		// Try to fix the distribution state for existing delegations
		if err := k.fixExistingDelegationDistribution(ctx, newAddress); err != nil {
			return errorsmod.Wrap(err, "failed to fix existing delegation distribution state")
		}

		// Skip the migration since delegations already exist
		k.Logger().Info("Distribution state fixed for existing delegations", "address", newAddress.String())
		return nil
	}

	// 1. Handle Distribution Rewards (withdraw old rewards first)
	if err := k.handleDistributionRewards(ctx, oldAddress, newAddress); err != nil {
		return errorsmod.Wrap(err, "failed to handle distribution rewards")
	}

	// 2. Handle Delegations
	if err := k.changeDelegations(ctx, oldAddress, newAddress); err != nil {
		return errorsmod.Wrap(err, "failed to change delegations")
	}

	// 3. Handle Unbonding Delegations
	if err := k.changeUnbondingDelegations(ctx, oldAddress, newAddress); err != nil {
		return errorsmod.Wrap(err, "failed to change unbonding delegations")
	}

	return nil
}

// changeDelegations handles changing delegator address for all delegations
// This follows the Cosmos SDK pattern for delegation modifications:
// 1. Process rewards before modifying (done in handleDistributionRewards)
// 2. Modify the delegation
// 3. Initialize new distribution tracking
func (k Keeper) changeDelegations(ctx sdk.Context, oldAddress, newAddress sdk.AccAddress) error {
	// Get all delegations for the old address
	delegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, oldAddress, 65535) // max uint16
	if err != nil {
		return err
	}

	// Process each delegation
	for _, delegation := range delegations {
		valAddr, err := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		if err != nil {
			return fmt.Errorf("failed to parse validator address %s: %w", delegation.ValidatorAddress, err)
		}

		// Store original shares for proper distribution initialization
		originalShares := delegation.Shares

		// Remove the old delegation
		if err := k.stakingKeeper.RemoveDelegation(ctx, delegation); err != nil {
			return fmt.Errorf("failed to remove old delegation: %w", err)
		}

		// Convert new address to string format
		newAddressStr, err := k.accountKeeper.AddressCodec().BytesToString(newAddress)
		if err != nil {
			return fmt.Errorf("failed to convert new address to string: %w", err)
		}

		// Create new delegation with updated address
		newDelegation := stakingtypes.NewDelegation(
			newAddressStr,
			delegation.ValidatorAddress,
			originalShares,
		)

		// Set the new delegation
		if err := k.stakingKeeper.SetDelegation(ctx, newDelegation); err != nil {
			return fmt.Errorf("failed to set new delegation: %w", err)
		}

		// Initialize distribution tracking for the new delegation
		// This follows the same pattern as Cosmos SDK's AfterDelegationModified hook

		// Get the validator to calculate proper stake (tokens from shares)
		validator, err := k.stakingKeeper.GetValidator(ctx, valAddr)
		if err != nil {
			return fmt.Errorf("failed to get validator for distribution initialization: %w", err)
		}

		// Calculate delegation stake in tokens (not shares)
		// This is the same calculation used in Cosmos SDK's initializeDelegation
		stake := validator.TokensFromSharesTruncated(originalShares)

		// Get the current validator period to properly initialize the delegation
		// This follows the exact pattern from Cosmos SDK's initializeDelegation
		valCurrentRewards, err := k.distributionKeeper.GetValidatorCurrentRewards(ctx, valAddr)
		if err != nil {
			return fmt.Errorf("failed to get validator current rewards: %w", err)
		}

		// For migrated delegations, we should use the current period as starting point
		// This is different from normal delegation creation which uses previousPeriod
		// Since we're migrating an existing delegation that already accumulated rewards
		currentPeriod := valCurrentRewards.Period

		k.Logger().Info("Using current period for migrated delegation initialization",
			"validator_current_period", valCurrentRewards.Period,
			"delegator_starting_period", currentPeriod)

		// For migrated delegations, we don't need to increment reference count for historical rewards
		// since we're using the current period, not a historical one
		// Just create the starting info with current period
		startingInfo := distrtypes.NewDelegatorStartingInfo(currentPeriod, stake, uint64(ctx.BlockHeight()))
		err = k.distributionKeeper.SetDelegatorStartingInfo(ctx, valAddr, newAddress, startingInfo)
		if err != nil {
			return fmt.Errorf("failed to set delegator starting info for new delegation: %w", err)
		}

		k.Logger().Info("Successfully migrated delegation",
			"old_delegator", oldAddress.String(),
			"new_delegator", newAddress.String(),
			"validator", delegation.ValidatorAddress,
			"shares", originalShares.String(),
			"stake_tokens", stake.String(),
			"validator_current_period", valCurrentRewards.Period,
			"delegator_starting_period", currentPeriod,
		)
	}

	return nil
}

// changeUnbondingDelegations handles changing delegator address for all unbonding delegations
func (k Keeper) changeUnbondingDelegations(ctx sdk.Context, oldAddress, newAddress sdk.AccAddress) error {
	// Get all unbonding delegations for the old address
	unbondingDelegations, err := k.stakingKeeper.GetUnbondingDelegations(ctx, oldAddress, 65535) // max uint16
	if err != nil {
		return err
	}

	// Process each unbonding delegation
	for _, ubd := range unbondingDelegations {
		// Remove the old unbonding delegation
		if err := k.stakingKeeper.RemoveUnbondingDelegation(ctx, ubd); err != nil {
			return fmt.Errorf("failed to remove old unbonding delegation: %w", err)
		}

		// Convert new address to string format
		newAddressStr, err := k.accountKeeper.AddressCodec().BytesToString(newAddress)
		if err != nil {
			return fmt.Errorf("failed to convert new address to string: %w", err)
		}

		// Create new unbonding delegation with updated address
		newUbd := stakingtypes.UnbondingDelegation{
			DelegatorAddress: newAddressStr,
			ValidatorAddress: ubd.ValidatorAddress,
			Entries:          ubd.Entries,
		}

		// Set the new unbonding delegation
		if err := k.stakingKeeper.SetUnbondingDelegation(ctx, newUbd); err != nil {
			return fmt.Errorf("failed to set new unbonding delegation: %w", err)
		}

		// Re-insert into queue if entries exist
		for _, entry := range ubd.Entries {
			if err := k.stakingKeeper.InsertUBDQueue(ctx, newUbd, entry.CompletionTime); err != nil {
				return fmt.Errorf("failed to insert unbonding delegation into queue: %w", err)
			}
		}
	}

	return nil
}

// handleDistributionRewards handles migrating reward distribution data
// This follows the same pattern as Cosmos SDK's Undelegate function:
// 1. Withdraw all existing rewards (same as BeforeDelegationSharesModified hook)
// 2. Clean up old distribution state
func (k Keeper) handleDistributionRewards(ctx sdk.Context, oldAddress, newAddress sdk.AccAddress) error {
	// Get all delegations for the old address to process rewards
	delegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, oldAddress, 65535)
	if err != nil {
		return fmt.Errorf("failed to get delegations for reward processing: %w", err)
	}

	// Process each delegation's rewards - withdraw them before migration
	// This mimics what BeforeDelegationSharesModified hook does in Cosmos SDK
	for _, delegation := range delegations {
		valAddr, err := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		if err != nil {
			return fmt.Errorf("failed to parse validator address %s: %w", delegation.ValidatorAddress, err)
		}

		// Withdraw any existing rewards to the old address
		// This is what BeforeDelegationSharesModified hook does in Cosmos SDK undelegate
		_, err = k.distributionKeeper.WithdrawDelegationRewards(ctx, oldAddress, valAddr)
		if err != nil {
			// Log the error but don't fail the migration - rewards might not exist yet
			// This matches the behavior in Cosmos SDK hooks
			k.Logger().Info("No rewards to withdraw during migration",
				"delegator", oldAddress.String(),
				"validator", delegation.ValidatorAddress,
				"error", err.Error())
		}

		// Clean up old delegator starting info to ensure clean state
		// This prepares for the new delegation initialization
		err = k.distributionKeeper.DeleteDelegatorStartingInfo(ctx, valAddr, oldAddress)
		if err != nil {
			// Log but continue - starting info might not exist
			k.Logger().Info("No starting info to delete during migration",
				"delegator", oldAddress.String(),
				"validator", delegation.ValidatorAddress,
				"error", err.Error())
		}
	}

	return nil
}

// fixExistingDelegationDistribution fixes broken distribution state for existing delegations
// This handles cases where delegations exist but distribution state is corrupted
func (k Keeper) fixExistingDelegationDistribution(ctx sdk.Context, delegatorAddr sdk.AccAddress) error {
	// Get all delegations for the address
	delegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, delegatorAddr, 65535)
	if err != nil {
		return fmt.Errorf("failed to get delegations for distribution fix: %w", err)
	}

	k.Logger().Info("Fixing distribution state for existing delegations",
		"delegator", delegatorAddr.String(),
		"delegation_count", len(delegations))

	// Process each delegation to fix its distribution state
	for _, delegation := range delegations {
		valAddr, err := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		if err != nil {
			return fmt.Errorf("failed to parse validator address %s: %w", delegation.ValidatorAddress, err)
		}

		// Check if delegator starting info exists and is valid
		startingInfo, err := k.distributionKeeper.GetDelegatorStartingInfo(ctx, valAddr, delegatorAddr)
		needsRecreation := false

		if err != nil {
			// Starting info doesn't exist, needs creation
			k.Logger().Info("Delegator starting info not found, needs creation",
				"delegator", delegatorAddr.String(),
				"validator", delegation.ValidatorAddress,
				"error", err.Error())
			needsRecreation = true
		} else if startingInfo.Stake.IsNil() || startingInfo.Stake.IsZero() {
			// Starting info exists but stake is nil or zero (corrupted), needs recreation
			k.Logger().Info("Delegator starting info corrupted (nil/zero stake), needs recreation",
				"delegator", delegatorAddr.String(),
				"validator", delegation.ValidatorAddress,
				"current_stake", startingInfo.Stake.String(),
				"current_period", startingInfo.PreviousPeriod,
				"current_height", startingInfo.Height)
			needsRecreation = true
		} else {
			k.Logger().Info("Distribution state is valid",
				"delegator", delegatorAddr.String(),
				"validator", delegation.ValidatorAddress,
				"stake", startingInfo.Stake.String(),
				"period", startingInfo.PreviousPeriod)
		}

		if needsRecreation {
			k.Logger().Info("Recreating delegator starting info",
				"delegator", delegatorAddr.String(),
				"validator", delegation.ValidatorAddress)

			// Delete existing corrupted starting info if it exists
			err = k.distributionKeeper.DeleteDelegatorStartingInfo(ctx, valAddr, delegatorAddr)
			if err != nil {
				k.Logger().Info("No existing starting info to delete (expected for new delegations)",
					"delegator", delegatorAddr.String(),
					"validator", delegation.ValidatorAddress)
			}

			// Get the validator to calculate proper stake
			validator, err := k.stakingKeeper.GetValidator(ctx, valAddr)
			if err != nil {
				return fmt.Errorf("failed to get validator for distribution fix: %w", err)
			}

			// Calculate delegation stake in tokens
			stake := validator.TokensFromSharesTruncated(delegation.Shares)

			// Get the current validator period
			valCurrentRewards, err := k.distributionKeeper.GetValidatorCurrentRewards(ctx, valAddr)
			if err != nil {
				return fmt.Errorf("failed to get validator current rewards for fix: %w", err)
			}

			// For fixing broken delegations, use the current period as starting point
			// This ensures proper alignment with the validator's current state
			currentPeriod := valCurrentRewards.Period

			// Create new starting info with current period
			startingInfo := distrtypes.NewDelegatorStartingInfo(currentPeriod, stake, uint64(ctx.BlockHeight()))
			err = k.distributionKeeper.SetDelegatorStartingInfo(ctx, valAddr, delegatorAddr, startingInfo)
			if err != nil {
				return fmt.Errorf("failed to set delegator starting info for fix: %w", err)
			}

			k.Logger().Info("Successfully fixed distribution state",
				"delegator", delegatorAddr.String(),
				"validator", delegation.ValidatorAddress,
				"shares", delegation.Shares.String(),
				"stake_tokens", stake.String(),
				"starting_period", currentPeriod)
		}
	}

	return nil
}

// GetDelegatorStakingInfo returns comprehensive staking information for a delegator
// This is useful for verifying the state before and after address changes
func (k Keeper) GetDelegatorStakingInfo(goCtx context.Context, delegatorAddr sdk.AccAddress) (*types.DelegatorStakingInfo, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	info := &types.DelegatorStakingInfo{
		DelegatorAddress: delegatorAddr.String(),
	}

	// Get delegations
	delegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, delegatorAddr, 65535)
	if err != nil {
		return nil, fmt.Errorf("failed to get delegations: %w", err)
	}
	info.Delegations = delegations

	// Get unbonding delegations
	unbondingDelegations, err := k.stakingKeeper.GetUnbondingDelegations(ctx, delegatorAddr, 65535)
	if err != nil {
		return nil, fmt.Errorf("failed to get unbonding delegations: %w", err)
	}
	info.UnbondingDelegations = unbondingDelegations

	return info, nil
}
