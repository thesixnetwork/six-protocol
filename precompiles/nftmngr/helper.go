package nftmngr

import (
	"encoding/json"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	erromod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/thesixnetwork/six-protocol/utils"

	nftmngrtype "github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (p PrecompileExecutor) AccAddressFromBech32(arg interface{}) (bec32Addr sdk.AccAddress, err error) {
	addr := arg.(string)
	bec32Addr, err = sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, erromod.Wrap(sdkerrors.ErrInvalidAddress, "invalid bech32 address")
	}
	return bec32Addr, nil
}

func (p PrecompileExecutor) AccAddressFromArg(arg interface{}) (sdk.AccAddress, error) {
	addr := arg.(common.Address)
	if addr == (common.Address{}) {
		return nil, errors.New("invalid addr")
	}
	bec32Addr := utils.EthToCosmosAddr(addr)
	return bec32Addr, nil
}

func (p PrecompileExecutor) StringFromArg(arg interface{}) (string, error) {
	stringArg, ok := arg.(string)
	if !ok {
		return "", errors.New("invalid argument type string")
	}
	return stringArg, nil
}

func (p PrecompileExecutor) ArrayOfstringFromArg(arg interface{}) ([]string, error) {
	arrayStringArg, ok := arg.([]string)
	if !ok {
		return nil, errors.New("invalid argument type string")
	}
	return arrayStringArg, nil
}

func (p PrecompileExecutor) BoolFromArg(arg interface{}) (bool, error) {
	boolArg, ok := arg.(bool)
	if !ok {
		return false, errors.New("invalid argument type string")
	}

	return boolArg, nil
}

func (p PrecompileExecutor) Uint64FromArg(arg interface{}) (uint64, error) {
	uint64Arg, ok := arg.(uint64)
	if !ok {
		return 0, errors.New("invalid argument type string")
	}

	return uint64Arg, nil
}

func (p PrecompileExecutor) Uint32FromArg(arg interface{}) (uint32, error) {
	uint32Arg, ok := arg.(uint32)
	if !ok {
		return 0, errors.New("invalid argument type string")
	}

	return uint32Arg, nil
}

func (p PrecompileExecutor) ParametersFromJSONString(arg interface{}) ([]*nftmngrtype.ActionParameter, error) {
	jsonStr, ok := arg.(string)
	if !ok {
		return nil, erromod.Wrap(sdkerrors.ErrInvalidRequest, "invalid argument type, expected string")
	}

	var params []nftmngrtype.ActionParameter
	if err := json.Unmarshal([]byte(jsonStr), &params); err != nil {
		return nil, erromod.Wrap(sdkerrors.ErrInvalidRequest, "invalid JSON format")
	}

	// Convert to slice of pointers to ActionParameter
	paramPointers := make([]*nftmngrtype.ActionParameter, len(params))
	for i := range params {
		paramPointers[i] = &params[i]
	}

	return paramPointers, nil
}

func (p PrecompileExecutor) ParseVoteOption(option int32) (nftmngrtype.RegistryStatus, error) {
	switch option {
	case 1:
		return nftmngrtype.RegistryStatus_REJECT, nil
	case 2:
		return nftmngrtype.RegistryStatus_ACCEPT, nil
	default:
		return nftmngrtype.RegistryStatus_PENDING, erromod.Wrap(sdkerrors.ErrInvalidRequest, "invalid vote option. Use 'YES' or 'NO'")
	}
}

func (p PrecompileExecutor) TokenIdMapFromJSONString(arg interface{}) ([]*nftmngrtype.TokenIdMap, error) {
	jsonStr, ok := arg.(string)
	if !ok {
		return nil, erromod.Wrap(sdkerrors.ErrInvalidRequest, "invalid argument type, expected string")
	}

	var tokenIdMap []nftmngrtype.TokenIdMap
	if err := json.Unmarshal([]byte(jsonStr), &tokenIdMap); err != nil {
		return nil, erromod.Wrap(sdkerrors.ErrInvalidRequest, "invalid JSON format")
	}

	// Convert to slice of pointers to ActionParameter
	tokenIdMapPointer := make([]*nftmngrtype.TokenIdMap, len(tokenIdMap))
	for i := range tokenIdMap {
		tokenIdMapPointer[i] = &tokenIdMap[i]
	}

	return tokenIdMapPointer, nil
}

func (p PrecompileExecutor) parseProposalType(option int32) (nftmngrtype.ProposalType, error) {
	switch option {
	case 0:
		return nftmngrtype.ProposalType_CREATE, nil
	case 1:
		return nftmngrtype.ProposalType_EDIT, nil
	default:
		return nftmngrtype.ProposalType_CREATE, erromod.Wrap(sdkerrors.ErrInvalidRequest, "Invalid proposal type. Use 'create(0)/edit(1)'")
	}
}
