// Copyright (C) 2024 SIX Network
// This file is part of the modified FeeMarket module from Ethermint (https://github.com/evmos/ethermint),
// and is licensed under the terms of the GNU Lesser General Public License v3
package simulation

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/thesixnetwork/six-protocol/x/feemarket/types"
)

// RandomizedGenState generates a random GenesisState for nft
func RandomizedGenState(simState *module.SimulationState) {
	params := types.NewParams(simState.Rand.Uint32()%2 == 0, simState.Rand.Uint32(), simState.Rand.Uint32(), simState.Rand.Uint64(), simState.Rand.Int63(), sdk.ZeroDec(), types.DefaultMinGasMultiplier, simState.Rand.Uint64(), sdk.ZeroDec())

	blockGas := simState.Rand.Uint64()
	feemarketGenesis := types.NewGenesisState(params, blockGas)

	bz, err := json.MarshalIndent(feemarketGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, bz)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(feemarketGenesis)
}
