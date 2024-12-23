package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

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
