package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) MintRequestAll(c context.Context, req *types.QueryAllMintRequestRequest) (*types.QueryAllMintRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var mintRequests []types.MintRequest
	ctx := sdk.UnwrapSDKContext(c)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	mintRequestStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintRequestKey))

	pageRes, err := query.Paginate(mintRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var mintRequest types.MintRequest
		if err := k.cdc.Unmarshal(value, &mintRequest); err != nil {
			return err
		}

		mintRequests = append(mintRequests, mintRequest)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMintRequestResponse{MintRequest: mintRequests, Pagination: pageRes}, nil
}

func (k Keeper) MintRequest(c context.Context, req *types.QueryGetMintRequestRequest) (*types.QueryGetMintRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	mintRequest, found := k.GetMintRequest(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMintRequestResponse{MintRequest: mintRequest}, nil
}
