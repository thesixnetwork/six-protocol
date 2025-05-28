package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) TokenBurnAll(ctx context.Context, req *types.QueryAllTokenBurnRequest) (*types.QueryAllTokenBurnResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokenBurns []types.TokenBurn

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	tokenBurnStore := prefix.NewStore(store, types.KeyPrefix(types.TokenBurnKeyPrefix))

	pageRes, err := query.Paginate(tokenBurnStore, req.Pagination, func(key []byte, value []byte) error {
		var tokenBurn types.TokenBurn
		if err := k.cdc.Unmarshal(value, &tokenBurn); err != nil {
			return err
		}

		tokenBurns = append(tokenBurns, tokenBurn)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTokenBurnResponse{TokenBurn: tokenBurns, Pagination: pageRes}, nil
}

func (k Keeper) TokenBurn(ctx context.Context, req *types.QueryGetTokenBurnRequest) (*types.QueryGetTokenBurnResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetTokenBurn(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTokenBurnResponse{TokenBurn: val}, nil
}
