package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client"
)

var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","

	_ = flagPacketTimeoutTimestamp
	_ = listSeparator
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Aliases:                    []string{"meta"},
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateNFTSchema())
	cmd.AddCommand(CmdCreateMetadata())
	cmd.AddCommand(CmdPerformActionByAdmin())
	cmd.AddCommand(CmdAddAttribute())
	cmd.AddCommand(CmdAddAction())
	cmd.AddCommand(CmdSetBaseUri())
	cmd.AddCommand(CmdToggleAction())
	cmd.AddCommand(CmdChangeSchemaOwner())
	cmd.AddCommand(CmdResyncAttributes())
	cmd.AddCommand(CmdShowAttributes())
	cmd.AddCommand(CmdSetFeeConfig())
	cmd.AddCommand(CmdSetMintauth())
	cmd.AddCommand(CmdChangeOrgOwner())
	cmd.AddCommand(CmdSetUriRetrievalMethod())
	cmd.AddCommand(CmdSetOriginChain())
	cmd.AddCommand(CmdSetOriginContract())
	cmd.AddCommand(CmdSetAttributeOveriding())
	cmd.AddCommand(CmdSetMetadataFormat())
	cmd.AddCommand(CmdCreateActionExecutor())
	cmd.AddCommand(CmdUpdateActionExecutor())
	cmd.AddCommand(CmdDeleteActionExecutor())
	cmd.AddCommand(CmdUpdateSchemaAttribute())
	cmd.AddCommand(CmdUpdateAction())
	cmd.AddCommand(CmdCreateVirtualAction())
	cmd.AddCommand(CmdUpdateVirtualAction())
	cmd.AddCommand(CmdDeleteVirtualAction())
	cmd.AddCommand(CmdCreateVirtualSchema())
	cmd.AddCommand(CmdPerformVirtualAction())
	cmd.AddCommand(CmdVoteVirtualSchemaProposal())

	return cmd
}
