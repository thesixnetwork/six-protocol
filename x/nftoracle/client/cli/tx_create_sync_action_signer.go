package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

var _ = strconv.Itoa(0)

func CmdCreateSyncActionSigner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-sync-action-signer [chain] [actor-address] [owner-address] [required-confirm]",
		Short: "Requesting to sync actionSigner with smartcontract",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChain := args[0]
			argActorAddress := args[1]
			argOwnerAddress := args[2]
			argRequiredConfirm, err := strconv.ParseUint(args[3], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSyncActionSigner(
				clientCtx.GetFromAddress().String(),
				argChain,
				argActorAddress,
				argOwnerAddress,
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
