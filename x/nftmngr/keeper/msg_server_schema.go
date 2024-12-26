package keeper

import (
	"context"
	"encoding/base64"
	"strconv"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateNFTSchema(goCtx context.Context, msg *types.MsgCreateNFTSchema) (*types.MsgCreateNFTSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	jsonSchema, err := base64.StdEncoding.DecodeString(msg.NftSchemaBase64)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}

	schema_input := types.NFTSchemaINPUT{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(jsonSchema, &schema_input)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingSchemaMessage, err.Error())
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

func (k msgServer) SetBaseUri(goCtx context.Context, msg *types.MsgSetBaseUri) (*types.MsgSetBaseUriResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	k.SetBaseURIKeeper(ctx, msg.Creator, msg.Code, msg.NewBaseUri)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetBaseURI,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.Code),
			sdk.NewAttribute(types.AttributeKeySetBaseURI, msg.NewBaseUri),
		),
	})

	return &types.MsgSetBaseUriResponse{
		Code: msg.Code,
		Uri:  msg.NewBaseUri,
	}, nil
}

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

func (k msgServer) SetOriginContract(goCtx context.Context, msg *types.MsgSetOriginContract) (*types.MsgSetOriginContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	err = k.SetOriginContractKeeper(ctx, msg.Creator, msg.SchemaCode, msg.NewContractAddress)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetOriginContract,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeySetOriginContract, msg.NewContractAddress),
			sdk.NewAttribute(types.AttributeKeySetOriginContractResult, "success"),
		),
	})

	return &types.MsgSetOriginContractResponse{
		SchemaCode:         msg.SchemaCode,
		NewContractAddress: msg.NewContractAddress,
	}, nil
}

func (k msgServer) SetOriginChain(goCtx context.Context, msg *types.MsgSetOriginChain) (*types.MsgSetOriginChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	err = k.SetOriginChainKeeper(ctx, msg.Creator, msg.SchemaCode, msg.NewOriginChain)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetBaseURI,
			sdk.NewAttribute(types.EventTypeSetOriginChain, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeySetOriginChain, msg.NewOriginChain),
			sdk.NewAttribute(types.AttributeKeySetOriginChainResult, "success"),
		),
	})

	return &types.MsgSetOriginChainResponse{
		SchemaCode:     msg.SchemaCode,
		NewOriginChain: msg.NewOriginChain,
	}, nil
}
