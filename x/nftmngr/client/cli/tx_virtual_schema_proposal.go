package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nftmngrutils "github.com/thesixnetwork/six-protocol/x/nftmngr/client/utils"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// TODO: [chore] Combine enable/disable to change virtual schema.
func CmdCreateVirtualSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "virtual-schema-proposal [proposalType] [proposalFile]",
		Short: "Virtual Schema Proposal proposal type [create(0)/edit(1)]",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get value arguments
			argType := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposalType, err := parseProposalType(argType)
			if err != nil {
				return err
			}

			argFilePath := args[2]

			proposal, err := nftmngrutils.ParseProposalFile(clientCtx.LegacyAmino, argFilePath)
			if err != nil {
				return err
			}

			msg := types.NewMsgProposalVirtualSchema(
				clientCtx.GetFromAddress().String(),
				proposalType,
				proposal.VirtualSchema,
				proposal.Actions,
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

func parseProposalType(option string) (types.ProposalType, error) {
	optionLower := strings.ToLower(option)
	switch optionLower {
	case "create", "0":
		return types.ProposalType_CREATE, nil
	case "edit", "1":
		return types.ProposalType_EDIT, nil
	default:
		return types.ProposalType_CREATE, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid proposal type. Use 'create(0)/edit(1)'")
	}
}
