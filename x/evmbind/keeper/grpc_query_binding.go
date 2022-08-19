package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/thesixnetwork/six-protocol/x/evmbind/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BindingAll(c context.Context, req *types.QueryAllBindingRequest) (*types.QueryAllBindingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var bindings []types.Binding
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	bindingStore := prefix.NewStore(store, types.KeyPrefix(types.BindingKeyPrefix))

	pageRes, err := query.Paginate(bindingStore, req.Pagination, func(key []byte, value []byte) error {
		var binding types.Binding
		if err := k.cdc.Unmarshal(value, &binding); err != nil {
			return err
		}

		bindings = append(bindings, binding)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBindingResponse{Binding: bindings, Pagination: pageRes}, nil
}

func (k Keeper) Binding(c context.Context, req *types.QueryGetBindingRequest) (*types.QueryGetBindingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetBinding(
		ctx,
		req.EthAddress,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBindingResponse{Binding: val}, nil
}
