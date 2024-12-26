package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func CmdListVirtualAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-virtual-action",
		Short: "list all virtual action",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVirtualActionRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VirtualActionAll(context.Background(), params)
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

func CmdShowVirtualAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-virtual-acion [nftSchemaCode] [name]",
		Short: "shows a virtual",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			nftSchemaCodeArg := args[0]
			nameArg := args[1]

			params := &types.QueryGetVirtualActionRequest{
				NftSchemaCode: nftSchemaCodeArg,
				Name: nameArg,
			}

			res, err := queryClient.VirtualAction(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
