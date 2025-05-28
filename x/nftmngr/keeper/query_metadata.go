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

func (k Keeper) MetadataCreatorAll(c context.Context, req *types.QueryAllMetadataCreatorRequest) (*types.QueryAllMetadataCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var metadataCreators []types.MetadataCreator
	ctx := sdk.UnwrapSDKContext(c)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MetadataCreatorKeyPrefix))

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
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

func (k Keeper) NftCollection(c context.Context, req *types.QueryGetNftCollectionRequest) (*types.QueryGetNftCollectionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var metadatas []*types.NftData
	ctx := sdk.UnwrapSDKContext(c)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	NftcollectionStore := prefix.NewStore(storeAdapter, types.CollectionkeyPrefix(req.NftSchemaCode))

	pageRes, err := query.Paginate(NftcollectionStore, req.Pagination, func(key []byte, value []byte) error {
		var nftData types.NftData
		if err := k.cdc.Unmarshal(value, &nftData); err != nil {
			return err
		}
		metadatas = append(metadatas, &nftData)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetNftCollectionResponse{NftCollection: metadatas, Pagination: pageRes}, nil
}

// function appendDataWithSchemaAttributes is used to append the data with schema attributes coz schema attributes are not stored in nftdata
func (k Keeper) appendDataWithSchemaAttributes(ctx sdk.Context, dataOnToken types.NftData) (updatedData types.NftData) {
	listOfAllschemaAttributeValue := k.GetAllSchemaAttribute(ctx)

	for _, schemaAttribute := range listOfAllschemaAttributeValue {
		if schemaAttribute.NftSchemaCode == dataOnToken.NftSchemaCode {
			scheamAttributeValues := ConverSchemaAttributeToNFTAttributeValue(&schemaAttribute)
			dataOnToken.OnchainAttributes = append(dataOnToken.OnchainAttributes, scheamAttributeValues)
		}
	}

	return dataOnToken
}
