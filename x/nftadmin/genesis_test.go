package nftadmin_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	nftadmin "github.com/thesixnetwork/six-protocol/x/nftadmin"
	"github.com/thesixnetwork/six-protocol/x/nftadmin/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Authorization: &types.Authorization{
			RootAdmin:   "12",
			Permissions: new(types.Permissions),
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NftAdminKeeper(t)
	nftadmin.InitGenesis(ctx, *k, genesisState)
	got := nftadmin.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Authorization, got.Authorization)
	// this line is used by starport scaffolding # genesis/test/assert
}
