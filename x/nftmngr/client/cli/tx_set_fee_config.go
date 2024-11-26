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

func CmdSetFeeConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-fee-config [new-fee-config-base-64]",
		Short: "To set fee config",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNewFeeConfigBase64 := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetFeeConfig(
				clientCtx.GetFromAddress().String(),
				argNewFeeConfigBase64,
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
