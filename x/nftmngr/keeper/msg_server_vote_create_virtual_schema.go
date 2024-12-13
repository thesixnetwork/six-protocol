package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// TODO:: Feat(VirtualSchema)
// 1. Check if virtual schema is exist
// 2. Check if virtual schema is enable
// 3. Check if voter(creator) is owner of some src schema
// 4. Check if src schema already vote
func (k msgServer) VoteCreateVirtualSchema(goCtx context.Context, msg *types.MsgVoteCreateVirtualSchema) (*types.MsgVoteCreateVirtualSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	virtualSchemaProposal, found := k.GetVirtualSchemaProposal(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.Id)
	}

	isOwner := false

	// Track if the vote has been processed
	voteProcessed := false

	// loop to find then schema registry
	for i, registry := range virtualSchemaProposal.Registry {
		srcSchema, found := k.GetNFTSchema(ctx, registry.NftSchemaCode)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, registry.NftSchemaCode)
		}

		if srcSchema.Owner == msg.Creator {
			isOwner = true
		}

		// Check if already voted
		if isOwner && registry.Status != types.RegistryStatus_PENDING {
			return nil, sdkerrors.Wrap(types.ErrAlreadyVote, msg.Creator)
		}

		// Update the status
		virtualSchemaProposal.Registry[i].Status = msg.Option
		voteProcessed = true
	}

	// Ensure the vote was processed
	if !voteProcessed {
		return nil, sdkerrors.Wrap(types.ErrSchemaNotInRegistry, msg.Creator)
	}

	if !isOwner {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	// Count votes
	var (
		acceptCount  int
		totalVotes   int
		voteTreshold = len(virtualSchemaProposal.Registry)
	)

	for _, registry := range virtualSchemaProposal.Registry {
		if registry.Status == types.RegistryStatus_ACCEPT {
			acceptCount++
		}
		if registry.Status != types.RegistryStatus_PENDING {
			totalVotes++
		}
	}

	// Check if all votes are in
	if totalVotes == voteTreshold {
		k.AfterProposalSuccess(ctx, virtualSchemaProposal.Id)
	}

	// save
	k.SetVirtualSchemaProposal(ctx, virtualSchemaProposal)

	return &types.MsgVoteCreateVirtualSchemaResponse{}, nil
}
