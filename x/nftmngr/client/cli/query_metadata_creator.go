package cli

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

func CmdListMetadataCreator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-metadata-creator",
		Short: "list all metadataCreator",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllMetadataCreatorRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.MetadataCreatorAll(context.Background(), params)
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

func CmdShowMetadataCreator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-metadata-creator [nft-schema-code]",
		Short: "shows a metadataCreator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argNftSchemaCode := args[0]

			params := &types.QueryGetMetadataCreatorRequest{
				NftSchemaCode: argNftSchemaCode,
			}

			res, err := queryClient.MetadataCreator(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
