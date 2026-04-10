package erc20
package erc20

import (




































}	}		TokenPairs: k.GetTokenPairs(ctx),		Params:     k.GetParams(ctx),	return &types.GenesisState{func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {// ExportGenesis exports module status}	}		k.SetToken(ctx, pair)	for _, pair := range data.TokenPairs {	}		panic("the erc20 module account has not been set")		// NOTE: shouldn't occur	if acc := accountKeeper.GetModuleAccount(ctx, types.ModuleName); acc == nil {	// ensure erc20 module account is set on genesis	}		panic("error setting erc20 params: " + err.Error())	if err := k.SetParams(ctx, data.Params); err != nil {) {	data types.GenesisState,	accountKeeper authkeeper.AccountKeeper,	k keeper.Keeper,	ctx sdk.Context,func InitGenesis(// InitGenesis imports module genesis)	"github.com/thesixnetwork/six-protocol/v4/x/erc20/types"	"github.com/thesixnetwork/six-protocol/v4/x/erc20/keeper"	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"	sdk "github.com/cosmos/cosmos-sdk/types"