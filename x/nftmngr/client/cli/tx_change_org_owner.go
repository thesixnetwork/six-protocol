package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdChangeOrgOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "change-org-owner [org-name] [to-new-owner]",
		Short: "Change Organization Owner",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOrgName := args[0]
			argToNewOwner := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgChangeOrgOwner(
				clientCtx.GetFromAddress().String(),
				argOrgName,
				argToNewOwner,
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
