package tokenmngr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the token
	for _, elem := range genState.TokenList {
		k.SetToken(ctx, elem)
	}
	// Set all the mintperm
	for _, elem := range genState.MintpermList {
		k.SetMintperm(ctx, elem)
	}
	// Set if defined
	if genState.Options != nil {
		k.SetOptions(ctx, *genState.Options)
	}
	// Set all the tokenBurn
	for _, elem := range genState.TokenBurnList {
		k.SetTokenBurn(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.TokenList = k.GetAllToken(ctx)
	genesis.MintpermList = k.GetAllMintperm(ctx)
	// Get all options
	options, found := k.GetOptions(ctx)
	if found {
		genesis.Options = &options
	}
	genesis.TokenBurnList = k.GetAllTokenBurn(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
