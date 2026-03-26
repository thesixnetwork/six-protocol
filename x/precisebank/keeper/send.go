package keeper

import (
	"context"
	"errors"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/thesixnetwork/six-protocol/v4/x/precisebank/types"
)

// IsSendEnabledCoins uses the parent x/bank keeper to check the coins provided
// and returns an ErrSendDisabled if any of the coins are not configured for sending.
func (k Keeper) IsSendEnabledCoins(ctx context.Context, coins ...sdk.Coin) error {
	return k.bk.IsSendEnabledCoins(ctx, coins...)
}

// SendCoins transfers amt coins from a sending account to a receiving account.
// It handles transfers including ExtendedCoinDenom and supports non-ExtendedCoinDenom
// transfers by passing through to x/bank.
func (k Keeper) SendCoins(
	ctx context.Context,
	from, to sdk.AccAddress,
	amt sdk.Coins,
) error {
	if !amt.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, amt.String())
	}

	passthroughCoins := amt
	extendedCoinAmount := amt.AmountOf(types.ExtendedCoinDenom)

	// Remove the extended coin amount from the passthrough coins
	if extendedCoinAmount.IsPositive() {
		subCoin := sdk.NewCoin(types.ExtendedCoinDenom, extendedCoinAmount)
		passthroughCoins = amt.Sub(subCoin)
	}

	// Send the passthrough coins through x/bank
	if passthroughCoins.IsAllPositive() {
		if err := k.bk.SendCoins(ctx, from, to, passthroughCoins); err != nil {
			return err
		}
	}

	// Send the extended coin amount through x/precisebank
	if extendedCoinAmount.IsPositive() {
		if err := k.sendExtendedCoins(ctx, from, to, extendedCoinAmount); err != nil {
			return err
		}
	}

	fullEmissionCoins := sdk.NewCoins(types.SumExtendedCoin(amt))

	if fullEmissionCoins.IsZero() {
		return nil
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			banktypes.EventTypeTransfer,
			sdk.NewAttribute(banktypes.AttributeKeyRecipient, to.String()),
			sdk.NewAttribute(banktypes.AttributeKeySender, from.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, fullEmissionCoins.String()),
		),
		banktypes.NewCoinSpentEvent(from, fullEmissionCoins),
		banktypes.NewCoinReceivedEvent(to, fullEmissionCoins),
	})

	return nil
}

// sendExtendedCoins transfers amt extended coins from a sending account to a
// receiving account. Handles fractional balance borrow/carry operations.
func (k Keeper) sendExtendedCoins(
	ctx context.Context,
	from, to sdk.AccAddress,
	amt sdkmath.Int,
) error {
	senderFracBal := k.GetFractionalBalance(ctx, from)
	recipientFracBal := k.GetFractionalBalance(ctx, to)

	integerAmt := amt.Quo(types.ConversionFactor())
	fractionalAmt := amt.Mod(types.ConversionFactor())

	senderNewFracBal, senderNeedsBorrow := subFromFractionalBalance(senderFracBal, fractionalAmt)
	recipientNewFracBal, recipientNeedsCarry := addToFractionalBalance(recipientFracBal, fractionalAmt)

	// Case #1: Sender borrow, recipient carry
	if senderNeedsBorrow && recipientNeedsCarry {
		integerAmt = integerAmt.AddRaw(1)
	}

	// Full integer amount transfer
	if integerAmt.IsPositive() {
		transferCoin := sdk.NewCoin(types.IntegerCoinDenom, integerAmt)
		if err := k.bk.SendCoins(ctx, from, to, sdk.NewCoins(transferCoin)); err != nil {
			return k.updateInsufficientFundsError(ctx, from, amt, err)
		}
	}

	// Case #2: Sender borrow, NO recipient carry
	if senderNeedsBorrow && !recipientNeedsCarry {
		borrowCoin := sdk.NewCoin(types.IntegerCoinDenom, sdkmath.NewInt(1))
		if err := k.bk.SendCoinsFromAccountToModule(
			ctx,
			from,
			types.ModuleName,
			sdk.NewCoins(borrowCoin),
		); err != nil {
			return k.updateInsufficientFundsError(ctx, from, amt, err)
		}
	}

	// Case #3: NO sender borrow, recipient carry
	if !senderNeedsBorrow && recipientNeedsCarry {
		reserveAddr := k.ak.GetModuleAddress(types.ModuleName)

		carryCoin := sdk.NewCoin(types.IntegerCoinDenom, sdkmath.NewInt(1))
		if err := k.bk.SendCoins(
			ctx,
			reserveAddr,
			to,
			sdk.NewCoins(carryCoin),
		); err != nil {
			panic(fmt.Errorf("failed to carry fractional coins to %s: %w", to, err))
		}
	}

	// Case #4: NO sender borrow, NO recipient carry - no additional operations

	// Persist new fractional balances to store
	k.SetFractionalBalance(ctx, from, senderNewFracBal)
	k.SetFractionalBalance(ctx, to, recipientNewFracBal)

	return nil
}

