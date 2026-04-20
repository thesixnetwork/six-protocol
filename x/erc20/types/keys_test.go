package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/x/erc20/types"
)

func TestModuleName(t *testing.T) {
	require.Equal(t, "erc20", types.ModuleName)
}

func TestStoreKey(t *testing.T) {
	require.Equal(t, types.ModuleName, types.StoreKey)
}

func TestRouterKey(t *testing.T) {
	require.Equal(t, types.ModuleName, types.RouterKey)
}

func TestKeyPrefixesAreUnique(t *testing.T) {
	prefixes := [][]byte{
		types.KeyPrefixTokenPair,
		types.KeyPrefixTokenPairByERC20,
		types.KeyPrefixTokenPairByDenom,
		types.KeyPrefixSTRv2Addresses,
	}

	seen := make(map[byte]bool)
	for _, p := range prefixes {
		require.Len(t, p, 1, "each prefix should be exactly one byte")
		require.False(t, seen[p[0]], "duplicate prefix byte found: %d", p[0])
		seen[p[0]] = true
	}
}

func TestModuleAddress(t *testing.T) {
	require.NotEmpty(t, types.ModuleAddress.Hex())
	require.NotEqual(t, "0x0000000000000000000000000000000000000000", types.ModuleAddress.Hex())
}

func TestOwnerConstants(t *testing.T) {
	require.NotEqual(t, types.OWNER_MODULE, types.OWNER_EXTERNAL)
	require.NotEqual(t, types.OWNER_MODULE, types.OWNER_UNSPECIFIED)
	require.NotEqual(t, types.OWNER_EXTERNAL, types.OWNER_UNSPECIFIED)
}
