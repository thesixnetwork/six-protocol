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

func CmdUpdateAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-action [nft-schema-code] [name] [base-64-update-action]",
		Short: "Update action of selected schema",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNftSchemaCode := args[0]
			argBase64UpdateAction := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAction(
				clientCtx.GetFromAddress().String(),
				argNftSchemaCode,
				argBase64UpdateAction,
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
