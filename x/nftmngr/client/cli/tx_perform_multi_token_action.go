package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

var _ = strconv.Itoa(0)

func CmdPerformMultiTokenAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perform-multi-token-action [nft-schema-code] [token-id,token-id,token-id,...] [action] [ref-id] [params, params, params, ...]",
		Short: "To do action for multiple NFTs ex: sixnftd performMultiTokenAction schema 1,2,3,4 action ref001 '[ [] ]'",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNftSchemaCode := args[0]
			argTokenIds := args[1]
			argAction := args[2]
			argRefId := args[3]
			argParameters := args[4]
			arryOfparams := strings.Split(args[4], "],")
			for i, params := range arryOfparams {
				if i != len(arryOfparams)-1 {
					arryOfparams[i] = params + "]"
				} else {
					arryOfparams[i] = params
				}
			}

			// Split tokenIds
			tokenId := strings.Split(argTokenIds, ",")

			// Split actions
			action := strings.Split(argAction, ",")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPerformMultiTokenAction(
				clientCtx.GetFromAddress().String(),
				argNftSchemaCode,
				tokenId,
				action,
				argRefId,
				argParameters,
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
