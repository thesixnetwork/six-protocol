package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) DisableVirtualSchemaProposal(goCtx context.Context, msg *types.MsgDisableVirtualSchemaProposal) (*types.MsgDisableVirtualSchemaProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	virtualSchema, found := k.GetVirtualSchema(ctx, msg.VirtualNftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.VirtualNftSchemaCode)
	}

	if !virtualSchema.Enable {
		return nil, sdkerrors.Wrap(types.ErrSchemaIsDisable, msg.VirtualNftSchemaCode)
	}

	isOwner := false
	// check if creator is path of schema registry owner
	for _, registry := range virtualSchema.Registry {
		schema, found := k.GetNFTSchema(ctx, registry.NftSchemaCode)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, registry.NftSchemaCode)
		}
		if schema.Owner == msg.Creator {
			isOwner = true
		}
	}
	if !isOwner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, msg.VirtualNftSchemaCode)
	}

	return &types.MsgDisableVirtualSchemaProposalResponse{}, nil
}
