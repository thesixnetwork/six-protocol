package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/consume/types"
)

func (k msgServer) UseNft(goCtx context.Context, msg *types.MsgUseNft) (*types.MsgUseNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUseNftResponse{}, nil
}
