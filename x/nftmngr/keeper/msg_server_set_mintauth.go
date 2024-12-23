package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) SetMintauth(goCtx context.Context, msg *types.MsgSetMintauth) (*types.MsgSetMintauthResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// Get nft schema from store
	schema, schemaFound := k.GetNFTSchema(ctx, msg.NftSchemaCode)
	// Check if the schema already exists
	if !schemaFound {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	err = k.SetMintAuthKeeper(ctx, msg.Creator, msg.NftSchemaCode, msg.AuthorizeTo)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetMintAuth,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeAutorizeTo, schema.MintAuthorization),
			sdk.NewAttribute(types.AttributeKeySetMinAuthResult, "success"),
		),
	})

	return &types.MsgSetMintauthResponse{
		NftSchemaCode: msg.NftSchemaCode,
	}, nil
}
