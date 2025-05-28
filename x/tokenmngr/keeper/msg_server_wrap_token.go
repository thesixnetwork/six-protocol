package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) WrapToken(goCtx context.Context, msg *types.MsgWrapToken) (*types.MsgWrapTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgWrapTokenResponse{}, nil
}
