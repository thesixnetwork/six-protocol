package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

var _ = strconv.Itoa(0)

func CmdSetUriRetrievalMethod() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-uri-retrieval-method [schema-code] [new-method: 0 or 1]",
		Short: "Update NewMethod for retrival base URI  ex 0:'BASE', 1:'TOKEN'",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSchemaCode := args[0]
			argNewMethod, err := cast.ToInt32E(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetUriRetrievalMethod(
				clientCtx.GetFromAddress().String(),
				argSchemaCode,
				argNewMethod,
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
