package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) SetUriRetrievalMethod(goCtx context.Context, msg *types.MsgSetUriRetrievalMethod) (*types.MsgSetUriRetrievalMethodResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	k.Keeper.SetURIRetrievalKeeper(ctx, msg.Creator, msg.SchemaCode, msg.NewMethod)
	strMethod := strconv.FormatInt(int64(msg.NewMethod), 10)
	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetBaseURI,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeySetRetrievalMethod, strMethod),
			sdk.NewAttribute(types.AttributeKeySetRetrivalResult, "success"),
		),
	})

	return &types.MsgSetUriRetrievalMethodResponse{
		SchemaCode: msg.SchemaCode,
		NewMethod:  strMethod,
	}, nil
}
