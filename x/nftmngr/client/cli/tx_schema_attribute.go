package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func CmdUpdateSchemaAttribute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-schema-attribute [nft-schema-code] [base64UpdateAttriuteDefenition]",
		Short: "Update a schema_attribute",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexNftSchemaCode := args[0]
			// Get value arguments
			argBase64NewAttriuteDefenition := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateSchemaAttribute(
				clientCtx.GetFromAddress().String(),
				indexNftSchemaCode,
				argBase64NewAttriuteDefenition,
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
