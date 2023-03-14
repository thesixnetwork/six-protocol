package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/thesixnetwork/six-protocol/app"
	"github.com/thesixnetwork/six-protocol/cmd/sixd/cmd"
	cmdcfg "github.com/thesixnetwork/six-protocol/cmd/sixd/config"
)

func main() {
	app.SetConfig()
	cmdcfg.RegisterDenoms()
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
