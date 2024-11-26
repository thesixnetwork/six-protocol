package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PerformActionByAdmin(goCtx context.Context, msg *types.MsgPerformActionByAdmin) (*types.MsgPerformActionByAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Emit events on metadata change
	changeList, err := k.ActionByAdmin(ctx, msg.Creator, msg.NftSchemaCode, msg.TokenId, msg.Action, msg.RefId, msg.Parameters)
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

	return &types.MsgPerformActionByAdminResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}
