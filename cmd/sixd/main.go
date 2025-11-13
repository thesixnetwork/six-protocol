package main

import (
	"fmt"
	"os"

	"github.com/thesixnetwork/six-protocol/v4/app"
	"github.com/thesixnetwork/six-protocol/v4/cmd/sixd/cmd"

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
