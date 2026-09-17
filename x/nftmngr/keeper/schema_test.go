package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNFTSchema(keeper keeper.Keeper, ctx sdk.Context, n int) []types.NFTSchema {
	items := make([]types.NFTSchema, n)
	for i := range items {
		items[i].Code = strconv.Itoa(i)
		items[i].Name = strconv.Itoa(i)
		items[i].Owner = strconv.Itoa(i)
		items[i].OnchainData = &types.OnChainData{}
		keeper.SetNFTSchema(ctx, items[i])
	}
	return items
}

func TestNFTSchemaGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchema(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNFTSchema(ctx,
			item.Code,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestNFTSchemaRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchema(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNFTSchema(ctx,
			item.Code,
		)
		_, found := keeper.GetNFTSchema(ctx,
			item.Code,
		)
		require.False(t, found)
	}
}

func TestNFTSchemaGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchema(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNFTSchema(ctx)),
	)
}

func createNNFTSchemaByContract(keeper keeper.Keeper, ctx sdk.Context, n int) []types.NFTSchemaByContract {
	items := make([]types.NFTSchemaByContract, n)
	for i := range items {
		items[i].OriginContractAddress = strconv.Itoa(i)
		keeper.SetNFTSchemaByContract(ctx, items[i])
	}
	return items
}

func TestNFTSchemaByContractGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchemaByContract(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNFTSchemaByContract(ctx,
			item.OriginContractAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestNFTSchemaByContractRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchemaByContract(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNFTSchemaByContract(ctx,
			item.OriginContractAddress,
		)
		_, found := keeper.GetNFTSchemaByContract(ctx,
			item.OriginContractAddress,
		)
		require.False(t, found)
	}
}

func TestNFTSchemaByContractGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchemaByContract(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNFTSchemaByContract(ctx)),
	)
}
