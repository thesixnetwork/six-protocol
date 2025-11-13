package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) CollectionOwnerRequestAll(c context.Context, req *types.QueryAllCollectionOwnerRequestRequest) (*types.QueryAllCollectionOwnerRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var collectionOwnerRequests []types.CollectionOwnerRequest
	ctx := sdk.UnwrapSDKContext(c)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	collectionOwnerRequestStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CollectionOwnerRequestKey))

	pageRes, err := query.Paginate(collectionOwnerRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var collectionOwnerRequest types.CollectionOwnerRequest
		if err := k.cdc.Unmarshal(value, &collectionOwnerRequest); err != nil {
			return err
		}

		collectionOwnerRequests = append(collectionOwnerRequests, collectionOwnerRequest)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCollectionOwnerRequestResponse{CollectionOwnerRequest: collectionOwnerRequests, Pagination: pageRes}, nil
}

func (k Keeper) CollectionOwnerRequest(c context.Context, req *types.QueryGetCollectionOwnerRequestRequest) (*types.QueryGetCollectionOwnerRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	collectionOwnerRequest, found := k.GetCollectionOwnerRequest(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetCollectionOwnerRequestResponse{CollectionOwnerRequest: collectionOwnerRequest}, nil
}
