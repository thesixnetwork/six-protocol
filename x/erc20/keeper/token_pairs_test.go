package keeper_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/erc20/types"
)

// validIBCDenom is a well-formed IBC denom whose hex hash is exactly 32 bytes.
const validIBCDenom = "ibc/DF63978F803A2E27CA5CC9B7631654CCF0BBC788B3B7F0A10200508E37C70992"

// testERC20Addr returns a deterministic test ERC20 address.
func testERC20Addr() common.Address {
	return common.HexToAddress("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed")
}

// TestSetTokenAndGetTokenPair verifies that SetToken stores a pair and it can be
// retrieved via GetTokenPair / GetTokenPairID.
func TestSetTokenAndGetTokenPair(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	pair := types.NewTokenPair(testERC20Addr(), "usix", types.OWNER_MODULE)
	k.SetToken(ctx, pair)

	id := k.GetTokenPairID(ctx, "usix")
	require.NotEmpty(t, id)

	got, found := k.GetTokenPair(ctx, id)
	require.True(t, found)
	require.Equal(t, pair, got)
}

// TestGetTokenPairs returns all stored token pairs.
func TestGetTokenPairs(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	addr1 := common.HexToAddress("0xAaBbCcDdEeFf0011223344556677889900AaBbCc")
	addr2 := common.HexToAddress("0x1122334455667788990011223344556677889900")

	pair1 := types.NewTokenPair(addr1, "usix", types.OWNER_MODULE)
	pair2 := types.NewTokenPair(addr2, "uatom", types.OWNER_EXTERNAL)
	k.SetToken(ctx, pair1)
	k.SetToken(ctx, pair2)

	all := k.GetTokenPairs(ctx)
	require.Len(t, all, 2)
}

// TestIsDenomRegistered verifies registration flag after SetToken / DeleteTokenPair.
func TestIsDenomRegistered(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	pair := types.NewTokenPair(testERC20Addr(), "usix", types.OWNER_MODULE)

	require.False(t, k.IsDenomRegistered(ctx, "usix"))

	k.SetToken(ctx, pair)
	require.True(t, k.IsDenomRegistered(ctx, "usix"))

	k.DeleteTokenPair(ctx, pair)
	require.False(t, k.IsDenomRegistered(ctx, "usix"))
}

// TestIsERC20Registered verifies ERC20 address registration flag.
func TestIsERC20Registered(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	addr := testERC20Addr()
	pair := types.NewTokenPair(addr, "usix", types.OWNER_MODULE)

	require.False(t, k.IsERC20Registered(ctx, addr))

	k.SetToken(ctx, pair)
	require.True(t, k.IsERC20Registered(ctx, addr))
}

// TestIsTokenPairRegistered checks the full ID registration.
func TestIsTokenPairRegistered(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	pair := types.NewTokenPair(testERC20Addr(), "usix", types.OWNER_MODULE)
	require.False(t, k.IsTokenPairRegistered(ctx, pair.GetID()))

	k.SetToken(ctx, pair)
	require.True(t, k.IsTokenPairRegistered(ctx, pair.GetID()))
}

// TestGetTokenPairIDByERC20Addr verifies lookup by ERC20 hex address string.
func TestGetTokenPairIDByERC20Addr(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	addr := testERC20Addr()
	pair := types.NewTokenPair(addr, "usix", types.OWNER_MODULE)
	k.SetToken(ctx, pair)

	id := k.GetTokenPairID(ctx, addr.Hex())
	require.NotEmpty(t, id)

	got, found := k.GetTokenPair(ctx, id)
	require.True(t, found)
	require.Equal(t, pair.Denom, got.Denom)
}

// TestGetCoinAddress returns the ERC20 address for a registered denom.
func TestGetCoinAddress(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	addr := testERC20Addr()
	pair := types.NewTokenPair(addr, "usix", types.OWNER_MODULE)
	k.SetToken(ctx, pair)

	got, err := k.GetCoinAddress(ctx, "usix")
	require.NoError(t, err)
	require.Equal(t, addr, got)
}

// TestGetTokenDenom returns the denom for a registered ERC20 address.
func TestGetTokenDenom(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	addr := testERC20Addr()
	pair := types.NewTokenPair(addr, "usix", types.OWNER_MODULE)
	k.SetToken(ctx, pair)

	got, err := k.GetTokenDenom(ctx, addr)
	require.NoError(t, err)
	require.Equal(t, "usix", got)
}

// TestCreateNewTokenPair_IBC verifies the STRv2 path with a proper IBC denom.
func TestCreateNewTokenPair_IBC(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	pair, err := k.CreateNewTokenPair(ctx, validIBCDenom)
	require.NoError(t, err)
	require.Equal(t, validIBCDenom, pair.Denom)
	require.True(t, pair.Enabled)
	require.Equal(t, types.OWNER_MODULE, pair.ContractOwner)

	require.True(t, k.IsDenomRegistered(ctx, validIBCDenom))
	require.True(t, k.IsERC20Registered(ctx, pair.GetERC20Contract()))
}

// TestCreateNewTokenPair_NonIBC verifies that a non-IBC denom returns an error.
func TestCreateNewTokenPair_NonIBC(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	_, err := k.CreateNewTokenPair(ctx, "usix")
	require.Error(t, err, "expected error for non-IBC denom")
}

// TestDeleteTokenPair verifies that a pair is fully removed from all maps.
func TestDeleteTokenPair(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	addr := testERC20Addr()
	pair := types.NewTokenPair(addr, "usix", types.OWNER_MODULE)
	k.SetToken(ctx, pair)

	id := k.GetTokenPairID(ctx, "usix")
	require.NotEmpty(t, id)

	k.DeleteTokenPair(ctx, pair)

	_, found := k.GetTokenPair(ctx, id)
	require.False(t, found)
	require.False(t, k.IsDenomRegistered(ctx, "usix"))
	require.False(t, k.IsERC20Registered(ctx, addr))
}
