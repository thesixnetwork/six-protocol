package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func CmdListVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-virtual-schema-proposal",
		Short: "list all virtualSchemaProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVirtualSchemaProposalRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VirtualSchemaProposalAll(context.Background(), params)
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

func CmdShowVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-virtual-schema-proposal [id]",
		Short: "shows a virtualSchemaProposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetVirtualSchemaProposalRequest{
				Index: argIndex,
			}

			res, err := queryClient.VirtualSchemaProposal(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListActiveDisableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-active-dislabe-virtual-schema-proposal",
		Short: "list all activeDisableVirtualSchemaProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllActiveDisableVirtualSchemaProposalRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ActiveDisableVirtualSchemaProposalAll(context.Background(), params)
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

func CmdShowActiveDisableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-active-dislabe-virtual-schema-proposal [id]",
		Short: "shows a activeDisableVirtualSchemaProposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetActiveDisableVirtualSchemaProposalRequest{
				Index: argIndex,
			}

			res, err := queryClient.ActiveDisableVirtualSchemaProposal(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListActiveEnableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-active-enable-virtual-schema-proposal",
		Short: "list all activeEnableVirtualSchemaProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllActiveEnableVirtualSchemaProposalRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ActiveEnableVirtualSchemaProposalAll(context.Background(), params)
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

func CmdShowActiveEnableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-active-enable-virtual-schema-proposal [id]",
		Short: "shows a activeEnableVirtualSchemaProposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetActiveEnableVirtualSchemaProposalRequest{
				Index: argIndex,
			}

			res, err := queryClient.ActiveEnableVirtualSchemaProposal(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListDisableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-disable-virtual-schema",
		Short: "list all disableVirtualSchema",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDisableVirtualSchemaProposalRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DisableVirtualSchemaProposalAll(context.Background(), params)
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

func CmdShowDisableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-disable-virtual-schema [id]",
		Short: "shows a disableVirtualSchema",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetDisableVirtualSchemaProposalRequest{
				Index: argIndex,
			}

			res, err := queryClient.DisableVirtualSchemaProposal(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListInactiveDisableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-inactive-disable-virtual-schema-proposal",
		Short: "list all inactiveDisableVirtualSchemaProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllInactiveDisableVirtualSchemaProposalRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.InactiveDisableVirtualSchemaProposalAll(context.Background(), params)
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

func CmdShowInactiveDisableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-inactive-disable-virtual-schema-proposal [index]",
		Short: "shows a inactiveDisableVirtualSchemaProposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetInactiveDisableVirtualSchemaProposalRequest{
				Index: argIndex,
			}

			res, err := queryClient.InactiveDisableVirtualSchemaProposal(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListInactiveEnableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-inactive-enable-virtual-schema-proposal",
		Short: "list all inactiveEnableVirtualSchemaProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllInactiveEnableVirtualSchemaProposalRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.InactiveEnableVirtualSchemaProposalAll(context.Background(), params)
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

func CmdShowInactiveEnableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-inactive-enable-virtual-schema-proposal [id]",
		Short: "shows a inactiveEnableVirtualSchemaProposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetInactiveEnableVirtualSchemaProposalRequest{
				Index: argIndex,
			}

			res, err := queryClient.InactiveEnableVirtualSchemaProposal(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListEnableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-enable-virtual-schema-proposal",
		Short: "list all enableVirtualSchemaProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllEnableVirtualSchemaProposalRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.EnableVirtualSchemaProposalAll(context.Background(), params)
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

func CmdShowEnableVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-enable-virtual-schema-proposal [id]",
		Short: "shows a enableVirtualSchemaProposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetEnableVirtualSchemaProposalRequest{
				Index: argIndex,
			}

			res, err := queryClient.EnableVirtualSchemaProposal(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
