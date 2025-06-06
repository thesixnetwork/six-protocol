package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) NftDataAll(ctx context.Context, req *types.QueryAllNftDataRequest) (*types.QueryAllNftDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nftDatas []types.NftData

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	nftDataStore := prefix.NewStore(store, types.KeyPrefix(types.NftDataKeyPrefix))

	pageRes, err := query.Paginate(nftDataStore, req.Pagination, func(key []byte, value []byte) error {
		var nftData types.NftData
		if err := k.cdc.Unmarshal(value, &nftData); err != nil {
			return err
		}

		nftDatas = append(nftDatas, nftData)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNftDataResponse{NftData: nftDatas, Pagination: pageRes}, nil
}

func (k Keeper) NftData(ctx context.Context, req *types.QueryGetNftDataRequest) (*types.QueryGetNftDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetNftData(
		ctx,
		req.NftSchemaCode,
		req.TokenId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNftDataResponse{NftData: val}, nil
}
