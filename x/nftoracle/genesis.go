package nftoracle

import (
	"github.com/thesixnetwork/six-protocol/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if genState.OracleConfig != nil {
		k.SetOracleConfig(ctx, *genState.OracleConfig)
	}

	// Set all the mintRequest
	for _, elem := range genState.MintRequestList {
		k.SetMintRequest(ctx, elem)
	}

	// Set mintRequest count
	k.SetMintRequestCount(ctx, genState.MintRequestCount)
	// Set all the actionRequest
	for _, elem := range genState.ActionRequestList {
		k.SetActionRequest(ctx, elem)
	}

	// Set actionRequest count
	k.SetActionRequestCount(ctx, genState.ActionRequestCount)
	// Set all the collectionOwnerRequest
	for _, elem := range genState.CollectionOwnerRequestList {
		k.SetCollectionOwnerRequest(ctx, elem)
	}

	// Set collectionOwnerRequest count
	k.SetCollectionOwnerRequestCount(ctx, genState.CollectionOwnerRequestCount)
	// Set if defined
	if genState.OracleConfig != nil {
		k.SetOracleConfig(ctx, *genState.OracleConfig)
	}
	// Set all the actionSigner
	for _, elem := range genState.ActionSignerList {
		k.SetActionSigner(ctx, elem)
	}
	// Set all the bindedSigner
	for _, elem := range genState.BindedSignerList {
		k.SetBindedSigner(ctx, elem)
	}
	// Set all the actionSignerConfig
	for _, elem := range genState.ActionSignerConfigList {
		k.SetActionSignerConfig(ctx, elem)
	}
	// Set all the syncActionSigner
	for _, elem := range genState.SyncActionSignerList {
		k.SetSyncActionSigner(ctx, elem)
	}

	// Set syncActionSigner count
	k.SetSyncActionSignerCount(ctx, genState.SyncActionSignerCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MintRequestList = k.GetAllMintRequest(ctx)
	genesis.MintRequestCount = k.GetMintRequestCount(ctx)
	genesis.ActionRequestList = k.GetAllActionRequest(ctx)
	genesis.ActionRequestCount = k.GetActionRequestCount(ctx)
	genesis.CollectionOwnerRequestList = k.GetAllCollectionOwnerRequest(ctx)
	genesis.CollectionOwnerRequestCount = k.GetCollectionOwnerRequestCount(ctx)
	// Get all oracleConfig
	oracleConfig, found := k.GetOracleConfig(ctx)
	if found {
		genesis.OracleConfig = &oracleConfig
	}
	genesis.ActionSignerList = k.GetAllActionSigner(ctx)
	genesis.BindedSignerList = k.GetAllBindedSigner(ctx)
	genesis.ActionSignerConfigList = k.GetAllActionSignerConfig(ctx)
	genesis.SyncActionSignerList = k.GetAllSyncActionSigner(ctx)
	genesis.SyncActionSignerCount = k.GetSyncActionSignerCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
