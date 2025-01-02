package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func CmdCreateActionExecutor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-action-executor [nft-schema-code] [executor-address]",
		Short: "Create a new actionExecutor",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexNftSchemaCode := args[0]
			indexExecutorAddress := args[1]

			// Get value arguments

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateActionExecutor(
				clientCtx.GetFromAddress().String(),
				indexNftSchemaCode,
				indexExecutorAddress,
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

func CmdUpdateActionExecutor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-action-executor [nft-schema-code] [executor-address]",
		Short: "Update a actionExecutor",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexNftSchemaCode := args[0]
			indexExecutorAddress := args[1]

			// Get value arguments

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateActionExecutor(
				clientCtx.GetFromAddress().String(),
				indexNftSchemaCode,
				indexExecutorAddress,
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

func CmdDeleteActionExecutor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-action-executor [nft-schema-code] [executor-address]",
		Short: "Delete a actionExecutor",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexNftSchemaCode := args[0]
			indexExecutorAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteActionExecutor(
				clientCtx.GetFromAddress().String(),
				indexNftSchemaCode,
				indexExecutorAddress,
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
