package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTokenBurn(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TokenBurn {
	items := make([]types.TokenBurn, n)
	for i := range items {
		items[i].Token = strconv.Itoa(i)

		keeper.SetTokenBurn(ctx, items[i])
	}
	return items
}

func TestTokenBurnGet(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	items := createNTokenBurn(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTokenBurn(ctx,
			item.Token,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTokenBurnRemove(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	items := createNTokenBurn(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTokenBurn(ctx,
			item.Token,
		)
		_, found := keeper.GetTokenBurn(ctx,
			item.Token,
		)
		require.False(t, found)
	}
}

func TestTokenBurnGetAll(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	items := createNTokenBurn(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTokenBurn(ctx)),
	)
}
