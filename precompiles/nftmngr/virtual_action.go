package nftmngr

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	pcommon "github.com/thesixnetwork/six-protocol/precompiles/common"
	nftmngrtypes "github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (p PrecompileExecutor) virtualSchemaProposal(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}

	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	proposalTypeUint, err := p.Uint32FromArg(args[0])
	if err != nil {
		return nil, err
	}

	proposalType, err := p.parseProposalType(int32(proposalTypeUint))
	if err != nil {
		return nil, err
	}

	base64VirtualSchemaProposal, err := p.StringFromArg(args[1])
	if err != nil {
		return nil, err
	}

	inputProposal, err := base64.StdEncoding.DecodeString(base64VirtualSchemaProposal)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingBase64, err.Error())
	}

	virtualSchemaRequest := nftmngrtypes.VirtualSchemaProposalRequest{}
	err = p.nftmngrKeeper.GetCodec().(*codec.ProtoCodec).UnmarshalJSON(inputProposal, &virtualSchemaRequest)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingMetadataMessage, err.Error())
	}

	virtualSchemaRegistryRequest := make([]*nftmngrtypes.VirtualSchemaRegistryRequest, len(virtualSchemaRequest.VirtualSchemaRegistry))

	for i, schemaCode := range virtualSchemaRequest.VirtualSchemaRegistry {
		virtualSchemaRegistryRequest[i] = &nftmngrtypes.VirtualSchemaRegistryRequest{
			NftSchemaCode: schemaCode,
		}
	}

	strPrposalId, err := p.nftmngrKeeper.ProposalVirtualSchemaKeeper(ctx, senderCosmoAddr.String(), virtualSchemaRequest.VirtualSchemaCode, proposalType, virtualSchemaRegistryRequest, virtualSchemaRequest.Actions, virtualSchemaRequest.Executors, virtualSchemaRequest.Enable)
	if err != nil {
		return nil, err
	}

	uintId, err := strconv.ParseUint(strPrposalId, 10, 64)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(uintId)
}

func (p PrecompileExecutor) voteVirtualSchema(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	fmt.Println("################# voteVirtualSchema #################")
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}

	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 3); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	proposalId, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	srcNftSchemaCode, err := p.StringFromArg(args[1])
	if err != nil {
		return nil, err
	}

	optionUint, err := p.Uint32FromArg(args[2])
	if err != nil {
		return nil, err
	}

	option, err := p.ParseVoteOption(int32(optionUint))
	if err != nil {
		return nil, err
	}

	err = p.nftmngrKeeper.VoteVirtualSchemaProposalKeeper(ctx, senderCosmoAddr.String(), proposalId, srcNftSchemaCode, option)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) perfromVirtualAction(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}

	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 5); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	virtualSchemaCode, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	tokenIdMap, err := p.TokenIdMapFromJSONString(args[1])
	if err != nil {
		return nil, err
	}

	actionName, err := p.StringFromArg(args[2])
	if err != nil {
		return nil, err
	}

	refId, err := p.StringFromArg(args[3])
	if err != nil {
		return nil, err
	}

	paramPointers, err := p.ParametersFromJSONString(args[4])
	if err != nil {
		return nil, err
	}

	_, err = p.nftmngrKeeper.PerformVirtualActionKeeper(ctx, senderCosmoAddr.String(), virtualSchemaCode, tokenIdMap, actionName, refId, paramPointers)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
