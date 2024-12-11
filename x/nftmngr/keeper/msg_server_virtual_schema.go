package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) CreateVirtualSchema(goCtx context.Context, msg *types.MsgCreateVirtualSchema) (*types.MsgCreateVirtualSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetVirtualSchema(
		ctx,
		msg.VirtualNftSchemaCode,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var virSchema = types.VirtualSchema{
		VirtualNftSchemaCode: "",
		Registry:             []*types.VirtualSchemaRegistry{},
	}

	k.SetVirtualSchema(
		ctx,
		virSchema,
	)
	return &types.MsgCreateVirtualSchemaResponse{}, nil
}

func (k msgServer) DeleteVirtualSchema(goCtx context.Context, msg *types.MsgDeleteVirtualSchema) (*types.MsgDeleteVirtualSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetVirtualSchema(
		ctx,
		msg.VirtualNftSchemaCode,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	_ = valFound.Registry
	// // Checks if the the msg creator is the same as the current owner
	// if msg.Creator != valFound.Creator {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	k.RemoveVirtualSchema(
		ctx,
		msg.VirtualNftSchemaCode,
	)

	return &types.MsgDeleteVirtualSchemaResponse{}, nil
}
