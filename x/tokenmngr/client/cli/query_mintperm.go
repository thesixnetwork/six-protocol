package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func CmdListMintperm() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-mintperm",
		Short: "list all mintperm",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllMintpermRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.MintpermAll(context.Background(), params)
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

func CmdShowMintperm() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-mintperm [token] [address]",
		Short: "shows a mintperm",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argToken := args[0]
			argAddress := args[1]

			params := &types.QueryGetMintpermRequest{
				Token:   argToken,
				Address: argAddress,
			}

			res, err := queryClient.Mintperm(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
