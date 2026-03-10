package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdCreateNFTSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-nft-schema [nft-schema-base-64]",
		Short: "To create NFT schema from JSON in Base64 format",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNftSchemaBase64 := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateNFTSchema(
				clientCtx.GetFromAddress().String(),
				argNftSchemaBase64,
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
