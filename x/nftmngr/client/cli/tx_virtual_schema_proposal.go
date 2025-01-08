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

// TODO:: TEST(VirtualSchema)
func CmdCreateVirtualSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "virtual-schema-proposal [schemaCode] [proposalType] [proposalFile]",
		Short: "Virtual Schema Proposal proposal type [create(0)/enable(1)/disable(2)/delete(3)]",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get value arguments
			argCode := args[0]
			argType := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposalType, err := parseProposalType(argType)
			if err != nil {
				return err
			}

			request := make([]types.VirtualSchemaRegistryRequest, 0)

			if proposalType == types.ProposalType_CREATE {

				argFilePath := args[2]

				proposal, err := nftmngrutils.ParseVirtualSchemaRegistryRequestJSON(clientCtx.LegacyAmino, argFilePath)
				if err != nil {
					return err
				}

				for _, req := range proposal {
					registry := types.NewVirtualSchemaRegistryRequest(req.Code, req.SharedAttributes)
					request = append(request, *registry)
				}
			}

			msg := types.NewMsgProposalVirtualSchema(
				clientCtx.GetFromAddress().String(),
				argCode,
				proposalType,
				request,
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
	case "enable", "1":
		return types.ProposalType_ENABLE, nil
	case "disable", "2":
		return types.ProposalType_DISABLE, nil
	default:
		return types.ProposalType_DISABLE, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid proposal type. Use 'create(0)/enable(1)/disable(2)/delete(3)'")
	}
}
