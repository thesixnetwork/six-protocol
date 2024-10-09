package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func CmdListNFTSchemaByContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-nft-schema-by-contract",
		Short: "list all NFTSchemaByContract",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNFTSchemaByContractRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.NFTSchemaByContractAll(context.Background(), params)
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

func CmdShowNFTSchemaByContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-nft-schema-by-contract [origin-contract-address]",
		Short: "shows a NFTSchemaByContract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argOriginContractAddress := args[0]

			params := &types.QueryGetNFTSchemaByContractRequest{
				OriginContractAddress: argOriginContractAddress,
			}

			res, err := queryClient.NFTSchemaByContract(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
