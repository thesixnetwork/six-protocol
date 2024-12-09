package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func CmdCreateVirtualAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-virtual [code] [base64NewAction]",
		Short: "Create a new virtual",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			
			argCode := args[0]
			argBase64ActionStruct := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVirtualAction(
				clientCtx.GetFromAddress().String(),
				argCode,
				argBase64ActionStruct,
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
		Use:   "update-virtual [code] [base64ActionToUpdate]",
		Short: "Update a virtual",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get value arguments
			argCode := args[0]
			argBase64ActionStruct := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateVirtual(
				clientCtx.GetFromAddress().String(),
				argCode,
				argBase64ActionStruct,
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
		Use:   "delete-virtual [code] [acion_name]",
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
