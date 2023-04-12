package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func (k msgServer) EnableContractConverter(goCtx context.Context, msg *types.MsgEnableContractConverter) (*types.MsgEnableContractConverterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pass := k.protocoladminKeeper.Authenticate(ctx, "super.admin", msg.Creator)
	if !pass {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "message creator super admin")
	}

	// get current evmparams
	evmParams := k.evmKeeper.GetParams(ctx)
	evmParams.ConverterParams.Enable = msg.Enable
	k.evmKeeper.SetParams(ctx, evmParams)

	return &types.MsgEnableContractConverterResponse{
		ContractAddress: evmParams.ConverterParams.ConverterContract,
		Enable:          msg.Enable,
	}, nil
}
