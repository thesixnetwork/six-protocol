package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

func CmdCreateActionSigner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-action-signer [base64-encoded-set-signer-action]",
		Short: "Create a new actionSigner",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			// Get value arguments
			argBase64SetSigner := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateActionSigner(
				clientCtx.GetFromAddress().String(),
				argBase64SetSigner,
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

func CmdUpdateActionSigner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-action-signer [base64-encoded-set-signer-action]",
		Short: "Update a actionSigner",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get value arguments
			argBase64SetSigner := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateActionSigner(
				clientCtx.GetFromAddress().String(),
				argBase64SetSigner,
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

func CmdDeleteActionSigner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-action-signer [base64-encoded-set-signer-action]",
		Short: "Delete a actionSigner",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBase64SetSigner := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteActionSigner(
				clientCtx.GetFromAddress().String(),
				argBase64SetSigner,
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
