package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

func CmdListActionOfSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-action-of-schema",
		Short: "list all action_of_schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllActionOfSchemaRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ActionOfSchemaAll(context.Background(), params)
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

func CmdShowActionOfSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-action-of-schema [nft-schema-code] [name]",
		Short: "shows a action_of_schema",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argNftSchemaCode := args[0]
			argName := args[1]

			params := &types.QueryGetActionOfSchemaRequest{
				NftSchemaCode: argNftSchemaCode,
				Name:          argName,
			}

			res, err := queryClient.ActionOfSchema(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
