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

func createNBindedSigner(keeper keeper.Keeper, ctx sdk.Context, n int) []types.BindedSigner {
	items := make([]types.BindedSigner, n)
	for i := range items {
		items[i].OwnerAddress = strconv.Itoa(i)

		keeper.SetBindedSigner(ctx, items[i])
	}
	return items
}

func TestBindedSignerGet(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNBindedSigner(k, ctx, 10)
	for _, item := range items {
		got, found := k.GetBindedSigner(ctx,
			item.OwnerAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestBindedSignerRemove(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNBindedSigner(k, ctx, 10)
	for _, item := range items {
		k.RemoveBindedSigner(ctx,
			item.OwnerAddress,
		)
		_, found := k.GetBindedSigner(ctx,
			item.OwnerAddress,
		)
		require.False(t, found)
	}
}

func TestBindedSignerGetAll(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNBindedSigner(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllBindedSigner(ctx)),
	)
}

func TestRemoveSignerFromBindedSignerList(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	expiredAt := time.Unix(100, 0).UTC()
	bindedSigner := types.BindedSigner{
		OwnerAddress: "owner",
		Signers: []*types.BindedSignerParams{
			{ActorAddress: "actor1", ExpiredAt: expiredAt},
			{ActorAddress: "actor2", ExpiredAt: expiredAt},
		},
		ActorCount: 2,
	}
	k.SetBindedSigner(ctx, bindedSigner)

	// removing an unknown owner is a no-op
	k.RemoveSignerFromBindedSignerList(ctx, "unknown_owner", "actor1")

	k.RemoveSignerFromBindedSignerList(ctx, "owner", "actor1")
	got, found := k.GetBindedSigner(ctx, "owner")
	require.True(t, found)
	require.Len(t, got.Signers, 1)
	require.Equal(t, "actor2", got.Signers[0].ActorAddress)

	// removing an unknown signer keeps the list untouched
	k.RemoveSignerFromBindedSignerList(ctx, "owner", "unknown_actor")
	got, found = k.GetBindedSigner(ctx, "owner")
	require.True(t, found)
	require.Len(t, got.Signers, 1)
}
