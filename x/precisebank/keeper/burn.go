package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/thesixnetwork/six-protocol/v4/x/precisebank/types"
)

// BurnCoins burns coins from the balance of the module account.
func (k Keeper) BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error {
	if moduleName == types.ModuleName {
		panic(errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "module account %s cannot be burned from", moduleName))
	}

	acc := k.ak.GetModuleAccount(ctx, moduleName)
	if acc == nil {
		panic(errorsmod.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", moduleName))
	}

	if !acc.HasPermission(authtypes.Burner) {
		panic(errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "module account %s does not have permissions to burn tokens", moduleName))
	}

	if !amt.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, amt.String())
	}

	passthroughCoins := amt
	extendedAmount := amt.AmountOf(types.ExtendedCoinDenom)
	if extendedAmount.IsPositive() {
		removeCoin := sdk.NewCoin(types.ExtendedCoinDenom, extendedAmount)
		passthroughCoins = amt.Sub(removeCoin)
	}

	if !passthroughCoins.Empty() {
		if err := k.bk.BurnCoins(ctx, moduleName, passthroughCoins); err != nil {
			return err
		}
	}

	if extendedAmount.IsPositive() {
		if err := k.burnExtendedCoin(ctx, moduleName, extendedAmount); err != nil {
			return err
		}
	}

	fullEmissionCoins := sdk.NewCoins(types.SumExtendedCoin(amt))
	if fullEmissionCoins.IsZero() {
		return nil
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvents(sdk.Events{
		banktypes.NewCoinBurnEvent(acc.GetAddress(), fullEmissionCoins),
		banktypes.NewCoinSpentEvent(acc.GetAddress(), fullEmissionCoins),
	})

	return nil
}

// burnExtendedCoin burns the fractional amount of the ExtendedCoinDenom from the module account.
func (k Keeper) burnExtendedCoin(
	ctx context.Context,
	moduleName string,
	amt sdkmath.Int,
) error {
	moduleAddr := k.ak.GetModuleAddress(moduleName)

	prevFractionalBalance := k.GetFractionalBalance(ctx, moduleAddr)
	prevRemainder := k.GetRemainderAmount(ctx)

	integerBurnAmount := amt.Quo(types.ConversionFactor())
	fractionalBurnAmount := amt.Mod(types.ConversionFactor())

	newFractionalBalance := prevFractionalBalance.Sub(fractionalBurnAmount)

	requiresBorrow := newFractionalBalance.IsNegative()

	newRemainder := prevRemainder.Add(fractionalBurnAmount)

	overflowingRemainder := newRemainder.GTE(types.ConversionFactor())

	// Case #1: optimization - direct burn instead of borrow + reserve burn
	if requiresBorrow && overflowingRemainder {
		newFractionalBalance = newFractionalBalance.Add(types.ConversionFactor())
		newRemainder = newRemainder.Sub(types.ConversionFactor())
		integerBurnAmount = integerBurnAmount.AddRaw(1)
	}

	// Case #2: Transfer 1 integer coin to reserve for integer borrow
	if requiresBorrow && !overflowingRemainder {
		newFractionalBalance = newFractionalBalance.Add(types.ConversionFactor())

		borrowCoin := sdk.NewCoin(types.IntegerCoinDenom, sdkmath.OneInt())
		if err := k.bk.SendCoinsFromModuleToModule(
			ctx,
			moduleName,
			types.ModuleName,
			sdk.NewCoins(borrowCoin),
		); err != nil {
			return k.updateInsufficientFundsError(ctx, moduleAddr, amt, err)
		}
	}

	// Case #3: No borrow needed, but remainder overflows
	if !requiresBorrow && overflowingRemainder {
		reserveBurnCoins := sdk.NewCoins(sdk.NewCoin(types.IntegerCoinDenom, sdkmath.OneInt()))
		if err := k.bk.BurnCoins(ctx, types.ModuleName, reserveBurnCoins); err != nil {
			return fmt.Errorf("failed to burn %s for reserve: %w", reserveBurnCoins, err)
		}
		newRemainder = newRemainder.Sub(types.ConversionFactor())
	}

	// Case #4: No borrow, no overflow - no additional work

	// Burn the integer amount
	if !integerBurnAmount.IsZero() {
		coin := sdk.NewCoin(types.IntegerCoinDenom, integerBurnAmount)
		if err := k.bk.BurnCoins(ctx, moduleName, sdk.NewCoins(coin)); err != nil {
			return k.updateInsufficientFundsError(ctx, moduleAddr, amt, err)
		}
	}

	// Set new fractional balance
	k.SetFractionalBalance(ctx, moduleAddr, newFractionalBalance)

	// Update remainder
	k.SetRemainderAmount(ctx, newRemainder)

	return nil
}
