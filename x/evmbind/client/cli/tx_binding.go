package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/evmbind/types"
)

func CmdCreateBinding() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-binding [eth-address] [eth-signature] [sign-message]",
		Short: "Create a new binding",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexEthAddress := args[0]

			// Get value arguments
			argEthSignature := args[1]
			argSignMessage := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBinding(
				clientCtx.GetFromAddress().String(),
				indexEthAddress,
				argEthSignature,
				argSignMessage,
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

func CmdUpdateBinding() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-binding [eth-address] [eth-signature] [sign-message]",
		Short: "Update a binding",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexEthAddress := args[0]

			// Get value arguments
			argEthSignature := args[1]
			argSignMessage := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBinding(
				clientCtx.GetFromAddress().String(),
				indexEthAddress,
				argEthSignature,
				argSignMessage,
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

func CmdDeleteBinding() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-binding [eth-address]",
		Short: "Delete a binding",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexEthAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteBinding(
				clientCtx.GetFromAddress().String(),
				indexEthAddress,
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
