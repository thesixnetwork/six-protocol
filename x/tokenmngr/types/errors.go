package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/tokenmngr module sentinel errors
var (
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")

	ErrSendCoinsFromAccountToModule = sdkerrors.Register(ModuleName, 1502, "unable to send coins from account to module")
	ErrBurnCoinsFromModuleAccount   = sdkerrors.Register(ModuleName, 1503, "unable to burn coins from module to account")
)
