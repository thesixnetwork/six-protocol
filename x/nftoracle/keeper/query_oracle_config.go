package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) OracleConfig(c context.Context, req *types.QueryGetOracleConfigRequest) (*types.QueryGetOracleConfigResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetOracleConfig(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOracleConfigResponse{OracleConfig: val}, nil
}
