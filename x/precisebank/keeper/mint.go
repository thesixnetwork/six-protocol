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

// MintCoins creates new coins from thin air and adds them to the module account.
// If ExtendedCoinDenom is provided, the corresponding fractional amount is
// added to the module state.
func (k Keeper) MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error {
	if moduleName == types.ModuleName {
		panic(errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "module account %s cannot be minted to", moduleName))
	}

	acc := k.ak.GetModuleAccount(ctx, moduleName)
	if acc == nil {
		panic(errorsmod.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", moduleName))
	}

	if !acc.HasPermission(authtypes.Minter) {
		panic(errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "module account %s does not have permissions to mint tokens", moduleName))
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
		if err := k.bk.MintCoins(ctx, moduleName, passthroughCoins); err != nil {
			return err
		}
	}

	if extendedAmount.IsPositive() {
		if err := k.mintExtendedCoin(ctx, moduleName, extendedAmount); err != nil {
			return err
		}
	}

	fullEmissionCoins := sdk.NewCoins(types.SumExtendedCoin(amt))
	if fullEmissionCoins.IsZero() {
		return nil
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvents(sdk.Events{
		banktypes.NewCoinMintEvent(acc.GetAddress(), fullEmissionCoins),
		banktypes.NewCoinReceivedEvent(acc.GetAddress(), fullEmissionCoins),
	})

	return nil
}

// mintExtendedCoin manages the minting of only extended coins, handling integer
// carry over from fractional balance and reserve management.
func (k Keeper) mintExtendedCoin(
	ctx context.Context,
	recipientModuleName string,
	amt sdkmath.Int,
) error {
	moduleAddr := k.ak.GetModuleAddress(recipientModuleName)

	fractionalAmount := k.GetFractionalBalance(ctx, moduleAddr)

	integerMintAmount := amt.Quo(types.ConversionFactor())
	fractionalMintAmount := amt.Mod(types.ConversionFactor())

	prevRemainder := k.GetRemainderAmount(ctx)

	newRemainder := prevRemainder.Sub(fractionalMintAmount)

	newFractionalBalance := fractionalAmount.Add(fractionalMintAmount)

	// Case #3 - Integer carry, remainder is sufficient
	if newFractionalBalance.GTE(types.ConversionFactor()) && newRemainder.GTE(sdkmath.ZeroInt()) {
		carryCoin := sdk.NewCoin(types.IntegerCoinDenom, sdkmath.OneInt())

		if err := k.bk.SendCoinsFromModuleToModule(
			ctx,
			types.ModuleName,
			recipientModuleName,
			sdk.NewCoins(carryCoin),
		); err != nil {
			return err
		}
	}

	// Case #4 - Integer carry, remainder is insufficient (optimization)
	if newFractionalBalance.GTE(types.ConversionFactor()) && newRemainder.IsNegative() {
		integerMintAmount = integerMintAmount.AddRaw(1)
	}

	// Adjust fractional balance if carry occurred
	if newFractionalBalance.GTE(types.ConversionFactor()) {
		newFractionalBalance = newFractionalBalance.Sub(types.ConversionFactor())
	}

	// Mint new integer amounts in x/bank
	if integerMintAmount.IsPositive() {
		integerMintCoin := sdk.NewCoin(types.IntegerCoinDenom, integerMintAmount)

		if err := k.bk.MintCoins(
			ctx,
			recipientModuleName,
			sdk.NewCoins(integerMintCoin),
		); err != nil {
			return err
		}
	}

	// Set new fractional balance
	k.SetFractionalBalance(ctx, moduleAddr, newFractionalBalance)

	// Mint additional reserve if remainder insufficient
	wasCarried := fractionalAmount.Add(fractionalMintAmount).GTE(types.ConversionFactor())
	if prevRemainder.LT(fractionalMintAmount) && !wasCarried {
		reserveMintCoins := sdk.NewCoins(sdk.NewCoin(types.IntegerCoinDenom, sdkmath.OneInt()))
		if err := k.bk.MintCoins(ctx, types.ModuleName, reserveMintCoins); err != nil {
			return fmt.Errorf("failed to mint %s for reserve: %w", reserveMintCoins, err)
		}
	}

	// Adjust remainder
	if newRemainder.IsNegative() {
		newRemainder = newRemainder.Add(types.ConversionFactor())
	}

	k.SetRemainderAmount(ctx, newRemainder)

	return nil
}
