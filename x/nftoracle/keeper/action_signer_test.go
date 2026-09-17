package keeper_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNActionSigner(keeper keeper.Keeper, ctx sdk.Context, n int) []types.ActionSigner {
	items := make([]types.ActionSigner, n)
	for i := range items {
		items[i].ActorAddress = strconv.Itoa(i)
		items[i].OwnerAddress = strconv.Itoa(i)
		items[i].CreatedAt = time.Unix(int64(i), 0).UTC()
		items[i].ExpiredAt = time.Unix(int64(i+100), 0).UTC()

		keeper.SetActionSigner(ctx, items[i])
	}
	return items
}

func TestActionSignerGet(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSigner(k, ctx, 10)
	for _, item := range items {
		got, found := k.GetActionSigner(ctx,
			item.ActorAddress,
			item.OwnerAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestActionSignerRemove(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSigner(k, ctx, 10)
	for _, item := range items {
		k.RemoveActionSigner(ctx,
			item.ActorAddress,
			item.OwnerAddress,
		)
		_, found := k.GetActionSigner(ctx,
			item.ActorAddress,
			item.OwnerAddress,
		)
		require.False(t, found)
	}
}

func TestActionSignerGetAll(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSigner(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllActionSigner(ctx)),
	)
}
