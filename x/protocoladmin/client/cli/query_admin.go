package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

func CmdListAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-admin",
		Short: "list all admin",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAdminRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AdminAll(context.Background(), params)
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

func CmdShowAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-admin [group] [admin]",
		Short: "shows a admin",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argGroup := args[0]
			argAdmin := args[1]

			params := &types.QueryGetAdminRequest{
				Group: argGroup,
				Admin: argAdmin,
			}

			res, err := queryClient.Admin(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
