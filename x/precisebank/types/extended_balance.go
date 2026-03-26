package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SumExtendedCoin returns a sdk.Coin of extended coin denomination
// with all integer and fractional amounts combined.
func SumExtendedCoin(amt sdk.Coins) sdk.Coin {
	// usix converted to asix
	integerAmount := amt.AmountOf(IntegerCoinDenom).Mul(conversionFactor)
	// asix as is
	extendedAmount := amt.AmountOf(ExtendedCoinDenom)

	// total of usix and asix amounts
	fullEmissionAmount := integerAmount.Add(extendedAmount)

	return sdk.NewCoin(
		ExtendedCoinDenom,
		fullEmissionAmount,
	)
}
