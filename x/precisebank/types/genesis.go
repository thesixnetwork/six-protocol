package types

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
)

// GenesisState defines the precisebank module's genesis state.
type GenesisState struct {
	Balances  FractionalBalances `json:"balances"`
	Remainder sdkmath.Int        `json:"remainder"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(
	balances FractionalBalances,
	remainder sdkmath.Int,
) *GenesisState {
	return &GenesisState{
		Balances:  balances,
		Remainder: remainder,
	}
}

// DefaultGenesisState returns a default genesis state.
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(FractionalBalances{}, sdkmath.ZeroInt())
}

// DefaultGenesis is an alias for DefaultGenesisState to follow six-protocol convention.
func DefaultGenesis() *GenesisState {
	return DefaultGenesisState()
}

// Validate performs basic validation of genesis data.
func (gs *GenesisState) Validate() error {
	if err := gs.Balances.Validate(); err != nil {
		return fmt.Errorf("invalid balances: %w", err)
	}

	if gs.Remainder.IsNil() {
		return fmt.Errorf("nil remainder amount")
	}

	if gs.Remainder.IsNegative() {
		return fmt.Errorf("negative remainder amount %s", gs.Remainder)
	}

	if gs.Remainder.GTE(conversionFactor) {
		return fmt.Errorf("remainder %v exceeds max of %v", gs.Remainder, conversionFactor.SubRaw(1))
	}

	// Sum(fractionalBalances) + remainder must be a whole integer value
	sum := gs.Balances.SumAmount()
	sumWithRemainder := sum.Add(gs.Remainder)

	offBy := sumWithRemainder.Mod(conversionFactor)

	if !offBy.IsZero() {
		return fmt.Errorf(
			"sum of fractional balances %v + remainder %v is not a multiple of %v",
			sum,
			gs.Remainder,
			conversionFactor,
		)
	}

	return nil
}

// TotalAmountWithRemainder returns the total amount of all balances in the
// genesis state, including both fractional balances and the remainder.
func (gs *GenesisState) TotalAmountWithRemainder() sdkmath.Int {
	return gs.Balances.SumAmount().Add(gs.Remainder)
}
