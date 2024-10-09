package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func CmdListActionExecutor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-action-executor",
		Short: "list all actionExecutor",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllActionExecutorRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ActionExecutorAll(context.Background(), params)
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

func CmdShowActionExecutor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-action-executor [nft-schema-code] [executor-address]",
		Short: "shows a actionExecutor",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argNftSchemaCode := args[0]
			argExecutorAddress := args[1]

			params := &types.QueryGetActionExecutorRequest{
				NftSchemaCode:   argNftSchemaCode,
				ExecutorAddress: argExecutorAddress,
			}

			res, err := queryClient.ActionExecutor(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
