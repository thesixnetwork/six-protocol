package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateToken())
	cmd.AddCommand(CmdDeleteToken())
	cmd.AddCommand(CmdCreateMintperm())
	cmd.AddCommand(CmdUpdateMintperm())
	cmd.AddCommand(CmdDeleteMintperm())
	cmd.AddCommand(CmdMint())
	cmd.AddCommand(CmdCreateOptions())
	cmd.AddCommand(CmdUpdateOptions())
	cmd.AddCommand(CmdDeleteOptions())
	cmd.AddCommand(CmdBurn())
	cmd.AddCommand(CmdConvertToAtto())
	cmd.AddCommand(CmdConvertToMicro())
	cmd.AddCommand(CmdSetConverterParams())
	// this line is used by starport scaffolding # 1

	return cmd
}
