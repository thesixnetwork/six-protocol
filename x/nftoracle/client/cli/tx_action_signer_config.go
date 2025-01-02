package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

func CmdCreateActionSignerConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-action-signer-config [chain] [contract-address]",
		Short: "Create a new actionSignerConfig",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexChain := args[0]

			// Get value arguments
			argContractAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateActionSignerConfig(
				clientCtx.GetFromAddress().String(),
				indexChain,
				argContractAddress,
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

func CmdUpdateActionSignerConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-action-signer-config [chain] [contract-address]",
		Short: "Update a actionSignerConfig",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexChain := args[0]

			// Get value arguments
			argContractAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateActionSignerConfig(
				clientCtx.GetFromAddress().String(),
				indexChain,
				argContractAddress,
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

func CmdDeleteActionSignerConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-action-signer-config [chain]",
		Short: "Delete a actionSignerConfig",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexChain := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteActionSignerConfig(
				clientCtx.GetFromAddress().String(),
				indexChain,
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
