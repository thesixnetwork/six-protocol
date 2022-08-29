package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/thesixnetwork/six-protocol/x/consume/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NftUsedAll(c context.Context, req *types.QueryAllNftUsedRequest) (*types.QueryAllNftUsedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nftUseds []types.NftUsed
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nftUsedStore := prefix.NewStore(store, types.KeyPrefix(types.NftUsedKeyPrefix))

	pageRes, err := query.Paginate(nftUsedStore, req.Pagination, func(key []byte, value []byte) error {
		var nftUsed types.NftUsed
		if err := k.cdc.Unmarshal(value, &nftUsed); err != nil {
			return err
		}

		nftUseds = append(nftUseds, nftUsed)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNftUsedResponse{NftUsed: nftUseds, Pagination: pageRes}, nil
}

func (k Keeper) NftUsed(c context.Context, req *types.QueryGetNftUsedRequest) (*types.QueryGetNftUsedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNftUsed(
		ctx,
		req.Token,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNftUsedResponse{NftUsed: val}, nil
}
