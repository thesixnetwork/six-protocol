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

func CmdCreateMultiMetadata() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-multi-metadata [nft-schema-code] [token-id,token-id,token-id,...] [base-64-nft-data]",
		Short: "To create Multiple NFT Metadata(s) from Base64 required flag in base64",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNftSchemaCode := args[0]
			argTokenIds := args[1]
			argBase64NFTData := args[2]

			// Split tokenIds
			tokenId := strings.Split(argTokenIds, ",")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateMultiMetadata(
				clientCtx.GetFromAddress().String(),
				argNftSchemaCode,
				tokenId,
				argBase64NFTData,
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
