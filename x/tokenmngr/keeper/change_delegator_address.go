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

// This file implements delegator address migration following the improved MigrateDelegatorWithRewards pattern.
//
// Key improvements over the previous version:
// 1. Proper Reward Tracking: Preserves reward calculation continuity with starting info
// 2. Active Delegations Only: Only migrates active delegations, ignoring unbonding delegations
// 3. Starting Info Preservation: Properly handles delegator starting info to avoid reward calculation bugs
// 4. Validation: Ensures addresses are different and new address is clean
// 5. Error Handling: Comprehensive error handling with specific error types
//
// The implementation focuses on two main operations:
// 1. migrateRewardTrackingWithStartingInfo: Withdraws rewards and sets proper starting info
// 2. migrateActiveDelegations: Migrates delegations without resetting starting info// ChangeDelegatorAddress migrates active delegations and rewards from old delegator to new delegator
// This function only handles active delegations and reward tracking, ignoring unbonding delegations
// It follows the improved MigrateDelegatorWithRewards pattern for proper reward handling
func (k Keeper) ChangeDelegatorAddress(goCtx context.Context, oldAddress, newAddress sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Validate addresses
	if oldAddress.Empty() || newAddress.Empty() {
		return errorsmod.Wrap(stakingtypes.ErrEmptyDelegatorAddr, "addresses cannot be empty")
	}

	if oldAddress.Equals(newAddress) {
		return types.ErrSameDelegatorAddress
	}

	// Check if new address already has delegations
	existingDelegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, newAddress, 1)
	if err != nil {
		return errorsmod.Wrap(err, "failed to check existing delegations for new address")
	}
	if len(existingDelegations) > 0 {
		return types.ErrNewDelegatorAlreadyExists
	}

	// 2. FIRST: Handle rewards before touching delegations
	if err := k.migrateRewardTrackingWithStartingInfo(ctx, oldAddress, newAddress); err != nil {
		return errorsmod.Wrap(err, "failed to migrate reward tracking")
	}

	// 3. Then migrate active delegations
	if err := k.migrateActiveDelegations(ctx, oldAddress, newAddress); err != nil {
		return errorsmod.Wrap(err, "failed to migrate active delegations")
	}

	return nil
}

// migrateActiveDelegations migrates all active delegations from old to new delegator
// This follows the improved pattern that avoids resetting starting info already set by migrateRewardTrackingWithStartingInfo
func (k Keeper) migrateActiveDelegations(ctx sdk.Context, oldDelAddr, newDelAddr sdk.AccAddress) error {
	// Get all delegations for old address
	delegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, oldDelAddr, 65535)
	if err != nil {
		return err
	}

	if len(delegations) == 0 {
		k.Logger().Info("No delegations found for migration", "old_delegator", oldDelAddr.String())
		return nil
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

		// DON'T initialize distribution tracking here since it was already done properly
		// in migrateRewardTrackingWithStartingInfo with the correct starting period and stake

		k.Logger().Info("Successfully migrated delegation",
			"old_delegator", oldDelAddrStr,
			"new_delegator", newDelAddrStr,
			"validator", delegation.ValidatorAddress,
			"shares", delegation.Shares.String())
	}

	return nil
}

