package protocoladmin

import (
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/keeper"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the admin
	for _, elem := range genState.AdminList {
		k.SetAdmin(ctx, elem)
	}
	// Set all the group
	for _, elem := range genState.GroupList {
		k.SetGroup(ctx, elem)
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

	genesis.AdminList = k.GetAllAdmin(ctx)
	genesis.GroupList = k.GetAllGroup(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
