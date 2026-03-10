package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"
)

var _ = strconv.Itoa(0)

func CmdCreateVerifyCollectionOwnerRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-verify-collection-owner-request [nft-schema-code] [base64VerifyRequestorSignature] [require-confirm]",
		Short: "To create Verify Collection Owner Request",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNftSchemaCode := args[0]
			argBase64VerifyRequestorSignature := args[1]
			argRequireConfirm, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVerifyCollectionOwnerRequest(
				clientCtx.GetFromAddress().String(),
				argNftSchemaCode,
				argBase64VerifyRequestorSignature,
				argRequireConfirm,
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
