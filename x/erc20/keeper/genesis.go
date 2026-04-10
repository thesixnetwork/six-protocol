package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/v4/x/erc20/types"
)

// InitGenesis initializes the erc20 module genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	if err := k.SetParams(ctx, data.Params); err != nil {
		panic("failed to set erc20 params: " + err.Error())
	}

	for _, pair := range data.TokenPairs {
		k.SetToken(ctx, pair)
	}
}

// ExportGenesis exports the erc20 module genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Params:     k.GetParams(ctx),
		TokenPairs: k.GetTokenPairs(ctx),
	}
}
