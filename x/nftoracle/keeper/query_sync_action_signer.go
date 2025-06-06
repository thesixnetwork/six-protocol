package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) SyncActionSignerAll(c context.Context, req *types.QueryAllSyncActionSignerRequest) (*types.QueryAllSyncActionSignerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var syncActionSigners []types.SyncActionSigner
	ctx := sdk.UnwrapSDKContext(c)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	syncActionSignerStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SyncActionSignerKey))

	pageRes, err := query.Paginate(syncActionSignerStore, req.Pagination, func(key []byte, value []byte) error {
		var syncActionSigner types.SyncActionSigner
		if err := k.cdc.Unmarshal(value, &syncActionSigner); err != nil {
			return err
		}

		syncActionSigners = append(syncActionSigners, syncActionSigner)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSyncActionSignerResponse{SyncActionSigner: syncActionSigners, Pagination: pageRes}, nil
}

func (k Keeper) SyncActionSigner(c context.Context, req *types.QueryGetSyncActionSignerRequest) (*types.QueryGetSyncActionSignerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	syncActionSigner, found := k.GetSyncActionSigner(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSyncActionSignerResponse{SyncActionSigner: syncActionSigner}, nil
}
