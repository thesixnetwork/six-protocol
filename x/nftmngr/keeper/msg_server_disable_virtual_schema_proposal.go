package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) DisableVirtualSchemaProposal(goCtx context.Context, msg *types.MsgDisableVirtualSchemaProposal) (*types.MsgDisableVirtualSchemaProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentBlock := ctx.BlockHeight()
	proposalEndBlock, err := strconv.ParseInt(msg.ProposalExpiredBlock, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidHeight, msg.ProposalExpiredBlock)
	}

	if currentBlock >= proposalEndBlock {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidHeight, msg.ProposalExpiredBlock)
	}

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
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
	}

	// iterate to assing id of proposal
	allProposal := k.GetAllDisableVirtualSchema(ctx)
	proposalId := len(allProposal) + 1

	k.SetDisableVirtualSchema(ctx, types.DisableVirtualSchema{
		Id:                   strconv.FormatInt(int64(proposalId), 10),
		VirtualNftSchemaCode: msg.VirtualNftSchemaCode,
		ProposalExpiredBlock: msg.ProposalExpiredBlock,
	})

	return &types.MsgDisableVirtualSchemaProposalResponse{}, nil
}
