package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

var _ = strconv.Itoa(0)

func CmdShowAttributes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-attributes [nft-schema-code] [true/false] [attribute-names(,separate)] ",
		Short: "To show attribute from market place aka delete from hidden list",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNftSchemaCode := args[0]
			argShow := args[1]
			argAttributeName := args[2]

			// Split attribute names
			attributeNames := strings.Split(argAttributeName, ",")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			show, _ := strconv.ParseBool(argShow)

			msg := types.NewMsgShowAttributes(
				clientCtx.GetFromAddress().String(),
				argNftSchemaCode,
				show,
				attributeNames,
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
