package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
