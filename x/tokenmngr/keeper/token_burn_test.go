package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createNTokenBurn(keeper keeper.Keeper, ctx context.Context, n int) []types.TokenBurn {
	items := make([]types.TokenBurn, n)
	for i := range items {
		items[i].Amount = sdk.NewCoin("token"+strconv.Itoa(i), sdkmath.NewInt(int64(i)))

		keeper.SetTokenBurn(ctx, items[i])
	}
	return items
}

func TestTokenBurnGet(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	items := createNTokenBurn(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTokenBurn(ctx,
			item.Amount.Denom,
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
			item.Amount.Denom,
		)
		_, found := keeper.GetTokenBurn(ctx,
			item.Amount.Denom,
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
