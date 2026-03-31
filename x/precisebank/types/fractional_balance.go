package types

import (
	"fmt"
	"math/big"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// conversionFactor is used to convert the fractional balance to integer
// balances. 10^12, since usix is 10^6 and asix is 10^18.
var conversionFactor = sdkmath.NewInt(1_000_000_000_000)

// ConversionFactor returns a safe deep copy of the conversionFactor used to
// convert the fractional balance to integer balances. This is also 1 greater
// than the max valid fractional amount (999_999_999_999):
// 0 < FractionalBalance < conversionFactor
func ConversionFactor() sdkmath.Int {
	return sdkmath.NewIntFromBigInt(new(big.Int).Set(conversionFactor.BigInt()))
}

// NewFractionalBalance returns a new FractionalBalance with the given address and amount.
func NewFractionalBalance(address string, amount sdkmath.Int) FractionalBalance {
	return FractionalBalance{
		Address: address,
		Amount:  amount,
	}
}

// Validate returns an error if the FractionalBalance has an invalid address or
// negative amount.
func (fb FractionalBalance) Validate() error {
	if _, err := sdk.AccAddressFromBech32(fb.Address); err != nil {
		return err
	}

	return ValidateFractionalAmount(fb.Amount)
}

// ValidateFractionalAmount checks if an sdkmath.Int is a valid fractional
// amount, ensuring it is positive and less than the conversion factor.
func ValidateFractionalAmount(amt sdkmath.Int) error {
	if amt.IsNil() {
		return fmt.Errorf("nil amount")
	}

	if !amt.IsPositive() {
		return fmt.Errorf("non-positive amount %v", amt)
	}

	if amt.GTE(conversionFactor) {
		return fmt.Errorf("amount %v exceeds max of %v", amt, conversionFactor.SubRaw(1))
	}

	return nil
}
