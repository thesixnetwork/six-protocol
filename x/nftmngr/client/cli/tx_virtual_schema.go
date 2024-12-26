package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	nftmngrutils "github.com/thesixnetwork/six-protocol/x/nftmngr/client/utils"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// TODO:: TEST(VirtualSchema)
func CmdCreateVirtualSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-virtual-schema [schemaCode] [proposal-file]",
		Short: "Request for create new virtual_schema",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get value arguments
			argCode := args[0]
			argFilePath := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposal, err := nftmngrutils.ParseVirtualSchemaRegistryRequestJSON(clientCtx.LegacyAmino, argFilePath)
			if err != nil {
				return err
			}

			request := make([]types.VirtualSchemaRegistryRequest, 0)

			for _, req := range proposal {
				registry := types.NewVirtualSchemaRegistryRequest(req.Code, req.SharedAttributes)
				request = append(request, *registry)
			}

			msg := types.NewMsgCreateVirtualSchemaProposal(
				clientCtx.GetFromAddress().String(),
				argCode,
				request,
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

func CmdDeleteVirtualSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-virtual-schema [schemaCode]",
		Short: "Delete a virtual_schema",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteVirtualSchema(
				clientCtx.GetFromAddress().String(),
				indexIndex,
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
