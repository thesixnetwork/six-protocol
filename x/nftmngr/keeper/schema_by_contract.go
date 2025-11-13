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

func (k Keeper) NFTSchemaByContract(c context.Context, req *types.QueryGetNFTSchemaByContractRequest) (*types.QueryGetNFTSchemaByContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNFTSchemaByContract(
		ctx,
		req.OriginContractAddress,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNFTSchemaByContractResponse{NFTSchemaByContract: val}, nil
}

func (k Keeper) NFTSchemaByContractAll(c context.Context, req *types.QueryAllNFTSchemaByContractRequest) (*types.QueryAllNFTSchemaByContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nFTSchemaByContracts []types.NFTSchemaByContract
	ctx := sdk.UnwrapSDKContext(c)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTSchemaByContractKeyPrefix))

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var nFTSchemaByContract types.NFTSchemaByContract
		if err := k.cdc.Unmarshal(value, &nFTSchemaByContract); err != nil {
			return err
		}

		nFTSchemaByContracts = append(nFTSchemaByContracts, nFTSchemaByContract)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNFTSchemaByContractResponse{NFTSchemaByContract: nFTSchemaByContracts, Pagination: pageRes}, nil
}
