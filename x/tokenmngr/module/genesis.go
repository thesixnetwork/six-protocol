package tokenmngr

import (
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the token
	for _, elem := range genState.TokenList {
		k.SetToken(ctx, elem)
	}
	// Set all the mintperm
	for _, elem := range genState.MintpermList {
		k.SetMintperm(ctx, elem)
	}
	// Set all the tokenBurn
	for _, elem := range genState.TokenBurnList {
		k.SetTokenBurn(ctx, elem)
	}
	// Set if defined
	if genState.Options != nil {
		k.SetOptions(ctx, *genState.Options)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TokenList = k.GetAllToken(ctx)
	genesis.MintpermList = k.GetAllMintperm(ctx)
	genesis.TokenBurnList = k.GetAllTokenBurn(ctx)
	// Get all options
	options, found := k.GetOptions(ctx)
	if found {
		genesis.Options = &options
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
