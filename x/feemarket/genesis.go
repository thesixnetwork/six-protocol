// Copyright (C) 2024 SIX Network
// This file is part of the modified FeeMarket module from Ethermint (https://github.com/evmos/ethermint),
// and is licensed under the terms of the GNU Lesser General Public License v3
package feemarket

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/thesixnetwork/six-protocol/x/feemarket/keeper"
	"github.com/thesixnetwork/six-protocol/x/feemarket/types"
)

// InitGenesis initializes genesis state based on exported genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	data types.GenesisState,
) []abci.ValidatorUpdate {
	k.SetParams(ctx, data.Params)
	k.SetBlockGasWanted(ctx, data.BlockGas)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports genesis state of the fee market module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:   k.GetParams(ctx),
		BlockGas: k.GetBlockGasWanted(ctx),
	}
}
