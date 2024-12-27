package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func TestNFTFeeBalanceQuery(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestNFTFeeBalance(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNFTFeeBalanceRequest
		response *types.QueryGetNFTFeeBalanceResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetNFTFeeBalanceRequest{},
			response: &types.QueryGetNFTFeeBalanceResponse{NFTFeeBalance: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NFTFeeBalance(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestNFTFeeConfigQuery(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestNFTFeeConfig(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNFTFeeConfigRequest
		response *types.QueryGetNFTFeeConfigResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetNFTFeeConfigRequest{},
			response: &types.QueryGetNFTFeeConfigResponse{NFTFeeConfig: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NFTFeeConfig(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
