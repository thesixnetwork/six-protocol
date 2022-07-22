package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
	"github.com/thesixnetwork/six-protocol/app"
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
		// AddGenesisAccountCmd(app.DefaultNodeHome), this module is already included in the scaffolding  # root/arguments
		// migratecmd.AddMigrateCmd(app.DefaultNodeHome), this module is already included in the scaffolding  # root/arguments
		// GenTxCmd(app.DefaultNodeHome), this module is already included in the scaffolding  # root/arguments
		AddGenesisWasmMsgCmd(app.DefaultNodeHome),
		AddEthKeyCommands(app.DefaultNodeHome),
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
