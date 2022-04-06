package tokenmngr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		TokenList: []types.Token{
			{
				Name: "0",
			},
			{
				Name: "1",
			},
		},
		MintpermList: []types.Mintperm{
			{
				Token:   "0",
				Address: "0",
			},
			{
				Token:   "1",
				Address: "1",
			},
		},
		Options: &types.Options{
			DefaultMintee: "11",
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokenmngrKeeper(t)
	tokenmngr.InitGenesis(ctx, *k, genesisState)
	got := tokenmngr.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.TokenList, got.TokenList)
	require.ElementsMatch(t, genesisState.MintpermList, got.MintpermList)
	require.Equal(t, genesisState.Options, got.Options)
	// this line is used by starport scaffolding # genesis/test/assert
}
