// Copyright (C) 2024 SIX Network
// This file is part of the modified FeeMarket module from Ethermint, 
// and is licensed under the terms of the GNU Lesser General Public License v3
package v012

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	v010types "github.com/thesixnetwork/six-protocol/x/feemarket/migrations/v010/types"
	"github.com/thesixnetwork/six-protocol/x/feemarket/types"
)

// MigrateStore adds the MinGasPrice param with a value of 0
// and MinGasMultiplier to 0,5
func MigrateStore(ctx sdk.Context, paramstore *paramtypes.Subspace) error {
	if !paramstore.HasKeyTable() {
		ps := paramstore.WithKeyTable(types.ParamKeyTable())
		paramstore = &ps
	}

	// add LegacyBaseFee
	paramstore.Set(ctx, types.ParamStoreKeyLegacyBaseFee, types.DefaultLegacyBaseFee)
	// add LegacyMinGasPrice
	paramstore.Set(ctx, types.ParamStoreKeyLegacyMinGasPrice, types.DefaultLegacyMinGasPrice)
	return nil
}

// MigrateJSON accepts exported v0.12 x/feemarket genesis state and migrates it to
// v0.12 x/feemarket genesis state. The migration includes:
// FIX Broken when perform legacy tx
// - add LegacyTXMinGasPrice param
// - add LegacyTXMinGasMultiplier param
func MigrateJSON(oldState v010types.GenesisState) types.GenesisState {
	return types.GenesisState{
		Params: types.Params{
			NoBaseFee:                oldState.Params.NoBaseFee,
			BaseFeeChangeDenominator: oldState.Params.BaseFeeChangeDenominator,
			ElasticityMultiplier:     oldState.Params.ElasticityMultiplier,
			EnableHeight:             oldState.Params.EnableHeight,
			BaseFee:                  oldState.Params.BaseFee,
			MinGasPrice:              types.DefaultMinGasPrice,
			MinGasMultiplier:         types.DefaultMinGasMultiplier,
			LegacyBaseFee:            types.DefaultLegacyBaseFee,
			LegacyMinGasPrice:        types.DefaultLegacyMinGasPrice,
		},
		BlockGas: oldState.BlockGas,
	}
}
