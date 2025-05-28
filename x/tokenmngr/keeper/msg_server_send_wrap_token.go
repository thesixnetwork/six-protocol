package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendWrapToken(goCtx context.Context, msg *types.MsgSendWrapToken) (*types.MsgSendWrapTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendWrapTokenResponse{}, nil
}
