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

func CmdSubmitActionResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-action-response [action-request-id] [base-64-nft-data]",
		Short: "To submit action response",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argActionRequestID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argBase64NftData := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitActionResponse(
				clientCtx.GetFromAddress().String(),
				argActionRequestID,
				argBase64NftData,
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
