package types

import (
	"fmt"
	"strings"

	sdkmath "cosmossdk.io/math"
)

// FractionalBalances is a slice of FractionalBalance
type FractionalBalances []FractionalBalance

// Validate returns an error if any FractionalBalance in the slice is invalid.
func (fbs FractionalBalances) Validate() error {
	seenAddresses := make(map[string]struct{})

	for _, fb := range fbs {
		if err := fb.Validate(); err != nil {
			return fmt.Errorf("invalid fractional balance for %s: %w", fb.Address, err)
		}

		// Make addresses all lowercase for unique check
		lowerAddr := strings.ToLower(fb.Address)

		if _, found := seenAddresses[lowerAddr]; found {
			return fmt.Errorf("duplicate address %v", lowerAddr)
		}

		seenAddresses[lowerAddr] = struct{}{}
	}

	return nil
}

// SumAmount returns the sum of all the amounts in the slice.
func (fbs FractionalBalances) SumAmount() sdkmath.Int {
	sum := sdkmath.ZeroInt()

	for _, fb := range fbs {
		sum = sum.Add(fb.Amount)
	}

	return sum
}
