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

func CmdSubmitSyncActionSigner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-sync-action-signer [sync-id] [chain] [actor-address] [owner-address] [expire-epoch]",
		Short: "Submit Verified actionSigner from smartcontract",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSyncId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argChain := args[1]
			argActorAddress := args[2]
			argOwnerAddress := args[3]
			argExpireEpochString := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitSyncActionSigner(
				clientCtx.GetFromAddress().String(),
				argChain,
				argSyncId,
				argActorAddress,
				argOwnerAddress,
				argExpireEpochString,
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
