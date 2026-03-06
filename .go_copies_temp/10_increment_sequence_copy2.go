// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package evm

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"

	evmtypes "github.com/evmos/evmos/v20/x/evm/types"
)

func IncrementNonce(ctx sdk.Context, ak evmtypes.AccountKeeper, acc sdk.AccountI, tx sdk.Tx, unsafeUnorderedTx bool) error {
	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return errorsmod.Wrapf(errortypes.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}

		tx := msgEthTx.AsTransaction()

		// increase sequence of sender
		from := msgEthTx.GetFrom()
		if acc == nil {
			return errorsmod.Wrapf(
				errortypes.ErrUnknownAddress,
				"account %s is nil", common.BytesToAddress(from.Bytes()),
			)
		}
		nonce := acc.GetSequence()

		if !unsafeUnorderedTx {
			// we merged the nonce verification to nonce increment, so when tx includes multiple messages
			// with same sender, they'll be accepted.
			if tx.Nonce() != nonce {
				return errorsmod.Wrapf(
					errortypes.ErrInvalidSequence,
					"invalid nonce; got %d, expected %d", tx.Nonce(), nonce,
				)
			}
		}

		if err := acc.SetSequence(nonce + 1); err != nil {
			return errorsmod.Wrapf(err, "failed to set sequence to %d", acc.GetSequence()+1)
		}

		ak.SetAccount(ctx, acc)
	}

	return nil
}
