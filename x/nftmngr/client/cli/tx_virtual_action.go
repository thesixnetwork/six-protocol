package cli

import (
	"github.com/spf13/cobra"
	nftmngrutils "github.com/thesixnetwork/six-protocol/x/nftmngr/client/utils"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

func CmdCreateVirtualAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-virtual-action [code] [new-action-file-path]",
		Short: "add a new virtual action to virtual schema",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCode := args[0]
			argFilePath := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			newAction, err := nftmngrutils.ParseNewVirtualActionJSON(clientCtx.LegacyAmino, argFilePath)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVirtualAction(
				clientCtx.GetFromAddress().String(),
				argCode,
				newAction,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateVirtualAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-virtual-action [code] [new-action-file-path]",
		Short: "Update a virtual",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get value arguments
			argCode := args[0]
			argFilePath := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			newAction, err := nftmngrutils.ParseNewVirtualActionJSON(clientCtx.LegacyAmino, argFilePath)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateVirtual(
				clientCtx.GetFromAddress().String(),
				argCode,
				newAction,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteVirtualAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-virtual-action [code] [acion_name]",
		Short: "Delete a virtual",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCode := args[0]
			argActionName := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteVirtual(
				clientCtx.GetFromAddress().String(),
				argCode,
				argActionName,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
