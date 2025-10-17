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

// This file implements delegator address migration following the MigrateDelegatorWithRewards pattern.
//
// Key features:
// 1. Simple Reward Migration: Withdraws all pending rewards and transfers them as tokens
// 2. Active Delegations Only: Only migrates active delegations, ignoring unbonding delegations
// 3. Automatic Reward Tracking: New delegations automatically start reward tracking
// 4. Validation: Ensures addresses are different and new address is clean
// 5. Error Handling: Comprehensive error handling with specific error types
//
// The implementation focuses on two main operations:
// 1. migrateRewardTrackingSimple: Withdraws and transfers all rewards
// 2. migrateActiveDelegations: Migrates delegations preserving shares exactly

// ChangeDelegatorAddress migrates active delegations and rewards from old delegator to new delegator
// This function only handles active delegations and reward tracking, ignoring unbonding delegations
// It follows the pattern from MigrateDelegatorWithRewards for proper reward handling
func (k Keeper) ChangeDelegatorAddress(goCtx context.Context, oldAddress, newAddress sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Validate addresses
	if oldAddress.Empty() || newAddress.Empty() {
		return errorsmod.Wrap(stakingtypes.ErrEmptyDelegatorAddr, "addresses cannot be empty")
	}

	if oldAddress.Equals(newAddress) {
		return types.ErrSameDelegatorAddress
	}

	// Check if new address already has delegations (optional validation)
	existingDelegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, newAddress, 1)
	if err != nil {
		return errorsmod.Wrap(err, "failed to check existing delegations for new address")
	}
	if len(existingDelegations) > 0 {
		return types.ErrNewDelegatorAlreadyExists
	}

	// 2. FIRST: Handle rewards before touching delegations
	if err := k.migrateRewardTrackingSimple(ctx, oldAddress, newAddress); err != nil {
		return errorsmod.Wrap(err, "failed to migrate reward tracking")
	}

	// 3. Then migrate active delegations (this will create new reward tracking)
	if err := k.migrateActiveDelegations(ctx, oldAddress, newAddress); err != nil {
		return errorsmod.Wrap(err, "failed to migrate active delegations")
	}

	return nil
}

// migrateActiveDelegations migrates all active delegations from old to new delegator
// This follows the pattern from MigrateDelegatorWithRewards for proper delegation migration
func (k Keeper) migrateActiveDelegations(ctx sdk.Context, oldDelAddr, newDelAddr sdk.AccAddress) error {
	// Get all delegations for old address
	delegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, oldDelAddr, 65535) // max retrieve
	if err != nil {
		return err
	}

	oldDelAddrStr, err := k.accountKeeper.AddressCodec().BytesToString(oldDelAddr)
	if err != nil {
		return err
	}

	newDelAddrStr, err := k.accountKeeper.AddressCodec().BytesToString(newDelAddr)
	if err != nil {
		return err
	}

	for _, delegation := range delegations {
		valAddr, err := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		if err != nil {
			return fmt.Errorf("failed to parse validator address %s: %w", delegation.ValidatorAddress, err)
		}

		k.Logger().Info("Migrating delegation",
			"old_delegator", oldDelAddrStr,
			"new_delegator", newDelAddrStr,
			"validator", delegation.ValidatorAddress,
			"shares", delegation.Shares.String())

		// Remove old delegation
		if err := k.stakingKeeper.RemoveDelegation(ctx, delegation); err != nil {
			return fmt.Errorf("failed to remove old delegation: %w", err)
		}

		// Create new delegation with same shares
		newDelegation := stakingtypes.NewDelegation(newDelAddrStr, delegation.ValidatorAddress, delegation.Shares)

		// Set new delegation
		if err := k.stakingKeeper.SetDelegation(ctx, newDelegation); err != nil {
			return fmt.Errorf("failed to set new delegation: %w", err)
		}

		// Initialize distribution tracking for the new delegation
		// Get the validator to calculate proper stake (tokens from shares)
		validator, err := k.stakingKeeper.GetValidator(ctx, valAddr)
		if err != nil {
			return fmt.Errorf("failed to get validator for distribution initialization: %w", err)
		}

		// Calculate delegation stake in tokens (not shares)
		stake := validator.TokensFromSharesTruncated(delegation.Shares)

		// Get the current validator period to properly initialize the delegation
		valCurrentRewards, err := k.distributionKeeper.GetValidatorCurrentRewards(ctx, valAddr)
		if err != nil {
			return fmt.Errorf("failed to get validator current rewards: %w", err)
		}

		// For new delegations, use current period as starting point
		currentPeriod := valCurrentRewards.Period

		// Create new starting info with current period
		startingInfo := distrtypes.NewDelegatorStartingInfo(currentPeriod, stake, uint64(ctx.BlockHeight()))
		err = k.distributionKeeper.SetDelegatorStartingInfo(ctx, valAddr, newDelAddr, startingInfo)
		if err != nil {
			return fmt.Errorf("failed to set delegator starting info for new delegation: %w", err)
		}

		k.Logger().Info("Successfully migrated delegation",
			"old_delegator", oldDelAddrStr,
			"new_delegator", newDelAddrStr,
			"validator", delegation.ValidatorAddress,
			"shares", delegation.Shares.String(),
			"stake_tokens", stake.String(),
			"starting_period", currentPeriod)
	}

	return nil
}

