package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

var _ = strconv.Itoa(0)

func CmdListAdminOfGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-admin-of-group [group]",
		Short: "Query list-admin-of-group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqGroup := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListAdminOfGroupRequest{
				Group: reqGroup,
			}

			res, err := queryClient.ListAdminOfGroup(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
