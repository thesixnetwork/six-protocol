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

func (k Keeper) NftDataAll(c context.Context, req *types.QueryAllNftDataRequest) (*types.QueryAllNftDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var listNFTData []types.NftData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nftDataStore := prefix.NewStore(store, types.KeyPrefix(types.NftDataKeyPrefix))

	pageRes, err := query.Paginate(nftDataStore, req.Pagination, func(key []byte, value []byte) error {
		var nftData types.NftData
		if err := k.cdc.Unmarshal(value, &nftData); err != nil {
			return err
		}

		if req.WithGlobal {
			updateddata := k.appendDataWithSchemaAttributes(ctx, nftData)
			listNFTData = append(listNFTData, updateddata)
		} else {
			listNFTData = append(listNFTData, nftData)
		}

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNftDataResponse{NftData: listNFTData, Pagination: pageRes}, nil
}

func (k Keeper) NftData(c context.Context, req *types.QueryGetNftDataRequest) (*types.QueryGetNftDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNftData(
		ctx,
		req.NftSchemaCode,
		req.TokenId,
	)

	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	updateddata := val
	if req.WithGlobal {
		updateddata = k.appendDataWithSchemaAttributes(ctx, val)
	}

	return &types.QueryGetNftDataResponse{NftData: updateddata}, nil
}

func (k Keeper) NftCollection(c context.Context, req *types.QueryGetNftCollectionRequest) (*types.QueryGetNftCollectionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var metadatas []*types.NftData
	ctx := sdk.UnwrapSDKContext(c)

	// store := ctx.KVStore(k.storeKey)
	// NftcollectionStore := prefix.NewStore(store, types.KeyPrefix(types.NftCollectionKeyPrefix))
	NftcollectionStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.CollectionkeyPrefix(req.NftSchemaCode))

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
