package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func (k msgServer) ConvertToAtto(goCtx context.Context, msg *types.MsgConvertToAtto) (*types.MsgConvertToAttoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgConvertToAttoResponse{}, nil
}
