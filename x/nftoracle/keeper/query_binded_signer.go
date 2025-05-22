package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

func (k Keeper) BindedSigner(c context.Context, req *types.QueryGetBindedSignerRequest) (*types.QueryGetBindedSignerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetBindedSigner(
		ctx,
		req.OwnerAddress,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBindedSignerResponse{BindedSigner: val}, nil
}
