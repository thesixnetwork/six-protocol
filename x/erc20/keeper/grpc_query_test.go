package keeper_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/erc20/types"
)

// TestTokenPairs_Empty verifies that an empty store returns an empty list.
func TestTokenPairs_Empty(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	resp, err := k.TokenPairs(ctx, &types.QueryTokenPairsRequest{})
	require.NoError(t, err)
	require.Empty(t, resp.TokenPairs)
}

// TestTokenPairs_WithPairs verifies pagination over stored pairs.
func TestTokenPairs_WithPairs(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	addr1 := common.HexToAddress("0xAaBbCcDdEeFf0011223344556677889900AaBbCc")
	addr2 := common.HexToAddress("0x1122334455667788990011223344556677889900")

	pair1 := types.NewTokenPair(addr1, "usix", types.OWNER_MODULE)
	pair2 := types.NewTokenPair(addr2, "uatom", types.OWNER_EXTERNAL)
	k.SetToken(ctx, pair1)
	k.SetToken(ctx, pair2)

	resp, err := k.TokenPairs(ctx, &types.QueryTokenPairsRequest{})
	require.NoError(t, err)
	require.Len(t, resp.TokenPairs, 2)
}

// TestTokenPairs_NilRequest verifies that a nil request returns an error.
func TestTokenPairs_NilRequest(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	_, err := k.TokenPairs(ctx, nil)
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}

// TestTokenPair_ByDenom looks up a token pair by denom.
func TestTokenPair_ByDenom(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	addr := common.HexToAddress("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed")
	pair := types.NewTokenPair(addr, "usix", types.OWNER_MODULE)
	k.SetToken(ctx, pair)

	resp, err := k.TokenPair(ctx, &types.QueryTokenPairRequest{Token: "usix"})
	require.NoError(t, err)
	require.Equal(t, pair.Denom, resp.TokenPair.Denom)
}

// TestTokenPair_ByERC20Address looks up a token pair by hex address.
func TestTokenPair_ByERC20Address(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	addr := common.HexToAddress("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed")
	pair := types.NewTokenPair(addr, "usix", types.OWNER_MODULE)
	k.SetToken(ctx, pair)

	resp, err := k.TokenPair(ctx, &types.QueryTokenPairRequest{Token: addr.Hex()})
	require.NoError(t, err)
	require.Equal(t, pair.Denom, resp.TokenPair.Denom)
}

// TestTokenPair_NotFound returns NotFound error for unknown token.
func TestTokenPair_NotFound(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	_, err := k.TokenPair(ctx, &types.QueryTokenPairRequest{Token: "nosuchtoken"})
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, st.Code())
}

// TestTokenPair_InvalidFormat tests an invalid token format that is neither hex nor denom.
func TestTokenPair_InvalidFormat(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	// An invalid hex string that starts with 0x but isn't a valid address and isn't a valid denom
	_, err := k.TokenPair(ctx, &types.QueryTokenPairRequest{Token: "not a hex or valid format!!"})
	require.Error(t, err)
}

// TestTokenPair_NilRequest verifies nil returns InvalidArgument.
func TestTokenPair_NilRequest(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	_, err := k.TokenPair(ctx, nil)
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}

// TestQueryParams verifies the Params query returns stored parameters.
func TestQueryParams(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	// Set params
	params := types.NewParams(true, []string{}, []string{})
	require.NoError(t, k.SetParams(ctx, params))

	resp, err := k.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.True(t, resp.Params.EnableErc20)
}

// TestQueryParams_Default verifies default params when store is empty.
func TestQueryParams_Default(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	resp, err := k.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.False(t, resp.Params.EnableErc20)
}
