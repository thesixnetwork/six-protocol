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

func (k Keeper) MintpermAll(ctx context.Context, req *types.QueryAllMintpermRequest) (*types.QueryAllMintpermResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var mintperms []types.Mintperm

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	mintpermStore := prefix.NewStore(store, types.KeyPrefix(types.MintpermKeyPrefix))

	pageRes, err := query.Paginate(mintpermStore, req.Pagination, func(key []byte, value []byte) error {
		var mintperm types.Mintperm
		if err := k.cdc.Unmarshal(value, &mintperm); err != nil {
			return err
		}

		mintperms = append(mintperms, mintperm)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMintpermResponse{Mintperm: mintperms, Pagination: pageRes}, nil
}

func (k Keeper) Mintperm(ctx context.Context, req *types.QueryGetMintpermRequest) (*types.QueryGetMintpermResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetMintperm(
		ctx,
		req.Token,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMintpermResponse{Mintperm: val}, nil
}
