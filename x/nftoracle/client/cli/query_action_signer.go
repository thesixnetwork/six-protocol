package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

func CmdListActionSigner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-action-signer",
		Short: "list all actionSigner",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllActionSignerRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ActionSignerAll(context.Background(), params)
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

func CmdShowActionSigner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-action-signer [actor-address] [owner-address]",
		Short: "shows a actionSigner",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argActorAddress := args[0]
			argOwnerAddress := args[1]

			params := &types.QueryGetActionSignerRequest{
				ActorAddress: argActorAddress,
				OwnerAddress: argOwnerAddress,
			}

			res, err := queryClient.ActionSigner(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
