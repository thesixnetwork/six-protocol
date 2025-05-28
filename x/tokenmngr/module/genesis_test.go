package tokenmngr_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	tokenmngr "github.com/thesixnetwork/six-protocol/x/tokenmngr/module"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

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
		TokenBurnList: []types.TokenBurn{
			{
				Amount: sdk.NewCoin("0", sdkmath.NewInt(0)),
			},
			{
				Amount: sdk.NewCoin("1", sdkmath.NewInt(1)),
			},
		},
		Options: &types.Options{
			DefaultMintee: "28",
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokenmngrKeeper(t)
	tokenmngr.InitGenesis(ctx, k, genesisState)
	got := tokenmngr.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TokenList, got.TokenList)
	require.ElementsMatch(t, genesisState.MintpermList, got.MintpermList)
	require.ElementsMatch(t, genesisState.TokenBurnList, got.TokenBurnList)
	require.Equal(t, genesisState.Options, got.Options)
	// this line is used by starport scaffolding # genesis/test/assert
}
