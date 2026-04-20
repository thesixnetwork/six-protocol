package cli

import (
	"fmt"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	evmostypes "github.com/evmos/evmos/v20/types"
	"github.com/spf13/cobra"

	"github.com/thesixnetwork/six-protocol/v4/x/erc20/types"
)

// NewTxCmd returns root CLI handler for erc20 transaction subcommands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "erc20 transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(NewConvertERC20Cmd())
	return txCmd
}

// NewConvertERC20Cmd returns a command to convert an ERC20 balance into
// the corresponding Cosmos SDK coin.
//
// Usage: sixd tx erc20 convert-erc20 CONTRACT_ADDRESS AMOUNT [RECEIVER]
func NewConvertERC20Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-erc20 CONTRACT_ADDRESS AMOUNT [RECEIVER]",
		Short: "Convert an ERC20 token balance to a Cosmos coin",
		Long: `Convert an ERC20 token balance to the registered Cosmos coin.

CONTRACT_ADDRESS  hex address of the registered ERC20 contract (0x...)
AMOUNT            integer amount of ERC20 token base units to convert
RECEIVER          (optional) bech32 cosmos address to receive the coins;
                  defaults to the --from sender address`,
		Args: cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			contract := args[0]
			if err := evmostypes.ValidateAddress(contract); err != nil {
				return fmt.Errorf("invalid ERC20 contract address: %w", err)
			}

			amount, ok := math.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("invalid amount %q: must be a positive integer", args[1])
			}

			from := common.BytesToAddress(cliCtx.GetFromAddress().Bytes())

			receiver := cliCtx.GetFromAddress()
			if len(args) == 3 {
				receiver, err = sdk.AccAddressFromBech32(args[2])
				if err != nil {
					return fmt.Errorf("invalid receiver address: %w", err)
				}
			}

			msg := &types.MsgConvertERC20{
				ContractAddress: contract,
				Amount:          amount,
				Receiver:        receiver.String(),
				Sender:          from.Hex(),
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
