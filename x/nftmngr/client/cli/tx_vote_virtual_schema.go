package cli

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	errormod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ = strconv.Itoa(0)

func CmdVoteVirtualSchemaProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "vote-virtual-schema [proposalId] [nft-schema-code] [yes(y)  no(n)]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			proposalId := args[0]
			argNftSchemaName := args[1]
			argOption := args[2]

			// Validate input option
			selectedVote, err := parseVoteOption(argOption)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgVoteVirtualSchemaProposal(
				clientCtx.GetFromAddress().String(),
				proposalId,
				argNftSchemaName,
				selectedVote,
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

func parseVoteOption(option string) (types.RegistryStatus, error) {
	switch option {
	case "YES", "Y", "yes", "y":
		return types.RegistryStatus_ACCEPT, nil
	case "NO", "N", "no", "n":
		return types.RegistryStatus_REJECT, nil
	default:
		return types.RegistryStatus_REJECT, errormod.Wrap(sdkerrors.ErrInvalidRequest, "invalid vote option. Use 'YES' or 'NO'")
	}
}
