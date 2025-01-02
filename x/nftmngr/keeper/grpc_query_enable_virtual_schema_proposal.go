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

func (k Keeper) EnableVirtualSchemaProposalAll(c context.Context, req *types.QueryAllEnableVirtualSchemaProposalRequest) (*types.QueryAllEnableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var enableVirtualSchemaProposals []types.EnableVirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	enableVirtualSchemaProposalStore := prefix.NewStore(store, types.KeyPrefix(types.EnableVirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(enableVirtualSchemaProposalStore, req.Pagination, func(key []byte, value []byte) error {
		var enableVirtualSchemaProposal types.EnableVirtualSchemaProposal
		if err := k.cdc.Unmarshal(value, &enableVirtualSchemaProposal); err != nil {
			return err
		}

		enableVirtualSchemaProposals = append(enableVirtualSchemaProposals, enableVirtualSchemaProposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEnableVirtualSchemaProposalResponse{EnableVirtualSchemaProposal: enableVirtualSchemaProposals, Pagination: pageRes}, nil
}

func (k Keeper) EnableVirtualSchemaProposal(c context.Context, req *types.QueryGetEnableVirtualSchemaProposalRequest) (*types.QueryGetEnableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEnableVirtualSchemaProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEnableVirtualSchemaProposalResponse{EnableVirtualSchemaProposal: val}, nil
}
