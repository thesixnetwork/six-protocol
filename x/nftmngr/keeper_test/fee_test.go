package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func createTestNFTFeeBalance(keeper *keeper.Keeper, ctx sdk.Context) types.NFTFeeBalance {
	item := types.NFTFeeBalance{}
	keeper.SetNFTFeeBalance(ctx, item)
	return item
}

func TestNFTFeeBalanceGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	item := createTestNFTFeeBalance(keeper, ctx)
	rst, found := keeper.GetNFTFeeBalance(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestNFTFeeBalanceRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	createTestNFTFeeBalance(keeper, ctx)
	keeper.RemoveNFTFeeBalance(ctx)
	_, found := keeper.GetNFTFeeBalance(ctx)
	require.False(t, found)
}

func createTestNFTFeeConfig(keeper *keeper.Keeper, ctx sdk.Context) types.NFTFeeConfig {
	item := types.NFTFeeConfig{}
	keeper.SetNFTFeeConfig(ctx, item)
	return item
}

func TestNFTFeeConfigGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	item := createTestNFTFeeConfig(keeper, ctx)
	rst, found := keeper.GetNFTFeeConfig(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestNFTFeeConfigRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	createTestNFTFeeConfig(keeper, ctx)
	keeper.RemoveNFTFeeConfig(ctx)
	_, found := keeper.GetNFTFeeConfig(ctx)
	require.False(t, found)
}