// migrateRewardTrackingWithStartingInfo properly migrates reward tracking including starting info
// This preserves reward calculation continuity by properly handling delegator starting info
func (k Keeper) migrateRewardTrackingWithStartingInfo(ctx sdk.Context, oldDelAddr, newDelAddr sdk.AccAddress) error {
	// Get all delegations for old delegator
	delegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, oldDelAddr, 65535)
	if err != nil {
		return err
	}

	if len(delegations) == 0 {
		k.Logger().Info("No delegations found for migration", "old_delegator", oldDelAddr.String())
		return nil
	}

	totalRewards := sdk.NewCoins()

	k.Logger().Info("Starting reward migration with starting info preservation",
		"old_delegator", oldDelAddr.String(),
		"new_delegator", newDelAddr.String(),
		"delegation_count", len(delegations))

	// Process each delegation
	for _, delegation := range delegations {
		valAddr, err := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		if err != nil {
			k.Logger().Error("Failed to parse validator address, skipping",
				"validator", delegation.ValidatorAddress,
				"error", err.Error())
			continue
		}

		// Get validator for stake calculation
		validator, err := k.stakingKeeper.GetValidator(ctx, valAddr)
		if err != nil {
			k.Logger().Error("Failed to get validator, skipping",
				"validator", delegation.ValidatorAddress,
				"error", err.Error())
			continue
		}

		// 1. Get the old delegator's starting info (if exists) for logging purposes
		_, err = k.distributionKeeper.GetDelegatorStartingInfo(ctx, valAddr, oldDelAddr)
		if err != nil {
			// If no starting info exists, we'll create a fresh one with current period
			k.Logger().Info("No existing starting info found for old delegator",
				"delegator", oldDelAddr.String(),
				"validator", delegation.ValidatorAddress)
		} else {
			k.Logger().Info("Found existing starting info for old delegator",
				"delegator", oldDelAddr.String(),
				"validator", delegation.ValidatorAddress)
		}

		// 2. Withdraw existing rewards for old delegator
		withdrawnRewards, err := k.distributionKeeper.WithdrawDelegationRewards(ctx, oldDelAddr, valAddr)
		if err != nil {
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

		// 3. Delete old delegator's starting info
		err = k.distributionKeeper.DeleteDelegatorStartingInfo(ctx, valAddr, oldDelAddr)
		if err != nil {
			k.Logger().Info("No starting info to delete",
				"delegator", oldDelAddr.String(),
				"validator", delegation.ValidatorAddress,
				"error", err.Error())
		}

		// 4. Initialize new delegator's starting info
		// Get current validator rewards to set proper starting period
		currentRewards, err := k.distributionKeeper.GetValidatorCurrentRewards(ctx, valAddr)
		if err != nil {
			k.Logger().Error("Failed to get validator current rewards, skipping",
				"validator", delegation.ValidatorAddress,
				"error", err.Error())
			continue
		}

		// Calculate proper stake from shares
		stake := validator.TokensFromSharesTruncated(delegation.Shares)

		// Set starting info for new delegator using current period and calculated stake
		// This ensures proper reward tracking going forward
		newStartingInfo := distrtypes.NewDelegatorStartingInfo(
			currentRewards.Period,
			stake,
			uint64(ctx.BlockHeight()),
		)

		err = k.distributionKeeper.SetDelegatorStartingInfo(ctx, valAddr, newDelAddr, newStartingInfo)
		if err != nil {
			return fmt.Errorf("failed to set new delegator starting info: %w", err)
		}

		k.Logger().Info("Set new delegator starting info",
			"new_delegator", newDelAddr.String(),
			"validator", delegation.ValidatorAddress,
			"period", currentRewards.Period,
			"stake", stake.String(),
			"height", ctx.BlockHeight())
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

// ValidateDelegatorMigration validates that the delegator migration was successful
// This is useful for testing and verification purposes
func (k Keeper) ValidateDelegatorMigration(goCtx context.Context, oldAddr, newAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check that old address has no more delegations
	oldDelegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, oldAddr, 65535)
	if err != nil {
		return fmt.Errorf("failed to check old delegations: %w", err)
	}
	if len(oldDelegations) > 0 {
		return fmt.Errorf("old address still has %d delegations after migration", len(oldDelegations))
	}

	// Check that new address has delegations
	newDelegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, newAddr, 65535)
	if err != nil {
		return fmt.Errorf("failed to check new delegations: %w", err)
	}
	if len(newDelegations) == 0 {
		return fmt.Errorf("new address has no delegations after migration")
	}

	// Validate distribution starting info exists for each delegation
	for _, delegation := range newDelegations {
		valAddr, err := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		if err != nil {
			return fmt.Errorf("failed to parse validator address: %w", err)
		}

		_, err = k.distributionKeeper.GetDelegatorStartingInfo(ctx, valAddr, newAddr)
		if err != nil {
			return fmt.Errorf("no starting info found for new delegator on validator %s: %w",
				delegation.ValidatorAddress, err)
		}
	}

	k.Logger().Info("Delegator migration validation successful",
		"old_delegator", oldAddr.String(),
		"new_delegator", newAddr.String(),
		"delegations_migrated", len(newDelegations))

	return nil
}
