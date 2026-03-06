package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"
)

func CmdListActionSignerConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-action-signer-config",
		Short: "list all actionSignerConfig",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllActionSignerConfigRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ActionSignerConfigAll(context.Background(), params)
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

func CmdShowActionSignerConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-action-signer-config [chain]",
		Short: "shows a actionSignerConfig",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChain := args[0]

			params := &types.QueryGetActionSignerConfigRequest{
				Chain: argChain,
			}

			res, err := queryClient.ActionSignerConfig(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
