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

func CmdSubmitVerifyCollectionOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-verify-collection-owner [verify-request-id] [schema-code] [base64OriginContractInfo]",
		Short: "To Sumint Verify Collection Owner",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVerifyRequestID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argSchemaCode := args[1]
			argContractOriginDataInfo := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitVerifyCollectionOwner(
				clientCtx.GetFromAddress().String(),
				argVerifyRequestID,
				argSchemaCode,
				argContractOriginDataInfo,
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
