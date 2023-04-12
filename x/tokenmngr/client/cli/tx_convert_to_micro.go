package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

var _ = strconv.Itoa(0)

func CmdConvertToMicro() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-to-micro [amount] [RECEIVER(optional)]",
		Short: "Convert native token from 10^18 to native token 10^6",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			coins, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			var receiver string
			if len(args) == 2 {
				receiver = args[1]
				_, err := sdk.AccAddressFromBech32(receiver)
				if err != nil {
					return fmt.Errorf("invalid receiver address: %s", err)
				}
			} else {
				receiver = creator
			}

			msg := types.NewMsgConvertToMicro(
				clientCtx.GetFromAddress().String(),
				coins,
				receiver,
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
