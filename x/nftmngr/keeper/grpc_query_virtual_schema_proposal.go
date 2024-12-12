package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VirtualSchemaProposalAll(c context.Context, req *types.QueryAllVirtualSchemaProposalRequest) (*types.QueryAllVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var virtualSchemaProposals []types.VirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	virtualSchemaProposalStore := prefix.NewStore(store, types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(virtualSchemaProposalStore, req.Pagination, func(key []byte, value []byte) error {
		var virtualSchemaProposal types.VirtualSchemaProposal
		if err := k.cdc.Unmarshal(value, &virtualSchemaProposal); err != nil {
			return err
		}

		virtualSchemaProposals = append(virtualSchemaProposals, virtualSchemaProposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVirtualSchemaProposalResponse{VirtualSchemaProposal: virtualSchemaProposals, Pagination: pageRes}, nil
}

func (k Keeper) VirtualSchemaProposal(c context.Context, req *types.QueryGetVirtualSchemaProposalRequest) (*types.QueryGetVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVirtualSchemaProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVirtualSchemaProposalResponse{VirtualSchemaProposal: val}, nil
}
