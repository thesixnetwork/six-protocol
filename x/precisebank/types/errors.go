package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrInsufficientFunds       = errorsmod.Register(ModuleName, 2, "insufficient funds")
	ErrInvalidFractionalAmount = errorsmod.Register(ModuleName, 3, "invalid fractional amount")
	ErrAmountNotPositive       = errorsmod.Register(ModuleName, 4, "amount must be positive")
)
