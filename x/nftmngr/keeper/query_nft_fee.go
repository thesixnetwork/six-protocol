package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) NFTFeeConfig(c context.Context, req *types.QueryGetNFTFeeConfigRequest) (*types.QueryGetNFTFeeConfigResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNFTFeeConfig(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNFTFeeConfigResponse{NFTFeeConfig: val}, nil
}

func (k Keeper) NFTFeeBalance(c context.Context, req *types.QueryGetNFTFeeBalanceRequest) (*types.QueryGetNFTFeeBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNFTFeeBalance(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNFTFeeBalanceResponse{NFTFeeBalance: val}, nil
}
