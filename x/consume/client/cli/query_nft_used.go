package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/consume/types"
)

func CmdListNftUsed() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-nft-used",
		Short: "list all NftUsed",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNftUsedRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.NftUsedAll(context.Background(), params)
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

func CmdShowNftUsed() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-nft-used [token]",
		Short: "shows a NftUsed",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argToken := args[0]

			params := &types.QueryGetNftUsedRequest{
				Token: argToken,
			}

			res, err := queryClient.NftUsed(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
