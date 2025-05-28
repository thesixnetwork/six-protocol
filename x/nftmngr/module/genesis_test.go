package nftmngr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	nftmngr "github.com/thesixnetwork/six-protocol/x/nftmngr/module"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		NFTSchemaList: []types.NFTSchema{
			{
				Code: "0",
			},
			{
				Code: "1",
			},
		},
		NftDataList: []types.NftData{
			{
				NftSchemaCode: "0",
				TokenId:       "0",
			},
			{
				NftSchemaCode: "1",
				TokenId:       "1",
			},
		},
		ActionByRefIdList: []types.ActionByRefId{
			{
				RefId: "0",
			},
			{
				RefId: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NftmngrKeeper(t)
	nftmngr.InitGenesis(ctx, k, genesisState)
	got := nftmngr.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.NFTSchemaList, got.NFTSchemaList)
	require.ElementsMatch(t, genesisState.NftDataList, got.NftDataList)
	require.ElementsMatch(t, genesisState.ActionByRefIdList, got.ActionByRefIdList)
	// this line is used by starport scaffolding # genesis/test/assert
}
