package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group nftmngr queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Aliases:                    []string{"meta"},
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListNFTSchema())
	cmd.AddCommand(CmdShowNFTSchema())
	cmd.AddCommand(CmdListNftData())
	cmd.AddCommand(CmdShowNftData())
	cmd.AddCommand(CmdListActionByRefId())
	cmd.AddCommand(CmdShowActionByRefId())
	cmd.AddCommand(CmdListOrganization())
	cmd.AddCommand(CmdShowOrganization())
	cmd.AddCommand(CmdShowNftCollection())
	cmd.AddCommand(CmdListNFTSchemaByContract())
	cmd.AddCommand(CmdShowNFTSchemaByContract())
	cmd.AddCommand(CmdShowNFTFeeConfig())
	cmd.AddCommand(CmdShowNFTFeeBalance())
	cmd.AddCommand(CmdListMetadataCreator())
	cmd.AddCommand(CmdShowMetadataCreator())
	cmd.AddCommand(CmdListActionExecutor())
	cmd.AddCommand(CmdShowActionExecutor())
	cmd.AddCommand(CmdListSchemaAttribute())
	cmd.AddCommand(CmdShowSchemaAttribute())
	cmd.AddCommand(CmdListActionOfSchema())
	cmd.AddCommand(CmdShowActionOfSchema())
	cmd.AddCommand(CmdListExecutorOfSchema())
	cmd.AddCommand(CmdShowExecutorOfSchema())

	cmd.AddCommand(CmdListVirtualAction())
	cmd.AddCommand(CmdShowVirtualAction())
	cmd.AddCommand(CmdListVirtualSchema())
	cmd.AddCommand(CmdShowVirtualSchema())
	cmd.AddCommand(CmdListDisableVirtualSchema())
	cmd.AddCommand(CmdShowDisableVirtualSchema())
	cmd.AddCommand(CmdShowVirtualSchemaProposal())
	cmd.AddCommand(CmdListVirtualSchemaProposal())
	cmd.AddCommand(CmdListActiveDislabeVirtualSchemaProposal())
	cmd.AddCommand(CmdShowActiveDislabeVirtualSchemaProposal())
	cmd.AddCommand(CmdListInactiveDisableVirtualSchemaProposal())
	cmd.AddCommand(CmdShowInactiveDisableVirtualSchemaProposal())
	// this line is used by starport scaffolding # 1

	return cmd
}
