package evmbind_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/evmbind"
	"github.com/thesixnetwork/six-protocol/x/evmbind/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		BindingList: []types.Binding{
			{
				EthAddress: "0",
			},
			{
				EthAddress: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EvmbindKeeper(t)
	evmbind.InitGenesis(ctx, *k, genesisState)
	got := evmbind.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.BindingList, got.BindingList)
	// this line is used by starport scaffolding # genesis/test/assert
}
