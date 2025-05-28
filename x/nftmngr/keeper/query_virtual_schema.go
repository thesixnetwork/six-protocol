package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) VirtualSchemaProposalAll(c context.Context, req *types.QueryAllVirtualSchemaProposalRequest) (*types.QueryAllVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var virtualSchemaProposals []types.VirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
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

func (k Keeper) VirtualSchemaAll(c context.Context, req *types.QueryAllVirtualSchemaRequest) (*types.QueryAllVirtualSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var virSchemas []types.VirtualSchema
	ctx := sdk.UnwrapSDKContext(c)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualSchemaKeyPrefix))

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var virSchema types.VirtualSchema
		if err := k.cdc.Unmarshal(value, &virSchema); err != nil {
			return err
		}

		virSchemas = append(virSchemas, virSchema)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVirtualSchemaResponse{VirtualSchema: virSchemas, Pagination: pageRes}, nil
}

func (k Keeper) VirtualSchema(c context.Context, req *types.QueryGetVirtualSchemaRequest) (*types.QueryGetVirtualSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVirtualSchema(
		ctx,
		req.NftSchemaCode,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVirtualSchemaResponse{VirtualSchema: val}, nil
}
