package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func CmdListTokenBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-burn",
		Short: "list all burn group by token",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTokenBurnRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TokenBurnAll(context.Background(), params)
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

func CmdShowTokenBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-burn [token]",
		Short: "shows a amount of burn per token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argToken := args[0]

			params := &types.QueryGetTokenBurnRequest{
				Token: argToken,
			}

			res, err := queryClient.TokenBurn(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
