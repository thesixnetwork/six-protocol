package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func (k msgServer) SetConverterParams(goCtx context.Context, msg *types.MsgSetConverterParams) (*types.MsgSetConverterParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pass := k.protocoladminKeeper.Authenticate(ctx, "super.admin", msg.Creator)
	if !pass {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "message creator super admin")
	}

	// validate the params
	// validate ethereum contact address
	if !common.IsHexAddress(msg.ContractAddress) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid ethereum contract address")
	}

	// validate abi
	_, err := abi.JSON(strings.NewReader(msg.Abi))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid abi input")
	}

	// get current evmparams
	evmParams := k.evmKeeper.GetParams(ctx)
	var converterParams evmtypes.ConverterParams
	converterParams.ConverterContract = msg.ContractAddress
	converterParams.EventAbi = msg.Abi
	converterParams.EventName = msg.EventName
	converterParams.EventTuple = msg.EventTuple

	// set the converter params
	evmParams.ConverterParams = converterParams
	k.evmKeeper.SetParams(ctx, evmParams)

	return &types.MsgSetConverterParamsResponse{
		ContractAddress: msg.ContractAddress,
		Abi:             msg.Abi,
	}, nil
}
