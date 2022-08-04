package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
	"github.com/thesixnetwork/six-protocol/app"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)
	rootCmd.AddCommand(
		AddGenesisWasmMsgCmd(app.DefaultNodeHome),
		Commands(app.DefaultNodeHome),
		GenGateTxCmd(app.ModuleBasics, cosmoscmd.MakeEncodingConfig(app.ModuleBasics).TxConfig, banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		CollectGenGateTxsCmd(banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
