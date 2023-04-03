package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

var _ = strconv.Itoa(0)

func CmdSetConverterParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-converter-params [contract-address] [event-name] [event-tuple] [abi]",
		Short: "Set Converter config for converting asix to usix by using smart contract",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argContractAddress := args[0]
			argEventName := args[1]
			argEventTuple := args[2]
			argAbi := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetConverterParams(
				clientCtx.GetFromAddress().String(),
				argContractAddress,
				argEventName,
				argEventTuple,
				argAbi,
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