// subFromFractionalBalance subtracts a fractional amount from the current balance,
// returning the new balance and whether an integer borrow is required.
func subFromFractionalBalance(
	currentFractionalBalance sdkmath.Int,
	amountToSub sdkmath.Int,
) (sdkmath.Int, bool) {
	if currentFractionalBalance.GTE(types.ConversionFactor()) {
		panic("currentFractionalBalance must be less than ConversionFactor")
	}
	if amountToSub.GTE(types.ConversionFactor()) {
		panic("amountToSub must be less than ConversionFactor")
	}

	newFractionalBalance := currentFractionalBalance.Sub(amountToSub)

	borrowRequired := newFractionalBalance.IsNegative()

	if borrowRequired {
		newFractionalBalance = newFractionalBalance.Add(types.ConversionFactor())
	}

	return newFractionalBalance, borrowRequired
}

// addToFractionalBalance adds a fractional amount to the current balance,
// returning the new balance and whether a carry is required.
func addToFractionalBalance(
	currentFractionalBalance sdkmath.Int,
	amountToAdd sdkmath.Int,
) (sdkmath.Int, bool) {
	if currentFractionalBalance.GTE(types.ConversionFactor()) {
		panic("currentFractionalBalance must be less than ConversionFactor")
	}
	if amountToAdd.GTE(types.ConversionFactor()) {
		panic("amountToAdd must be less than ConversionFactor")
	}

	newFractionalBalance := currentFractionalBalance.Add(amountToAdd)

	carryRequired := newFractionalBalance.GTE(types.ConversionFactor())

	if carryRequired {
		newFractionalBalance = newFractionalBalance.Sub(types.ConversionFactor())
	}

	return newFractionalBalance, carryRequired
}

// SendCoinsFromAccountToModule transfers coins from an account to a module.
func (k Keeper) SendCoinsFromAccountToModule(
	ctx context.Context,
	senderAddr sdk.AccAddress,
	recipientModule string,
	amt sdk.Coins,
) error {
	recipientAcc := k.ak.GetModuleAccount(ctx, recipientModule)
	if recipientAcc == nil {
		panic(errorsmod.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", recipientModule))
	}

	if recipientModule == types.ModuleName {
		return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "module account %s is not allowed to receive funds", types.ModuleName)
	}

	return k.SendCoins(ctx, senderAddr, recipientAcc.GetAddress(), amt)
}

// SendCoinsFromModuleToAccount transfers coins from a module to an account.
func (k Keeper) SendCoinsFromModuleToAccount(
	ctx context.Context,
	senderModule string,
	recipientAddr sdk.AccAddress,
	amt sdk.Coins,
) error {
	senderAddr := k.ak.GetModuleAddress(senderModule)
	if senderAddr == nil {
		panic(errorsmod.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", senderModule))
	}

	if senderModule == types.ModuleName {
		return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "module account %s is not allowed to send funds", types.ModuleName)
	}

	if k.bk.BlockedAddr(recipientAddr) {
		return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", recipientAddr)
	}

	return k.SendCoins(ctx, senderAddr, recipientAddr, amt)
}

// updateInsufficientFundsError returns a modified ErrInsufficientFunds with
// extended coin amounts if the error is due to insufficient funds.
func (k Keeper) updateInsufficientFundsError(
	ctx context.Context,
	addr sdk.AccAddress,
	amt sdkmath.Int,
	err error,
) error {
	if !errors.Is(err, sdkerrors.ErrInsufficientFunds) {
		return err
	}

	bal := k.GetBalance(ctx, addr, types.ExtendedCoinDenom)
	coin := sdk.NewCoin(types.ExtendedCoinDenom, amt)

	spendable := sdk.NewCoins(bal)

	return errorsmod.Wrapf(
		sdkerrors.ErrInsufficientFunds,
		"spendable balance %s is smaller than %s",
		spendable, coin,
	)
}
