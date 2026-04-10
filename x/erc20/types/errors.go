package types

import (
	evmostypes "github.com/evmos/evmos/v20/x/erc20/types"
)

// Re-export errors from evmos erc20 types to avoid double registration.
// These are already registered by the evmos types package init().
var (
	ErrERC20Disabled            = evmostypes.ErrERC20Disabled
	ErrInternalTokenPair        = evmostypes.ErrInternalTokenPair
	ErrTokenPairNotFound        = evmostypes.ErrTokenPairNotFound
	ErrTokenPairAlreadyExists   = evmostypes.ErrTokenPairAlreadyExists
	ErrUndefinedOwner           = evmostypes.ErrUndefinedOwner
	ErrBalanceInvariance        = evmostypes.ErrBalanceInvariance
	ErrUnexpectedEvent          = evmostypes.ErrUnexpectedEvent
	ErrABIPack                  = evmostypes.ErrABIPack
	ErrABIUnpack                = evmostypes.ErrABIUnpack
	ErrEVMDenom                 = evmostypes.ErrEVMDenom
	ErrEVMCall                  = evmostypes.ErrEVMCall
	ErrERC20TokenPairDisabled   = evmostypes.ErrERC20TokenPairDisabled
	ErrInvalidIBC               = evmostypes.ErrInvalidIBC
	ErrTokenPairOwnedByModule   = evmostypes.ErrTokenPairOwnedByModule
	ErrNativeConversionDisabled = evmostypes.ErrNativeConversionDisabled
)
