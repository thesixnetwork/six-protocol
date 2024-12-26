package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// TODO:: Feat(VirtualSchema)
func (k msgServer) CreateVirtualAction(goCtx context.Context, msg *types.MsgCreateVirtualAction) (*types.MsgCreateVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	virtualSchema, found := k.GetVirtualSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	isOwner := false
	for _, schemaRegistry := range virtualSchema.Registry {
		srcSchema, found := k.GetNFTSchema(ctx, schemaRegistry.NftSchemaCode)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, schemaRegistry.NftSchemaCode)
		}
		if msg.Creator == srcSchema.Owner {
			isOwner = true
		}
	}

	if !isOwner {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
	}

	listNewVirtualAction := []*types.VirtualAction{}

	if err := msg.CheckDuplicateAction(); err != nil {
		return nil, err
	}

	for _, newAction := range msg.NewActions {

		virtualAction := types.VirtualAction{
			NftSchemaCode:   msg.NftSchemaCode,
			Name:            newAction.Name,
			Desc:            newAction.Desc,
			When:            newAction.When,
			Then:            newAction.Then,
			Disable:         newAction.Disable,
			AllowedActioner: newAction.AllowedActioner,
			Params:          newAction.Params,
		}

		// Check if the virtual action already exists
		_, found = k.GetVirtualAction(
			ctx,
			msg.NftSchemaCode,
			virtualAction.Name,
		)

		if found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
		}

		err := ValidateVirutualAction(virtualAction.ToAction())
		if err != nil {
			return nil, err
		}

		k.SetVirtualAction(
			ctx,
			virtualAction,
		)

		listNewVirtualAction = append(listNewVirtualAction, &virtualAction)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddAction,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
		),
	})

	return &types.MsgCreateVirtualActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		VirtualAction: listNewVirtualAction,
	}, nil
}

// TODO:: Feat(VirtualSchema)
func (k msgServer) UpdateVirtualAction(goCtx context.Context, msg *types.MsgUpdateVirtualAction) (*types.MsgUpdateVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	virtualSchema, found := k.GetVirtualSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	isOwner := false
	for _, schemaRegistry := range virtualSchema.Registry {
		srcSchema, found := k.GetNFTSchema(ctx, schemaRegistry.NftSchemaCode)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, schemaRegistry.NftSchemaCode)
		}
		if msg.Creator == srcSchema.Owner {
			isOwner = true
		}
	}

	if !isOwner {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
	}

	if err := msg.CheckDuplicateAction(); err != nil {
		return nil, err
	}

	listNewVirtualAction := []*types.VirtualAction{}

	for _, newAction := range msg.NewActions {
		toUpdateVirtualAction := types.VirtualAction{
			NftSchemaCode:   msg.NftSchemaCode,
			Name:            newAction.Name,
			Desc:            newAction.Desc,
			Disable:         newAction.Disable,
			When:            newAction.When,
			Then:            newAction.Then,
			AllowedActioner: newAction.AllowedActioner,
			Params:          newAction.Params,
		}

		// Check if the virtual action already exists
		_, found = k.GetVirtualAction(
			ctx,
			msg.NftSchemaCode,
			toUpdateVirtualAction.Name,
		)

		if !found {
			return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, toUpdateVirtualAction.Name)
		}

		k.SetVirtualAction(ctx, toUpdateVirtualAction)
		listNewVirtualAction = append(listNewVirtualAction, &toUpdateVirtualAction)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddAction,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
		),
	})

	return &types.MsgUpdateVirtualActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		VirtualAction: listNewVirtualAction,
	}, nil
}

// TODO:: Feat(VirtualSchema)
func (k msgServer) DeleteVirtualAction(goCtx context.Context, msg *types.MsgDeleteVirtualAction) (*types.MsgDeleteVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	virtualSchema, found := k.GetVirtualSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	isOwner := false
	for _, schemaRegistry := range virtualSchema.Registry {
		srcSchema, found := k.GetNFTSchema(ctx, schemaRegistry.NftSchemaCode)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, schemaRegistry.NftSchemaCode)
		}
		if msg.Creator == srcSchema.Owner {
			isOwner = true
		}
	}

	if !isOwner {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
	}

	// Check if the virtual action already exists
	_, found = k.GetVirtualAction(
		ctx,
		msg.NftSchemaCode,
		msg.Name,
	)

	if !found {
		return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, msg.Name)
	}

	k.RemoveVirtualAction(
		ctx,
		msg.NftSchemaCode,
		msg.Name,
	)

	return &types.MsgDeleteVirtualActionResponse{
		Creator: msg.Creator,
		Status:  "success",
	}, nil
}
