package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group nftoracle queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListMintRequest())
	cmd.AddCommand(CmdShowMintRequest())
	cmd.AddCommand(CmdListActionRequest())
	cmd.AddCommand(CmdShowActionRequest())
	cmd.AddCommand(CmdListCollectionOwnerRequest())
	cmd.AddCommand(CmdShowCollectionOwnerRequest())
	cmd.AddCommand(CmdShowOracleConfig())
	cmd.AddCommand(CmdListActionSigner())
	cmd.AddCommand(CmdShowBindedSigner())
	cmd.AddCommand(CmdShowActionSigner())
	cmd.AddCommand(CmdListActionSignerConfig())
	cmd.AddCommand(CmdShowActionSignerConfig())
	cmd.AddCommand(CmdListSyncActionSigner())
	cmd.AddCommand(CmdShowSyncActionSigner())
	// this line is used by starport scaffolding # 1

	return cmd
}
