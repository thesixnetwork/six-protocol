package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// TODO:: Feat(VirtualSchema)
func CmdListVirtualSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-virtual-schema",
		Short: "list all virtual_schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVirtualSchemaRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VirtualSchemaAll(context.Background(), params)
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

// TODO:: Feat(VirtualSchema)
func CmdShowVirtualSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-virtual-schema [schemaCode]",
		Short: "shows a virtual_schema",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetVirtualSchemaRequest{
				NftSchemaCode: argIndex,
			}

			res, err := queryClient.VirtualSchema(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
