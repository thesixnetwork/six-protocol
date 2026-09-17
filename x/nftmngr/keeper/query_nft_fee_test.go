package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

func TestNFTFeeConfigQuery(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)

	_, err := k.NFTFeeConfig(ctx, nil)
	require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))

	_, err = k.NFTFeeConfig(ctx, &types.QueryGetNFTFeeConfigRequest{})
	require.ErrorIs(t, err, status.Error(codes.NotFound, "not found"))

	item := types.NFTFeeConfig{
		SchemaFee: &types.FeeConfig{
			FeeAmount: "200000000usix",
		},
	}
	k.SetNFTFeeConfig(ctx, item)

	resp, err := k.NFTFeeConfig(ctx, &types.QueryGetNFTFeeConfigRequest{})
	require.NoError(t, err)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&resp.NFTFeeConfig),
	)
}

func TestNFTFeeBalanceQuery(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)

	_, err := k.NFTFeeBalance(ctx, nil)
	require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))

	_, err = k.NFTFeeBalance(ctx, &types.QueryGetNFTFeeBalanceRequest{})
	require.ErrorIs(t, err, status.Error(codes.NotFound, "not found"))

	item := types.NFTFeeBalance{
		FeeBalances: []string{"0usix"},
	}
	k.SetNFTFeeBalance(ctx, item)

	resp, err := k.NFTFeeBalance(ctx, &types.QueryGetNFTFeeBalanceRequest{})
	require.NoError(t, err)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&resp.NFTFeeBalance),
	)
}
