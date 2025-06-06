package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) ActionRequestAll(c context.Context, req *types.QueryAllActionRequestRequest) (*types.QueryAllActionRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionRequests []types.ActionOracleRequest
	ctx := sdk.UnwrapSDKContext(c)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	actionRequestStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionRequestKey))

	pageRes, err := query.Paginate(actionRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var actionRequest types.ActionOracleRequest
		if err := k.cdc.Unmarshal(value, &actionRequest); err != nil {
			return err
		}

		actionRequests = append(actionRequests, actionRequest)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActionRequestResponse{ActionOracleRequest: actionRequests, Pagination: pageRes}, nil
}

func (k Keeper) ActionOracleRequest(c context.Context, req *types.QueryGetActionRequestRequest) (*types.QueryGetActionRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	actionRequest, found := k.GetActionRequest(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetActionRequestResponse{ActionOracleRequest: actionRequest}, nil
}
