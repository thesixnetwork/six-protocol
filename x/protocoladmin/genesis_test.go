package protocoladmin_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		GroupList: []types.Group{
			{
				Name: "0",
			},
			{
				Name: "1",
			},
		},
		AdminList: []types.Admin{
			{
				Group: "0",
				Admin: "0",
			},
			{
				Group: "1",
				Admin: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ProtocoladminKeeper(t)
	protocoladmin.InitGenesis(ctx, *k, genesisState)
	got := protocoladmin.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.GroupList, got.GroupList)
	require.ElementsMatch(t, genesisState.AdminList, got.AdminList)
	// this line is used by starport scaffolding # genesis/test/assert
}
