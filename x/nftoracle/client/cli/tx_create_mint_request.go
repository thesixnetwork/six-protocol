package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

var _ = strconv.Itoa(0)

func CmdCreateMintRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-mint-request [nft-schema-code] [token-id] [required-confirm]",
		Short: "To create Mint Request",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNftSchemaCode := args[0]
			argTokenId := args[1]
			argRequiredConfirm, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateMintRequest(
				clientCtx.GetFromAddress().String(),
				argNftSchemaCode,
				argTokenId,
				argRequiredConfirm,
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
