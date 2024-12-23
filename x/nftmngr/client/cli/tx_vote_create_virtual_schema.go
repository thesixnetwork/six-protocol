package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

var _ = strconv.Itoa(0)

func CmdVoteCreateVirtualSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote-create-virtual-schema [virtual-nft-schema-code] [your-nft-schema-code] [yes(y)  no(n)]",
		Short: "Broadcast message voteCreateVirtualSchema",
		Args:  cobra.ExactArgs(3),
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

			msg := types.NewMsgVoteCreateVirtualSchema(
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
	fmt.Println(option)
	switch option {
	case "YES", "Y", "yes", "y":
		return types.RegistryStatus_ACCEPT, nil
	case "NO", "N", "no", "n":
		return types.RegistryStatus_REJECT, nil
	default:
		return types.RegistryStatus_REJECT, fmt.Errorf("invalid vote option. Use 'YES' or 'NO'")
	}
}
