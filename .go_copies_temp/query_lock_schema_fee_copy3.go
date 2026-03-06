package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

func CmdListLockSchemaFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-lock-schema-fee",
		Short: "list all lockSchemaFee",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLockSchemaFeeRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LockSchemaFeeAll(context.Background(), params)
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

func CmdShowLockSchemaFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-lock-schema-fee [index]",
		Short: "shows a lockSchemaFee",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetLockSchemaFeeRequest{
				Index: argIndex,
			}

			res, err := queryClient.LockSchemaFee(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
