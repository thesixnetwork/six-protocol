package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k Keeper) MetadataCreatorAll(c context.Context, req *types.QueryAllMetadataCreatorRequest) (*types.QueryAllMetadataCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var metadataCreators []types.MetadataCreator
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	metadataCreatorStore := prefix.NewStore(store, types.KeyPrefix(types.MetadataCreatorKeyPrefix))

	pageRes, err := query.Paginate(metadataCreatorStore, req.Pagination, func(key []byte, value []byte) error {
		var metadataCreator types.MetadataCreator
		if err := k.cdc.Unmarshal(value, &metadataCreator); err != nil {
			return err
		}

		metadataCreators = append(metadataCreators, metadataCreator)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMetadataCreatorResponse{MetadataCreator: metadataCreators, Pagination: pageRes}, nil
}

func (k Keeper) MetadataCreator(c context.Context, req *types.QueryGetMetadataCreatorRequest) (*types.QueryGetMetadataCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMetadataCreator(
		ctx,
		req.NftSchemaCode,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMetadataCreatorResponse{MetadataCreator: val}, nil
}
