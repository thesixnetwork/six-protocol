package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBindedSigner(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.BindedSigner {
	items := make([]types.BindedSigner, n)
	for i := range items {
		items[i].OwnerAddress = strconv.Itoa(i)

		keeper.SetBindedSigner(ctx, items[i])
	}
	return items
}

func TestBindedSignerGet(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNBindedSigner(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBindedSigner(ctx,
			item.OwnerAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestBindedSignerRemove(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNBindedSigner(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBindedSigner(ctx,
			item.OwnerAddress,
		)
		_, found := keeper.GetBindedSigner(ctx,
			item.OwnerAddress,
		)
		require.False(t, found)
	}
}

func TestBindedSignerGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNBindedSigner(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBindedSigner(ctx)),
	)
}
