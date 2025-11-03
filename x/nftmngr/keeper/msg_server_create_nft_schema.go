package keeper

import (
	"context"
	"encoding/base64"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	errormod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateNftschema(goCtx context.Context, msg *types.MsgCreateNFTSchema) (*types.MsgCreateNFTSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	jsonSchema, err := base64.StdEncoding.DecodeString(msg.NftSchemaBase64)
	if err != nil {
		return nil, errormod.Wrap(types.ErrParsingBase64, err.Error())
	}

	schema_input := types.NFTSchemaINPUT{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(jsonSchema, &schema_input)
	if err != nil {
		return nil, errormod.Wrap(types.ErrParsingSchemaMessage, err.Error())
	}

	err = k.CreateNftSchemaKeeper(ctx, msg.Creator, schema_input)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateSchema,
			sdk.NewAttribute(types.AttributeKeyCreateSchemaCode, schema_input.Code),
			sdk.NewAttribute(types.AttributeKeyCreateSchemaResult, "success"),
		),
	})

	return &types.MsgCreateNFTSchemaResponse{
		Code: schema_input.Code,
	}, nil
}
