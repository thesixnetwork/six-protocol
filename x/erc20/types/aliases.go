// Package types provides type definitions for the x/erc20 module.
// Core protobuf types (TokenPair, GenesisState, Params, Msg* types) are
// type-aliased from the evmos erc20 types package to ensure proto wire
// compatibility and avoid duplicating generated code.
package types

import (
	evmostypes "github.com/evmos/evmos/v20/x/erc20/types"
)

// Re-export types from evmos erc20 types for use within this package
type (
	// TokenPair defines an instance that records a pairing consisting of a native
	// Cosmos Coin and an ERC20 token address.
	TokenPair = evmostypes.TokenPair

	// Owner enumerates the ownership of an ERC20 contract.
	Owner = evmostypes.Owner

	// GenesisState defines the erc20 module genesis state.
	GenesisState = evmostypes.GenesisState

	// Params defines the parameters of the erc20 module.
	Params = evmostypes.Params

	// MsgConvertERC20 defines a message to convert ERC20 tokens to native Cosmos coins.
	MsgConvertERC20 = evmostypes.MsgConvertERC20

	// MsgConvertERC20Response is the response type of ConvertERC20.
	MsgConvertERC20Response = evmostypes.MsgConvertERC20Response

	// MsgConvertCoin defines a message to convert a native Cosmos coin to ERC20.
	MsgConvertCoin = evmostypes.MsgConvertCoin

	// MsgConvertCoinResponse is the response type of ConvertCoin.
	MsgConvertCoinResponse = evmostypes.MsgConvertCoinResponse

	// MsgUpdateParams defines a message to update the erc20 module params.
	MsgUpdateParams = evmostypes.MsgUpdateParams

	// MsgUpdateParamsResponse is the response type of UpdateParams.
	MsgUpdateParamsResponse = evmostypes.MsgUpdateParamsResponse

	// MsgRegisterERC20 defines a message to register an ERC20 contract.
	MsgRegisterERC20 = evmostypes.MsgRegisterERC20

	// MsgRegisterERC20Response is the response type of RegisterERC20.
	MsgRegisterERC20Response = evmostypes.MsgRegisterERC20Response

	// MsgToggleConversion defines a message to toggle erc20 conversion.
	MsgToggleConversion = evmostypes.MsgToggleConversion

	// MsgToggleConversionResponse is the response type of ToggleConversion.
	MsgToggleConversionResponse = evmostypes.MsgToggleConversionResponse

	// QueryTokenPairsRequest is the request type for the Query/TokenPairs RPC method.
	QueryTokenPairsRequest = evmostypes.QueryTokenPairsRequest

	// QueryTokenPairsResponse is the response type for the Query/TokenPairs RPC method.
	QueryTokenPairsResponse = evmostypes.QueryTokenPairsResponse

	// QueryTokenPairRequest is the request type for the Query/TokenPair RPC method.
	QueryTokenPairRequest = evmostypes.QueryTokenPairRequest

	// QueryTokenPairResponse is the response type for the Query/TokenPair RPC method.
	QueryTokenPairResponse = evmostypes.QueryTokenPairResponse

	// QueryParamsRequest is the request type for the Query/Params RPC method.
	QueryParamsRequest = evmostypes.QueryParamsRequest

	// QueryParamsResponse is the response type for the Query/Params RPC method.
	QueryParamsResponse = evmostypes.QueryParamsResponse

	// MsgServer is the server API for the erc20 Msg service.
	MsgServer = evmostypes.MsgServer

	// QueryServer is the server API for the erc20 Query service.
	QueryServer = evmostypes.QueryServer

	// Subspace defines an interface that implements the legacy Cosmos SDK x/params Subspace type.
	Subspace = evmostypes.Subspace
)

// Re-export Owner enum constants
const (
	OWNER_UNSPECIFIED = evmostypes.OWNER_UNSPECIFIED
	OWNER_MODULE      = evmostypes.OWNER_MODULE
	OWNER_EXTERNAL    = evmostypes.OWNER_EXTERNAL
)

// Re-export functions from evmos erc20 types
var (
	NewTokenPair       = evmostypes.NewTokenPair
	NewTokenPairSTRv2  = evmostypes.NewTokenPairSTRv2
	NewGenesisState    = evmostypes.NewGenesisState
	DefaultGenesisState = evmostypes.DefaultGenesisState
	NewParams          = evmostypes.NewParams
	DefaultParams      = evmostypes.DefaultParams
	ValidatePrecompiles = evmostypes.ValidatePrecompiles
	SanitizeERC20Name  = evmostypes.SanitizeERC20Name
	EqualMetadata      = evmostypes.EqualMetadata
	IsModuleAccount    = evmostypes.IsModuleAccount
	GetDisabledAndEnabledPrecompiles = evmostypes.GetDisabledAndEnabledPrecompiles
	CreateDenom        = evmostypes.CreateDenom
	CreateDenomDescription = evmostypes.CreateDenomDescription

	RegisterMsgServer          = evmostypes.RegisterMsgServer
	RegisterQueryServer        = evmostypes.RegisterQueryServer
	RegisterQueryHandlerClient = evmostypes.RegisterQueryHandlerClient
	NewQueryClient             = evmostypes.NewQueryClient

	RegisterInterfaces         = evmostypes.RegisterInterfaces
	RegisterLegacyAminoCodec   = evmostypes.RegisterLegacyAminoCodec

	// ParamStoreKeys
	ParamStoreKeyEnableErc20        = evmostypes.ParamStoreKeyEnableErc20
	ParamStoreKeyDynamicPrecompiles = evmostypes.ParamStoreKeyDynamicPrecompiles
	ParamStoreKeyNativePrecompiles  = evmostypes.ParamStoreKeyNativePrecompiles

	// Erc20Bytecode is the bytecode of the ERC20 precompile
	Erc20Bytecode = evmostypes.Erc20Bytecode
)
