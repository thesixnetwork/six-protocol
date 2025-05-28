package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UnwrapToken(goCtx context.Context, msg *types.MsgUnwrapToken) (*types.MsgUnwrapTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUnwrapTokenResponse{}, nil
}
