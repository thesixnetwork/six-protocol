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

func (k Keeper) InactiveVirtualSchemaProposalAll(c context.Context, req *types.QueryAllInactiveVirtualSchemaProposalRequest) (*types.QueryAllInactiveVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var inactiveVirtualSchemaProposals []types.InactiveVirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	inactiveVirtualSchemaProposalStore := prefix.NewStore(store, types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(inactiveVirtualSchemaProposalStore, req.Pagination, func(key []byte, value []byte) error {
		var inactiveVirtualSchemaProposal types.InactiveVirtualSchemaProposal
		if err := k.cdc.Unmarshal(value, &inactiveVirtualSchemaProposal); err != nil {
			return err
		}

		inactiveVirtualSchemaProposals = append(inactiveVirtualSchemaProposals, inactiveVirtualSchemaProposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInactiveVirtualSchemaProposalResponse{InactiveVirtualSchemaProposal: inactiveVirtualSchemaProposals, Pagination: pageRes}, nil
}

func (k Keeper) InactiveVirtualSchemaProposal(c context.Context, req *types.QueryGetInactiveVirtualSchemaProposalRequest) (*types.QueryGetInactiveVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetInactiveVirtualSchemaProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetInactiveVirtualSchemaProposalResponse{InactiveVirtualSchemaProposal: val}, nil
}
