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

func (k Keeper) MintpermAll(c context.Context, req *types.QueryAllMintpermRequest) (*types.QueryAllMintpermResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var mintperms []types.Mintperm
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
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

func (k Keeper) Mintperm(c context.Context, req *types.QueryGetMintpermRequest) (*types.QueryGetMintpermResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMintperm(
		ctx,
		req.Token,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetMintpermResponse{Mintperm: val}, nil
}
