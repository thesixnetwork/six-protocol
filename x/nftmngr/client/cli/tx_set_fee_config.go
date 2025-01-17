package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	nftmngrutils "github.com/thesixnetwork/six-protocol/x/nftmngr/client/utils"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

var _ = strconv.Itoa(0)

func CmdSetFeeConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-fee-config [file-path]",
		Short: "To set fee config",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFeeConfigFilePath := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			feeConfig, err := nftmngrutils.ParseFeeConfigJSON(clientCtx.LegacyAmino, argFeeConfigFilePath)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetFeeConfig(
				clientCtx.GetFromAddress().String(),
				feeConfig,
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

func parseFeeSubject(option string) (types.FeeSubject, error) {
	switch option {
	case "0":
		return types.FeeSubject_CREATE_NFT_SCHEMA, nil
	default:
		return types.FeeSubject_CREATE_NFT_SCHEMA, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid subject. Use 0 or 1")
	}
}
