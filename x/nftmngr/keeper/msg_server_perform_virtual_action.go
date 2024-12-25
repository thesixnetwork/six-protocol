package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) PerformVirtualAction(goCtx context.Context, msg *types.MsgPerformVirtualAction) (*types.MsgPerformVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Emit events on metadata change
	changeList, err := k.PerformVirtualKeeper(ctx, msg.Creator, msg.NftSchemaName, msg.TokenIdMap, msg.Action, msg.RefId, msg.Parameters)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.EventMessage, types.EventTypeRunAction),
			sdk.NewAttribute(types.AttributeKeyRunActionChangeList, string(changeList)),
		),
	)

	return &types.MsgPerformVirtualActionResponse{
		NftSchemaName: msg.NftSchemaName,
	}, nil
}
