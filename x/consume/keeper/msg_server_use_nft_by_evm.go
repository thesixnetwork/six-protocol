package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/consume/types"
)

func (k msgServer) UseNftByEVM(goCtx context.Context, msg *types.MsgUseNftByEVM) (*types.MsgUseNftByEVMResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUseNftByEVMResponse{}, nil
}
