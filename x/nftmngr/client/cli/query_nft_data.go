package cli

import (
	"context"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

func CmdListNftData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-nft-data",
		Short: "list all NftData",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			withGlobal := false
			if len(args) > 0 {
				argGlobal := args[0]
				_withGlobal, _ := strconv.ParseBool(argGlobal)
				withGlobal = _withGlobal
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNftDataRequest{
				WithGlobal: withGlobal,
				Pagination: pageReq,
			}

			res, err := queryClient.NftDataAll(context.Background(), params)
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

func CmdShowNftData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-nft-data [nft-schema-code] [token-id]",
		Short: "shows a NftData",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argNftSchemaCode := args[0]
			argTokenId := args[1]

			withGlobal := false
			if len(args) > 2 {
				argGlobal := args[2]
				_withGlobal, _ := strconv.ParseBool(argGlobal)
				withGlobal = _withGlobal
			}

			params := &types.QueryGetNftDataRequest{
				NftSchemaCode: argNftSchemaCode,
				TokenId:       argTokenId,
				WithGlobal:    withGlobal,
			}

			res, err := queryClient.NftData(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
