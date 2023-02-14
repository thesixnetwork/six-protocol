package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TokenBurnAll(c context.Context, req *types.QueryAllTokenBurnRequest) (*types.QueryAllTokenBurnResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokenBurns []types.TokenBurn
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	tokenBurnStore := prefix.NewStore(store, types.KeyPrefix(types.TokenBurnKeyPrefix))
	// TODO:: Check paginate limit
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

func (k Keeper) TokenBurnAllV202(c context.Context, req *types.QueryAllTokenBurnRequest) (*types.QueryAllTokenBurnResponseV202, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokenBurns []types.TokenBurnV202
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	tokenBurnStore := prefix.NewStore(store, types.KeyPrefix(types.TokenBurnKeyPrefix))
	// TODO:: Check paginate limit
	pageRes, err := query.Paginate(tokenBurnStore, req.Pagination, func(key []byte, value []byte) error {
		var tokenBurn types.TokenBurnV202
		if err := k.cdc.Unmarshal(value, &tokenBurn); err != nil {
			return err
		}

		tokenBurns = append(tokenBurns, tokenBurn)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTokenBurnResponseV202{TokenBurn: tokenBurns, Pagination: pageRes}, nil
}

func (k Keeper) TokenBurn(c context.Context, req *types.QueryGetTokenBurnRequest) (*types.QueryGetTokenBurnResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTokenBurn(
		ctx,
		req.Token,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTokenBurnResponse{TokenBurn: val}, nil
}

func (k Keeper) TokenBurnV202(c context.Context, req *types.QueryGetTokenBurnRequest) (*types.QueryGetTokenBurnResponseV202, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTokenBurnV202(
		ctx,
		req.Token,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTokenBurnResponseV202{TokenBurn: val}, nil
}
