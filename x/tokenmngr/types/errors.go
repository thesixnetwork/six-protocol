package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/tokenmngr module sentinel errors
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample        = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")

	ErrSendCoinsFromAccountToModule = sdkerrors.Register(ModuleName, 1502, "unable to send coins from account to module")
	ErrBurnCoinsFromModuleAccount   = sdkerrors.Register(ModuleName, 1503, "unable to burn coins from module to account")
)
