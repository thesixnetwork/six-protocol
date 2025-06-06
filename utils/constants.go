package utils

import (
	"math"
	"math/big"

	sdkmath "cosmossdk.io/math"
)

var (
	Big0      = big.NewInt(0)
	Big1      = big.NewInt(1)
	Big2      = big.NewInt(2)
	Big8      = big.NewInt(8)
	Big27     = big.NewInt(27)
	Big35     = big.NewInt(35)
	BigMaxI64 = big.NewInt(math.MaxInt64)
	BigMaxU64 = new(big.Int).SetUint64(math.MaxUint64)
)

var Sdk0 = sdkmath.NewInt(0)
