package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

func TestNFTFeeBalance(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)

	_, found := keeper.GetNFTFeeBalance(ctx)
	require.False(t, found)

	item := types.NFTFeeBalance{
		FeeBalances: []string{"0usix"},
	}
	keeper.SetNFTFeeBalance(ctx, item)

	rst, found := keeper.GetNFTFeeBalance(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)

	keeper.RemoveNFTFeeBalance(ctx)
	_, found = keeper.GetNFTFeeBalance(ctx)
	require.False(t, found)
}

func TestNFTFeeConfig(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)

	_, found := keeper.GetNFTFeeConfig(ctx)
	require.False(t, found)

	item := types.NFTFeeConfig{
		SchemaFee: &types.FeeConfig{
			FeeAmount: "200000000usix",
		},
	}
	keeper.SetNFTFeeConfig(ctx, item)

	rst, found := keeper.GetNFTFeeConfig(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)

	keeper.RemoveNFTFeeConfig(ctx)
	_, found = keeper.GetNFTFeeConfig(ctx)
	require.False(t, found)
}
