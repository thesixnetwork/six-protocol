package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	nftmngrutils "github.com/thesixnetwork/six-protocol/x/nftmngr/client/utils"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdPerformVirtualAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perform-virtual-action [action-file] [ref-id]",
		Short: "Broadcast message performVirtualAction",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var refId string
			argActionFilePath := args[0]
			if len(args) > 1 {
				refId = args[1]
			} else {
				refId = ""
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			actionJSON, err := nftmngrutils.ParseActionJSON(clientCtx.LegacyAmino, argActionFilePath)
			if err != nil {
				return err
			}

			var (
				tokenIdMapInput []*types.TokenIdMap
				paramInput      []*types.ActionParameter
			)

			for _, actionParam := range actionJSON.Params {
				paramInput = append(paramInput, &types.ActionParameter{
					Name:  actionParam.Name,
					Value: actionParam.Value,
				})
			}

			for _, tokenId := range actionJSON.TokenIds {
				tokenIdMapInput = append(tokenIdMapInput, &types.TokenIdMap{
					NftSchemaName: tokenId.Schema,
					TokenId:       tokenId.TokenId,
				})
			}

			msg := types.NewMsgPerformVirtualAction(
				clientCtx.GetFromAddress().String(),
				actionJSON.Code,
				tokenIdMapInput,
				actionJSON.Action,
				paramInput,
				refId,
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
