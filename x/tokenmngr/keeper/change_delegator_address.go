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


func (k Keeper) ChangeDelegatorAddress(goCtx context.Context, oldAddress, newAddress sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fmt.Printf("Starting ChangeDelegatorAddress from %s to %s\n", oldAddress.String(), newAddress.String())

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

	// 2. Get all delegations for old delegator
	delegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, oldAddress, 65535)
	if err != nil {
		return errorsmod.Wrap(err, "failed to get delegations for old address")
	}

	if len(delegations) == 0 {
		k.Logger().Info("No delegations found for migration", "old_delegator", oldAddress.String())
		return nil // Nothing to migrate
	}

	// 3.	Migrate reward state to preserve continuity
	if err := k.migrateRewardState(ctx, oldAddress, newAddress, delegations); err != nil {
		return errorsmod.Wrap(err, "failed to migrate reward state")
	}

	// 4.  Migrate delegations without touching reward calculation
	if err := k.migrateDelegationsPreserveRewards(ctx, oldAddress, newAddress, delegations); err != nil {
		return errorsmod.Wrap(err, "failed to migrate delegations")
	}

	return nil
}

// migrateRewardState migrates the delegator starting info to preserve reward continuity
func (k Keeper) migrateRewardState(
	ctx sdk.Context,
	oldDelAddr, newDelAddr sdk.AccAddress,
	delegations []stakingtypes.Delegation,
) error {
	for _, delegation := range delegations {
		valAddr, err := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		if err != nil {
			k.Logger().Error("Failed to parse validator address, skipping",
				"validator", delegation.ValidatorAddress,
				"error", err.Error())
			continue
		}

		// Get the existing starting info from old delegator
		oldStartingInfo, err := k.distributionKeeper.GetDelegatorStartingInfo(ctx, valAddr, oldDelAddr)
		if err != nil {
			// If no starting info exists, create proper starting info
			// based on current delegation state
			validator, err := k.stakingKeeper.GetValidator(ctx, valAddr)
			if err != nil {
				k.Logger().Error("Failed to get validator, skipping",
					"validator", delegation.ValidatorAddress,
					"error", err.Error())
				continue
			}

			// Get current validator rewards period
			currentRewards, err := k.distributionKeeper.GetValidatorCurrentRewards(ctx, valAddr)
			if err != nil {
				k.Logger().Error("Failed to get validator current rewards, skipping",
					"validator", delegation.ValidatorAddress,
					"error", err.Error())
				continue
			}

			// Calculate stake from delegation shares
			stake := validator.TokensFromSharesTruncated(delegation.Shares)

			// Create new starting info with current state
			// Use current period, not previous period to avoid negative values
			newStartingInfo := distrtypes.NewDelegatorStartingInfo(
				currentRewards.Period,     
				stake,                     // Calculated stake from shares
				uint64(ctx.BlockHeight()), // Current height
			)

			err = k.distributionKeeper.SetDelegatorStartingInfo(ctx, valAddr, newDelAddr, newStartingInfo)
			if err != nil {
				k.Logger().Error("Failed to set new starting info",
					"new_delegator", newDelAddr.String(),
					"validator", delegation.ValidatorAddress,
					"error", err.Error())
				continue
			}

			k.Logger().Info("Created new starting info for delegator",
				"new_delegator", newDelAddr.String(),
				"validator", delegation.ValidatorAddress,
				"period", currentRewards.Period,
				"stake", stake.String())
		} else {

			// Check if old starting info has valid stake
			if oldStartingInfo.Stake.IsNil() || oldStartingInfo.Stake.IsZero() {
				// Recalculate stake from current delegation
				validator, err := k.stakingKeeper.GetValidator(ctx, valAddr)
				if err != nil {
					k.Logger().Error("Failed to get validator for stake calculation, skipping",
						"validator", delegation.ValidatorAddress,
						"error", err.Error())
					continue
				}

				// Calculate proper stake
				stake := validator.TokensFromSharesTruncated(delegation.Shares)
				oldStartingInfo.Stake = stake

				k.Logger().Info("Recalculated stake for starting info",
					"validator", delegation.ValidatorAddress,
					"old_stake", "nil/zero",
					"new_stake", stake.String())
			}

			// Transfer the EXACT starting info to preserve reward continuity
			err = k.distributionKeeper.SetDelegatorStartingInfo(ctx, valAddr, newDelAddr, oldStartingInfo)
			if err != nil {
				return fmt.Errorf("failed to transfer starting info to new delegator: %w", err)
			}

			// Delete the old starting info only after successful transfer
			err = k.distributionKeeper.DeleteDelegatorStartingInfo(ctx, valAddr, oldDelAddr)
			if err != nil {
				k.Logger().Info("No starting info to delete for old delegator",
					"old_delegator", oldDelAddr.String(),
					"validator", delegation.ValidatorAddress)
			}

			k.Logger().Info("Transferred exact starting info to preserve reward continuity",
				"old_delegator", oldDelAddr.String(),
				"new_delegator", newDelAddr.String(),
				"validator", delegation.ValidatorAddress,
				"period", oldStartingInfo.PreviousPeriod,
				"stake", oldStartingInfo.Stake.String(),
				"height", oldStartingInfo.Height)
		}
	}

	return nil
}

// migrateDelegationsPreserveRewards migrates delegations without calling reward-related hooks
func (k Keeper) migrateDelegationsPreserveRewards(
	ctx sdk.Context,
	oldDelAddr, newDelAddr sdk.AccAddress,
	delegations []stakingtypes.Delegation,
) error {
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

		// Remove old delegation WITHOUT calling reward hooks
		if err := k.stakingKeeper.RemoveDelegation(ctx, delegation); err != nil {
			return fmt.Errorf("failed to remove old delegation: %w", err)
		}

		// Create new delegation with exact same shares
		newDelegation := stakingtypes.NewDelegation(newDelAddrStr, delegation.ValidatorAddress, delegation.Shares)

		// Set new delegation
		if err := k.stakingKeeper.SetDelegation(ctx, newDelegation); err != nil {
			return fmt.Errorf("failed to set new delegation: %w", err)
		}


		k.Logger().Info("Successfully migrated delegation",
			"old_delegator", oldDelAddrStr,
			"new_delegator", newDelAddrStr,
			"validator", delegation.ValidatorAddress,
			"shares", delegation.Shares.String())
	}

	return nil
}






