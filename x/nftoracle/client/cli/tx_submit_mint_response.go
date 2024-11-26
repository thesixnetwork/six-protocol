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

func CmdSubmitMintResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-mint-response [mint-request-id] [base-64-nft-data]",
		Short: "To submit mint response",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMintRequestID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argBase64NftData := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitMintResponse(
				clientCtx.GetFromAddress().String(),
				argMintRequestID,
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
