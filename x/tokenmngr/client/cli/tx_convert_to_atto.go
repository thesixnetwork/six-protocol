package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

var _ = strconv.Itoa(0)

func CmdConvertToAtto() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-to-atto [amount]",
		Short: "Convert native token from 10^6 to native token 10^18 for usign with Evm",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			coins, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgConvertToAtto(
				clientCtx.GetFromAddress().String(),
				coins,
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
