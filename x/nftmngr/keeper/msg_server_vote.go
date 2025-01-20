package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) VoteVirtualSchemaProposal(goCtx context.Context, msg *types.MsgVoteVirtualSchemaProposal) (*types.MsgVoteVirtualSchemaProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.VoteVirtualSchemaProposalKeeper(ctx, msg.Creator, msg.Id, msg.NftSchemaCode, msg.Option)
	if err != nil {
		return nil, err
	}

	return &types.MsgVoteVirtualSchemaProposalResponse{}, nil
}

func (k Keeper) VoteVirtualSchemaProposalKeeper(ctx sdk.Context, creator, proposalId, srcNftSchemaCode string, option types.RegistryStatus) error {
	virtualSchemaProposal, found := k.GetVirtualSchemaProposal(ctx, proposalId)
	if !found {
		return sdkerrors.Wrap(types.ErrProposalIdDoesNotExists, proposalId)
	}

	// chck if proposal still active
	if active := k.IsProposalActive(ctx, virtualSchemaProposal); !active {
		return sdkerrors.Wrap(types.ErrProposalExpired, proposalId)
	}

	srcSchema, found := k.GetNFTSchema(ctx, srcNftSchemaCode)
	if !found {
		return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, srcNftSchemaCode)
	}

	if srcSchema.Owner != creator {
		return sdkerrors.Wrap(types.ErrUnauthorized, creator)
	}
	registryIndex := -1
	for i, registry := range virtualSchemaProposal.VirtualSchema.Registry {
		if registry.NftSchemaCode == srcNftSchemaCode {
			registryIndex = i
			// Check if already voted
			if registry.Decision != types.RegistryStatus_PENDING {
				return sdkerrors.Wrapf(types.ErrAlreadyVote, "schema %s has already voted", srcNftSchemaCode)
			}
			break
		}
	}

	if registryIndex == -1 {
		return sdkerrors.Wrapf(types.ErrSchemaNotInRegistry, "schema %s not found in registry", srcNftSchemaCode)
	}

	// Update vote
	virtualSchemaProposal.VirtualSchema.Registry[registryIndex].Decision = option

	k.SetVirtualSchemaProposal(ctx, virtualSchemaProposal)

	return nil
}
