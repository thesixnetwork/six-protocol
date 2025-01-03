package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func CmdListInactiveEnableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-inactive-enable-virtual-schema-proposal",
		Short: "list all inactiveEnableVirtualSchemaProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllInactiveEnableVirtualSchemaProposalRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.InactiveEnableVirtualSchemaProposalAll(context.Background(), params)
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

func CmdShowInactiveEnableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-inactive-enable-virtual-schema-proposal [index]",
		Short: "shows a inactiveEnableVirtualSchemaProposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetInactiveEnableVirtualSchemaProposalRequest{
				Index: argIndex,
			}

			res, err := queryClient.InactiveEnableVirtualSchemaProposal(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
