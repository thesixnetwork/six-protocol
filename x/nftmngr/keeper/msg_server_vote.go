package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) VoteVirtualSchemaProposal(goCtx context.Context, msg *types.MsgVoteVirtualSchemaProposal) (*types.MsgVoteVirtualSchemaProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	virtualSchemaProposal, found := k.GetVirtualSchemaProposal(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrProposalIdDoesNotExists, msg.Id)
	}

	// chck if proposal still active
	if active := k.IsProposalActive(ctx, virtualSchemaProposal); !active {
		return nil, sdkerrors.Wrap(types.ErrProposalExpired, msg.Id)
	}

	srcSchema, found := k.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	if srcSchema.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
	}

	registryIndex := -1
	for i, registry := range virtualSchemaProposal.VirtualSchema.Registry {
		if registry.NftSchemaCode == msg.NftSchemaCode {
			registryIndex = i
			// Check if already voted
			if registry.Decision != types.RegistryStatus_PENDING {
				return nil, sdkerrors.Wrapf(types.ErrAlreadyVote, "schema %s has already voted", msg.NftSchemaCode)
			}
			break
		}
	}

	if registryIndex == -1 {
		return nil, sdkerrors.Wrapf(types.ErrSchemaNotInRegistry, "schema %s not found in registry", msg.NftSchemaCode)
	}

	// Update vote
	virtualSchemaProposal.VirtualSchema.Registry[registryIndex].Decision = msg.Option

	k.SetVirtualSchemaProposal(ctx, virtualSchemaProposal)

	return &types.MsgVoteVirtualSchemaProposalResponse{}, nil
}
