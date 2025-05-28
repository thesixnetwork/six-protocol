package nftmngr

import (
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the nftschema
	for _, elem := range genState.NFTSchemaList {
		k.SetNftschema(ctx, elem)
	}
	// Set all the nftData
	for _, elem := range genState.NftDataList {
		k.SetNftData(ctx, elem)
	}
	// Set all the actionByRefId
	for _, elem := range genState.ActionByRefIdList {
		k.SetActionByRefId(ctx, elem)
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

	genesis.NFTSchemaList = k.GetAllNftschema(ctx)
	genesis.NftDataList = k.GetAllNftData(ctx)
	genesis.ActionByRefIdList = k.GetAllActionByRefId(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
