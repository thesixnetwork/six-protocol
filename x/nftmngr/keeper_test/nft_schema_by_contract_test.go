package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNFTSchemaByContract(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NFTSchemaByContract {
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
