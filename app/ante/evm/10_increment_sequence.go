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

		if unsafeUnorderedTx && ctx.IsCheckTx() {
			// Unordered mode, mempool admission (CheckTx / ReCheckTx).
			//
			// Accept any current-or-future nonce so that several txs from the
			// same sender that arrive out of order (e.g. a foundry "fast"
			// deploy that broadcasts nonces 0..N concurrently) are all admitted
			// to the mempool. The app uses a PriorityNonce mempool with an EVM
			// signer extractor, so it re-orders each sender's txs by nonce
			// before a block is built — the strict check below then passes in
			// DeliverTx.
			//
			// Reject only a nonce BELOW the committed sequence: that tx has
			// already been executed, so this keeps replays/stale txs out of the
			// mempool. Crucially we do NOT advance the sequence here — advancing
			// during admission would turn it into a per-round counter and would
			// wrongly reject a lower nonce that legitimately arrives later.
			if tx.Nonce() < nonce {
				return errorsmod.Wrapf(
					errortypes.ErrInvalidSequence,
					"nonce too low; got %d, expected >= %d", tx.Nonce(), nonce,
				)
			}
			continue
		}

		// Ordered mode (all contexts), or unordered mode during block execution
		// (DeliverTx): enforce an exact match against the account sequence. In
		// unordered mode the proposer emits each sender's txs in nonce order, so
		// legitimately queued txs pass here; a re-broadcast past tx has a nonce
		// below the advanced sequence and is rejected. This exact-match check at
		// execution time — where tx ordering is deterministic — is the actual
		// replay guard. It also accepts multiple messages from one sender in a
		// single tx because the sequence is incremented per message below.
		if tx.Nonce() != nonce {
			return errorsmod.Wrapf(
				errortypes.ErrInvalidSequence,
				"invalid nonce; got %d, expected %d", tx.Nonce(), nonce,
			)
		}

		// Advance the sequence monotonically by one per message. Using nonce+1
		// (the account's current sequence, never the tx nonce) guarantees the
		// stored sequence only ever moves forward and can never be rolled back.
		if err := acc.SetSequence(nonce + 1); err != nil {
			return errorsmod.Wrapf(err, "failed to set sequence to %d", nonce+1)
		}

		ak.SetAccount(ctx, acc)
	}

	return nil
}
