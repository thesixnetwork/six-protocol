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

func (k Keeper) ActiveEnableVirtualSchemaProposalAll(c context.Context, req *types.QueryAllActiveEnableVirtualSchemaProposalRequest) (*types.QueryAllActiveEnableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var activeEnableVirtualSchemaProposals []types.ActiveEnableVirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	activeEnableVirtualSchemaProposalStore := prefix.NewStore(store, types.KeyPrefix(types.ActiveEnableVirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(activeEnableVirtualSchemaProposalStore, req.Pagination, func(key []byte, value []byte) error {
		var activeEnableVirtualSchemaProposal types.ActiveEnableVirtualSchemaProposal
		if err := k.cdc.Unmarshal(value, &activeEnableVirtualSchemaProposal); err != nil {
			return err
		}

		activeEnableVirtualSchemaProposals = append(activeEnableVirtualSchemaProposals, activeEnableVirtualSchemaProposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActiveEnableVirtualSchemaProposalResponse{ActiveEnableVirtualSchemaProposal: activeEnableVirtualSchemaProposals, Pagination: pageRes}, nil
}

func (k Keeper) ActiveEnableVirtualSchemaProposal(c context.Context, req *types.QueryGetActiveEnableVirtualSchemaProposalRequest) (*types.QueryGetActiveEnableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetActiveEnableVirtualSchemaProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetActiveEnableVirtualSchemaProposalResponse{ActiveEnableVirtualSchemaProposal: val}, nil
}
