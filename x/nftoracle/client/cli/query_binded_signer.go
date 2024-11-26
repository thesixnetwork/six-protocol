package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

func CmdShowBindedSigner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-binded-signer [owner-address]",
		Short: "shows a bindedSigner",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argOwnerAddress := args[0]

			params := &types.QueryGetBindedSignerRequest{
				OwnerAddress: argOwnerAddress,
			}

			res, err := queryClient.BindedSigner(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
