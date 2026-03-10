package keeper

import (
	"context"
	"encoding/base64"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	errormod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMetadata(goCtx context.Context, msg *types.MsgCreateMetadata) (*types.MsgCreateMetadataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	newMetadata, err := base64.StdEncoding.DecodeString(msg.Base64NFTData)
	if err != nil {
		return nil, errormod.Wrap(types.ErrParsingBase64, err.Error())
	}

	metadata := types.NftData{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(newMetadata, &metadata)
	if err != nil {
		return nil, errormod.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	err = k.CreateNewMetadataKeeper(ctx, msg.Creator, msg.NftSchemaCode, msg.TokenId, metadata)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateMetadata,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyCreateMetaDataTokenID, msg.TokenId),
			sdk.NewAttribute(types.AttributeKeyCreateMetaDataResult, "success"),
		),
	})

	return &types.MsgCreateMetadataResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}

func (k msgServer) SetMintauth(goCtx context.Context, msg *types.MsgSetMintauth) (*types.MsgSetMintauthResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// Get nft schema from store
	schema, schemaFound := k.GetNFTSchema(ctx, msg.NftSchemaCode)
	// Check if the schema already exists
	if !schemaFound {
		return nil, errormod.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
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

func (k msgServer) SetMetadataFormat(goCtx context.Context, msg *types.MsgSetMetadataFormat) (*types.MsgSetMetadataFormatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	k.SetMetadataFormatKeeper(ctx, msg.Creator, msg.SchemaCode, msg.NewFormat)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetMetadataFormat,
			sdk.NewAttribute(types.EventTypeSetMetadataFormat, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NewFormat),
		),
	})

	return &types.MsgSetMetadataFormatResponse{
		SchemaCode: msg.SchemaCode,
		NewFormat:  msg.NewFormat,
	}, nil
}
