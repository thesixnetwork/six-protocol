package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNLockSchemaFee(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LockSchemaFee {
	items := make([]types.LockSchemaFee, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)

		keeper.SetLockSchemaFee(ctx, items[i])
	}
	return items
}

func TestLockSchemaFeeGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNLockSchemaFee(&keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLockSchemaFee(ctx,
			item.Id,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestLockSchemaFeeRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNLockSchemaFee(&keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLockSchemaFee(ctx,
			item.Id,
		)
		_, found := keeper.GetLockSchemaFee(ctx,
			item.Id,
		)
		require.False(t, found)
	}
}

func TestLockSchemaFeeGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNLockSchemaFee(&keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLockSchemaFee(ctx)),
	)
}