// migrateRewardTrackingSimple migrates reward tracking by withdrawing all rewards and transferring them
// This follows the pattern from MigrateDelegatorWithRewards for proper reward handling
func (k Keeper) migrateRewardTrackingSimple(ctx sdk.Context, oldDelAddr, newDelAddr sdk.AccAddress) error {
	// Get all delegations for reward calculation
	delegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, oldDelAddr, 65535)
	if err != nil {
		return err
	}

	totalRewards := sdk.NewCoins()

	k.Logger().Info("Starting reward migration",
		"old_delegator", oldDelAddr.String(),
		"new_delegator", newDelAddr.String(),
		"delegation_count", len(delegations))

	// Calculate and withdraw all rewards
	for _, delegation := range delegations {
		valAddr, err := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		if err != nil {
			return fmt.Errorf("failed to parse validator address %s: %w", delegation.ValidatorAddress, err)
		}

		// Withdraw rewards from the old delegator
		withdrawnRewards, err := k.distributionKeeper.WithdrawDelegationRewards(ctx, oldDelAddr, valAddr)
		if err != nil {
			// Log but continue - rewards might not exist yet
			k.Logger().Info("No rewards to withdraw during migration",
				"delegator", oldDelAddr.String(),
				"validator", delegation.ValidatorAddress,
				"error", err.Error())
		} else if !withdrawnRewards.IsZero() {
			totalRewards = totalRewards.Add(withdrawnRewards...)
			k.Logger().Info("Withdrew delegation rewards",
				"delegator", oldDelAddr.String(),
				"validator", delegation.ValidatorAddress,
				"rewards", withdrawnRewards.String())
		}

		// Clean up old delegator starting info to ensure clean state
		err = k.distributionKeeper.DeleteDelegatorStartingInfo(ctx, valAddr, oldDelAddr)
		if err != nil {
			// Log but continue - starting info might not exist
			k.Logger().Info("No starting info to delete during migration",
				"delegator", oldDelAddr.String(),
				"validator", delegation.ValidatorAddress,
				"error", err.Error())
		}
	}

	// Transfer accumulated rewards from old to new address
	if !totalRewards.IsZero() {
		err := k.bankKeeper.SendCoins(ctx, oldDelAddr, newDelAddr, totalRewards)
		if err != nil {
			return fmt.Errorf("failed to transfer rewards from old to new address: %w", err)
		}
		k.Logger().Info("Transferred total rewards",
			"from", oldDelAddr.String(),
			"to", newDelAddr.String(),
			"total_rewards", totalRewards.String())
	}

	// Note: New delegations will automatically start tracking rewards
	// when they are created in migrateActiveDelegations

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
