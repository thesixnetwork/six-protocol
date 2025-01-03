package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func CmdListActiveDisableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-active-dislabe-virtual-schema-proposal",
		Short: "list all activeDisableVirtualSchemaProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllActiveDisableVirtualSchemaProposalRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ActiveDisableVirtualSchemaProposalAll(context.Background(), params)
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

func CmdShowActiveDisableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-active-dislabe-virtual-schema-proposal [index]",
		Short: "shows a activeDisableVirtualSchemaProposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetActiveDisableVirtualSchemaProposalRequest{
				Index: argIndex,
			}

			res, err := queryClient.ActiveDisableVirtualSchemaProposal(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
