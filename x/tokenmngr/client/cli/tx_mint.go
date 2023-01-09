package cli

import (
	"strconv"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	// "github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

var _ = strconv.Itoa(0)

func CmdMint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [amount]",
		Short: "Broadcast message mint",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}


			coins, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgMint(
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
