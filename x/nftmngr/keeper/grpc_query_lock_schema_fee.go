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

func (k Keeper) LockSchemaFeeAll(c context.Context, req *types.QueryAllLockSchemaFeeRequest) (*types.QueryAllLockSchemaFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var lockSchemaFees []types.LockSchemaFee
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	lockSchemaFeeStore := prefix.NewStore(store, types.KeyPrefix(types.LockSchemaFeeKeyPrefix))

	pageRes, err := query.Paginate(lockSchemaFeeStore, req.Pagination, func(key []byte, value []byte) error {
		var lockSchemaFee types.LockSchemaFee
		if err := k.cdc.Unmarshal(value, &lockSchemaFee); err != nil {
			return err
		}

		lockSchemaFees = append(lockSchemaFees, lockSchemaFee)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLockSchemaFeeResponse{LockSchemaFee: lockSchemaFees, Pagination: pageRes}, nil
}

func (k Keeper) LockSchemaFee(c context.Context, req *types.QueryGetLockSchemaFeeRequest) (*types.QueryGetLockSchemaFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLockSchemaFee(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLockSchemaFeeResponse{LockSchemaFee: val}, nil
}
