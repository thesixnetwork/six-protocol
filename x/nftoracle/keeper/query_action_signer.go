package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) ActionSignerAll(c context.Context, req *types.QueryAllActionSignerRequest) (*types.QueryAllActionSignerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionSigners []types.ActionSigner
	ctx := sdk.UnwrapSDKContext(c)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	actionSignerStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionSignerKeyPrefix))

	pageRes, err := query.Paginate(actionSignerStore, req.Pagination, func(key []byte, value []byte) error {
		var actionSigner types.ActionSigner
		if err := k.cdc.Unmarshal(value, &actionSigner); err != nil {
			return err
		}

		actionSigners = append(actionSigners, actionSigner)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActionSignerResponse{ActionSigner: actionSigners, Pagination: pageRes}, nil
}

func (k Keeper) ActionSigner(c context.Context, req *types.QueryGetActionSignerRequest) (*types.QueryGetActionSignerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetActionSigner(
		ctx,
		req.ActorAddress,
		req.OwnerAddress,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetActionSignerResponse{ActionSigner: val}, nil
}
