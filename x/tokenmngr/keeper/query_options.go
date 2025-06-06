package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Options(goCtx context.Context, req *types.QueryGetOptionsRequest) (*types.QueryGetOptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetOptions(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOptionsResponse{Options: val}, nil
}
