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

func (k Keeper) ActiveVirtualSchemaProposalAll(c context.Context, req *types.QueryAllActiveVirtualSchemaProposalRequest) (*types.QueryAllActiveVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var activeVirtualSchemaProposals []types.ActiveVirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	activeVirtualSchemaProposalStore := prefix.NewStore(store, types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(activeVirtualSchemaProposalStore, req.Pagination, func(key []byte, value []byte) error {
		var activeVirtualSchemaProposal types.ActiveVirtualSchemaProposal
		if err := k.cdc.Unmarshal(value, &activeVirtualSchemaProposal); err != nil {
			return err
		}

		activeVirtualSchemaProposals = append(activeVirtualSchemaProposals, activeVirtualSchemaProposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActiveVirtualSchemaProposalResponse{ActiveVirtualSchemaProposal: activeVirtualSchemaProposals, Pagination: pageRes}, nil
}

func (k Keeper) ActiveVirtualSchemaProposal(c context.Context, req *types.QueryGetActiveVirtualSchemaProposalRequest) (*types.QueryGetActiveVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetActiveVirtualSchemaProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetActiveVirtualSchemaProposalResponse{ActiveVirtualSchemaProposal: val}, nil
}
