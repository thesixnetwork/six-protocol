package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func CmdListActiveVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-active-virtual-schema-proposal",
		Short: "list all activeVirtualSchemaProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllActiveVirtualSchemaProposalRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ActiveVirtualSchemaProposalAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowActiveVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-active-virtual-schema-proposal [index]",
		Short: "shows a activeVirtualSchemaProposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetActiveVirtualSchemaProposalRequest{
				Index: argIndex,
			}

			res, err := queryClient.ActiveVirtualSchemaProposal(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
