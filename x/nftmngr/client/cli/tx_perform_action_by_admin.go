package cli

import (
	"strconv"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPerformActionByAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perform-action-by-nftadmin [nft-schema-code] [token-id] [action] [action-params]  [ref-id(optional)] ",
		Short: "To Perfrom action",
		Args:  cobra.RangeArgs(4, 5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNftSchemaCode := args[0]
			argTokenId := args[1]
			argAction := args[2]
			argActionParams := args[3]
			argRefId := args[4]

			if argActionParams == "" || argActionParams == " " {
				argActionParams = "[]"
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPerformActionByAdmin(
				clientCtx.GetFromAddress().String(),
				argNftSchemaCode,
				argTokenId,
				argAction,
				argActionParams,
				argRefId,
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
