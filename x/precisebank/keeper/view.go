package keeper

import (
	"context"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/v4/x/precisebank/types"
)

// GetBalance returns the balance of a specific denom for an address.
// For ExtendedCoinDenom, it returns the combined integer + fractional balance.
func (k Keeper) GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	// Module balance should display as empty for extended denom.
	if denom == types.ExtendedCoinDenom && addr.Equals(k.ak.GetModuleAddress(types.ModuleName)) {
		return sdk.NewCoin(denom, sdkmath.ZeroInt())
	}

	// Pass through to x/bank for denoms except ExtendedCoinDenom
	if denom != types.ExtendedCoinDenom {
		return k.bk.GetBalance(ctx, addr, denom)
	}

	// x/bank for integer balance
	integerCoins := k.bk.GetBalance(ctx, addr, types.IntegerCoinDenom)

	// x/precisebank for fractional balance
	fractionalAmount := k.GetFractionalBalance(ctx, addr)

	// (Integer * ConversionFactor) + Fractional
	fullAmount := integerCoins.Amount.Mul(types.ConversionFactor()).Add(fractionalAmount)

	return sdk.NewCoin(types.ExtendedCoinDenom, fullAmount)
}

// SpendableCoin returns the spendable balance for an account and denom.
func (k Keeper) SpendableCoin(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	if denom == types.ExtendedCoinDenom && addr.Equals(k.ak.GetModuleAddress(types.ModuleName)) {
		return sdk.NewCoin(denom, sdkmath.ZeroInt())
	}

	if denom != types.ExtendedCoinDenom {
		return k.bk.SpendableCoin(ctx, addr, denom)
	}

	// x/bank for integer balance - excluding locked
	integerCoin := k.bk.SpendableCoin(ctx, addr, types.IntegerCoinDenom)

	// x/precisebank for fractional balance
	fractionalAmount := k.GetFractionalBalance(ctx, addr)

	// Spendable = (Integer * ConversionFactor) + Fractional
	fullAmount := integerCoin.Amount.Mul(types.ConversionFactor()).Add(fractionalAmount)

	return sdk.NewCoin(types.ExtendedCoinDenom, fullAmount)
}
