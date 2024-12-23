package nftoracle_test

import (
	"testing"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"

	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MintRequestList: []types.MintRequest{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MintRequestCount: 2,
		ActionRequestList: []types.ActionOracleRequest{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ActionRequestCount: 2,
		CollectionOwnerRequestList: []types.CollectionOwnerRequest{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CollectionOwnerRequestCount: 2,
		OracleConfig:                &types.OracleConfig{},
		ActionSignerList: []types.ActionSigner{
			{
				ActorAddress: "0",
			},
			{
				ActorAddress: "1",
			},
		},
		BindedSignerList: []types.BindedSigner{
			{
				OwnerAddress: "0",
			},
			{
				OwnerAddress: "1",
			},
		},
		ActionSignerConfigList: []types.ActionSignerConfig{
			{
				Chain: "0",
			},
			{
				Chain: "1",
			},
		},
		SyncActionSignerList: []types.SyncActionSigner{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SyncActionSignerCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NftoracleKeeper(t)
	nftoracle.InitGenesis(ctx, *k, genesisState)
	got := nftoracle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MintRequestList, got.MintRequestList)
	require.Equal(t, genesisState.MintRequestCount, got.MintRequestCount)
	require.ElementsMatch(t, genesisState.ActionRequestList, got.ActionRequestList)
	require.Equal(t, genesisState.ActionRequestCount, got.ActionRequestCount)
	require.ElementsMatch(t, genesisState.CollectionOwnerRequestList, got.CollectionOwnerRequestList)
	require.Equal(t, genesisState.CollectionOwnerRequestCount, got.CollectionOwnerRequestCount)
	require.Equal(t, genesisState.OracleConfig, got.OracleConfig)
	require.ElementsMatch(t, genesisState.ActionSignerList, got.ActionSignerList)
	require.ElementsMatch(t, genesisState.BindedSignerList, got.BindedSignerList)
	require.ElementsMatch(t, genesisState.ActionSignerConfigList, got.ActionSignerConfigList)
	require.ElementsMatch(t, genesisState.SyncActionSignerList, got.SyncActionSignerList)
	require.Equal(t, genesisState.SyncActionSignerCount, got.SyncActionSignerCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
