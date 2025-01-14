package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func CmdListSchemaAttribute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-schema-attribute",
		Short: "list all schema_attribute",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSchemaAttributeRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SchemaAttributeAll(context.Background(), params)
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

func CmdShowSchemaAttribute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-schema-attribute [nft-schema-code] [name]",
		Short: "shows a schema_attribute",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argNftSchemaCode := args[0]
			argName := args[1]

			params := &types.QueryGetSchemaAttributeRequest{
				NftSchemaCode: argNftSchemaCode,
				Name:          argName,
			}

			res, err := queryClient.SchemaAttribute(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
