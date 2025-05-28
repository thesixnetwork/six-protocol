package main

import (
	"fmt"
	"os"

	"github.com/thesixnetwork/six-protocol/app"
	"github.com/thesixnetwork/six-protocol/cmd/sixd/cmd"

	clienthelpers "cosmossdk.io/client/v2/helpers"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, clienthelpers.EnvPrefix, app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
