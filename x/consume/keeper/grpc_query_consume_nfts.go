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

func (k Keeper) ConsumeNfts(c context.Context, req *types.QueryConsumeNftsRequest) (*types.QueryConsumeNftsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Define a variable that will store a list of use_nfts
	var use_nfts []*types.UseNft
	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(c)
	// Get the key-value module store using the store key (in our case store key is "chain")
	store := ctx.KVStore(k.storeKey)
	// Get the part of the store that keeps use_nfts (using use_nft key, which is "UseNft-value-")
	use_nftStore := prefix.NewStore(store, []byte(types.UseNftKey))
	// Paginate the use_nfts store based on PageRequest
	pageRes, err := query.Paginate(use_nftStore, req.Pagination, func(key []byte, value []byte) error {
		var use_nft types.UseNft
		if err := k.cdc.Unmarshal(value, &use_nft); err != nil {
			return err
		}
		use_nfts = append(use_nfts, &use_nft)
		return nil
	})
	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// Return a struct containing a list of use_nfts and pagination info

	return &types.QueryConsumeNftsResponse{UseNft: use_nfts, Pagination: pageRes}, nil
}
