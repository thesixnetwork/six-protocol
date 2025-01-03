package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

var _ = strconv.Itoa(0)

func CmdEnableVirtualSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable-virtual-schema [virtual-nft-schema-code]",
		Short: "Broadcast message enableVirtualSchema",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVirtualNftSchemaCode := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEnableVirtualSchema(
				clientCtx.GetFromAddress().String(),
				argVirtualNftSchemaCode,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
