package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func (k Keeper) ChangeDelegatorAddress(goCtx context.Context, oldAddress, newAddress sdk.AccAddress) error {
	fmt.Println("debug: ChangeDelegatorAddress called======")

	// ctx := sdk.UnwrapSDKContext(goCtx)
	// Validate addresses
	if oldAddress.Empty() || newAddress.Empty() {
		return errorsmod.Wrap(stakingtypes.ErrEmptyDelegatorAddr, "addresses cannot be empty")
	}

	if oldAddress.Equals(newAddress) {
		return errorsmod.Wrap(stakingtypes.ErrBadDelegatorAddr, "old and new addresses are the same")
	}

	// err := k.withdrawAllRewards(ctx, oldAddress)
	// if err != nil {
	// 	return errorsmod.Wrap(err, "failed to withdraw all rewards")
	// }

	// err = k.distributionKeeper.SetWithdrawAddr(ctx, oldAddress, newAddress)
	// if err != nil {
	// 	return errorsmod.Wrap(err, "failed to set delegator withdraw address")
	// }

	// // Check if new address already has delegations to avoid conflicts
	// existingDelegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, newAddress, 1)
	// if err != nil {
	// 	return errorsmod.Wrap(err, "failed to check existing delegations for new address")
	// }
	// if len(existingDelegations) > 0 {
	// 	return errorsmod.Wrap(stakingtypes.ErrBadDelegatorAddr, "new address already has existing delegations")
	// }

	// // 1. Handle Delegations
	// if err := k.changeDelegations(ctx, oldAddress, newAddress); err != nil {
	// 	return errorsmod.Wrap(err, "failed to change delegations")
	// }

	// // 2. Handle Unbonding Delegations
	// if err := k.changeUnbondingDelegations(ctx, oldAddress, newAddress); err != nil {
	// 	return errorsmod.Wrap(err, "failed to change unbonding delegations")
	// }

	return nil
}

func (k Keeper) withdrawAllRewards(ctx sdk.Context, delegatorAddr sdk.AccAddress) error {
	// Get all delegations for the address
	delegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, delegatorAddr, 65535)
	if err != nil {
		return fmt.Errorf("failed to get delegations for reward withdrawal: %w", err)
	}

	k.Logger().Info("Withdrawing all rewards before migration",
		"delegator", delegatorAddr.String(),
		"delegation_count", len(delegations))

	// Withdraw rewards from each validator
	for _, delegation := range delegations {
		valAddr, err := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		if err != nil {
			return fmt.Errorf("failed to parse validator address %s: %w", delegation.ValidatorAddress, err)
		}

		// Withdraw rewards for this delegation
		rewards, err := k.distributionKeeper.WithdrawDelegationRewards(ctx, delegatorAddr, valAddr)
		if err != nil {
			// Log but continue - rewards might not exist
			k.Logger().Info("No rewards to withdraw for delegation",
				"delegator", delegatorAddr.String(),
				"validator", delegation.ValidatorAddress,
				"error", err.Error())
			continue
		}

		k.Logger().Info("Withdrew rewards for delegation",
			"delegator", delegatorAddr.String(),
			"validator", delegation.ValidatorAddress,
			"rewards", rewards.String())
	}

	return nil
}

// changeDelegations handles changing delegator address for all delegations
func (k Keeper) changeDelegations(ctx sdk.Context, oldAddress, newAddress sdk.AccAddress) error {
	// Get all delegations for the old address
	delegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, oldAddress, 65535) // max uint16
	if err != nil {
		return err
	}

	// Process each delegation
	for _, delegation := range delegations {
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
			delegation.Shares,
		)

		// Set the new delegation
		if err := k.stakingKeeper.SetDelegation(ctx, newDelegation); err != nil {
			return fmt.Errorf("failed to set new delegation: %w", err)
		}
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
