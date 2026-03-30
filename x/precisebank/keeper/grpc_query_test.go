package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/precisebank/types"
)

// ---------------------------------------------------------------------------
// Remainder
// ---------------------------------------------------------------------------

func TestRemainder_NilRequest(t *testing.T) {
	k, ctx := keepertest.PrecisebankKeeper(t)

	_, err := k.Remainder(ctx, nil)
	require.Error(t, err)

	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}

func TestRemainder_Default(t *testing.T) {
	k, ctx := keepertest.PrecisebankKeeper(t)

	res, err := k.Remainder(ctx, &types.QueryRemainderRequest{})
	require.NoError(t, err)
	require.Equal(t, "0", res.Remainder)
}

func TestRemainder_WithValue(t *testing.T) {
	k, ctx := keepertest.PrecisebankKeeper(t)

	expected := sdkmath.NewInt(500_000_000_000)
	k.SetRemainderAmount(ctx, expected)

	res, err := k.Remainder(ctx, &types.QueryRemainderRequest{})
	require.NoError(t, err)
	require.Equal(t, expected.String(), res.Remainder)
}

// ---------------------------------------------------------------------------
// FractionalBalance
// ---------------------------------------------------------------------------

func TestFractionalBalance_NilRequest(t *testing.T) {
	k, ctx := keepertest.PrecisebankKeeper(t)

	_, err := k.FractionalBalance(ctx, nil)
	require.Error(t, err)

	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}

func TestFractionalBalance_InvalidAddress(t *testing.T) {
	k, ctx := keepertest.PrecisebankKeeper(t)

	_, err := k.FractionalBalance(ctx, &types.QueryFractionalBalanceRequest{Address: "not-an-address"})
	require.Error(t, err)

	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}

func TestFractionalBalance_Zero(t *testing.T) {
	k, ctx := keepertest.PrecisebankKeeper(t)

	addr := sdk.AccAddress([]byte("testaddress1________"))
	res, err := k.FractionalBalance(ctx, &types.QueryFractionalBalanceRequest{Address: addr.String()})
	require.NoError(t, err)
	require.Equal(t, "0", res.FractionalBalance)
}

func TestFractionalBalance_WithValue(t *testing.T) {
	k, ctx := keepertest.PrecisebankKeeper(t)

	addr := sdk.AccAddress([]byte("testaddress1________"))
	expected := sdkmath.NewInt(123_456_789)
	k.SetFractionalBalance(ctx, addr, expected)

	res, err := k.FractionalBalance(ctx, &types.QueryFractionalBalanceRequest{Address: addr.String()})
	require.NoError(t, err)
	require.Equal(t, expected.String(), res.FractionalBalance)
}

// ---------------------------------------------------------------------------
// FractionalBalances
// ---------------------------------------------------------------------------

func TestFractionalBalances_NilRequest(t *testing.T) {
	k, ctx := keepertest.PrecisebankKeeper(t)

	_, err := k.FractionalBalances(ctx, nil)
	require.Error(t, err)

	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}

func TestFractionalBalances_Empty(t *testing.T) {
	k, ctx := keepertest.PrecisebankKeeper(t)

	res, err := k.FractionalBalances(ctx, &types.QueryFractionalBalancesRequest{})
	require.NoError(t, err)
	require.Empty(t, res.FractionalBalances)
}

func TestFractionalBalances_Multiple(t *testing.T) {
	k, ctx := keepertest.PrecisebankKeeper(t)

	addr1 := sdk.AccAddress([]byte("testaddress1________"))
	addr2 := sdk.AccAddress([]byte("testaddress2________"))

	amt1 := sdkmath.NewInt(100_000_000_000)
	amt2 := sdkmath.NewInt(200_000_000_000)

	k.SetFractionalBalance(ctx, addr1, amt1)
	k.SetFractionalBalance(ctx, addr2, amt2)

	res, err := k.FractionalBalances(ctx, &types.QueryFractionalBalancesRequest{})
	require.NoError(t, err)
	require.Len(t, res.FractionalBalances, 2)

	// Build a map for order-independent comparison
	balanceMap := make(map[string]string, len(res.FractionalBalances))
	for _, entry := range res.FractionalBalances {
		balanceMap[entry.Address] = entry.FractionalBalance
	}

	require.Equal(t, amt1.String(), balanceMap[addr1.String()])
	require.Equal(t, amt2.String(), balanceMap[addr2.String()])
}
