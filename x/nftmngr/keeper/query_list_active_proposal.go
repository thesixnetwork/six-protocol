package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO:: Add pagination on next release
func (k Keeper) ListActiveProposal(goCtx context.Context, req *types.QueryListActiveProposalRequest) (*types.QueryListActiveProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var virtualSchemaProposals []types.VirtualSchemaProposal
	listActiveProposal := k.GetAllActiveVirtualSchemaProposal(ctx)

	for _, virtualSchemaProposal := range listActiveProposal {
		proposal, found := k.GetVirtualSchemaProposal(ctx, virtualSchemaProposal.Id)
		if !found {
			return nil, status.Error(codes.NotFound, "not found")
		}
		virtualSchemaProposals = append(virtualSchemaProposals, proposal)
	}

	return &types.QueryListActiveProposalResponse{
		VirtualSchemaProposal: virtualSchemaProposals,
	}, nil
}
