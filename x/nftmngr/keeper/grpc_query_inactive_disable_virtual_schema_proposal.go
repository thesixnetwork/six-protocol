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

func (k Keeper) InactiveDisableVirtualSchemaProposalAll(c context.Context, req *types.QueryAllInactiveDisableVirtualSchemaProposalRequest) (*types.QueryAllInactiveDisableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var inactiveDisableVirtualSchemaProposals []types.InactiveDisableVirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	inactiveDisableVirtualSchemaProposalStore := prefix.NewStore(store, types.KeyPrefix(types.InactiveDisableVirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(inactiveDisableVirtualSchemaProposalStore, req.Pagination, func(key []byte, value []byte) error {
		var inactiveDisableVirtualSchemaProposal types.InactiveDisableVirtualSchemaProposal
		if err := k.cdc.Unmarshal(value, &inactiveDisableVirtualSchemaProposal); err != nil {
			return err
		}

		inactiveDisableVirtualSchemaProposals = append(inactiveDisableVirtualSchemaProposals, inactiveDisableVirtualSchemaProposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInactiveDisableVirtualSchemaProposalResponse{InactiveDisableVirtualSchemaProposal: inactiveDisableVirtualSchemaProposals, Pagination: pageRes}, nil
}

func (k Keeper) InactiveDisableVirtualSchemaProposal(c context.Context, req *types.QueryGetInactiveDisableVirtualSchemaProposalRequest) (*types.QueryGetInactiveDisableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetInactiveDisableVirtualSchemaProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetInactiveDisableVirtualSchemaProposalResponse{InactiveDisableVirtualSchemaProposal: val}, nil
}
