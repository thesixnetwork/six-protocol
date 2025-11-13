package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) VirtualActionAll(c context.Context, req *types.QueryAllVirtualActionRequest) (*types.QueryAllVirtualActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var virtuals []types.VirtualAction
	ctx := sdk.UnwrapSDKContext(c)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualActionKeyPrefix))

	// chekc if input specify schemaCode
	if req.NftSchemaCode == "" {
		pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
			var virtual types.VirtualAction
			if err := k.cdc.Unmarshal(value, &virtual); err != nil {
				return err
			}

			virtuals = append(virtuals, virtual)
			return nil
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return &types.QueryAllVirtualActionResponse{VirtualAction: virtuals, Pagination: pageRes}, nil
	} else {
		pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
			var virtual types.VirtualAction
			if err := k.cdc.Unmarshal(value, &virtual); err != nil {
				return err
			}

			if virtual.VirtualNftSchemaCode == req.NftSchemaCode {
				virtuals = append(virtuals, virtual)
			}

			return nil
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		pageRes.Total = uint64(len(virtuals))
		pageRes.NextKey = nil
		return &types.QueryAllVirtualActionResponse{VirtualAction: virtuals, Pagination: pageRes}, nil
	}
}

func (k Keeper) VirtualAction(c context.Context, req *types.QueryGetVirtualActionRequest) (*types.QueryGetVirtualActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVirtualAction(
		ctx,
		req.NftSchemaCode,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVirtualActionResponse{VirtualAction: val}, nil
}
