package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/consume/types"
)

var _ = strconv.Itoa(0)

func CmdUseNftByEVM() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "use-nft-by-evm [token] [eth-address] [eth-signature] [sign-message]",
		Short: "Broadcast message useNftByEVM",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argToken := args[0]
			argEthAddress := args[1]
			argEthSignature := args[2]
			argSignMessage := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUseNftByEVM(
				clientCtx.GetFromAddress().String(),
				argToken,
				argEthAddress,
				argEthSignature,
				argSignMessage,
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
