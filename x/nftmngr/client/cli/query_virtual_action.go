package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

func CmdListVirtualAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-virtual-action",
		Short: "list all virtual action [nftSchemaCode(optional)]",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			argSchemaCode := ""
			if len(args) > 0 {
				argSchemaCode = args[0]
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVirtualActionRequest{
				NftSchemaCode: argSchemaCode,
				Pagination:    pageReq,
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
				Name:          nameArg,
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
