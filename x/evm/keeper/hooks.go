// Copyright (C) 2024 SIX Network
// This file is part of the modified EVM module from Ethermint (https://github.com/evmos/ethermint),
// and is licensed under the terms of the GNU Lesser General Public License v3
package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/thesixnetwork/six-protocol/x/evm/types"
)

var _ types.EvmHooks = MultiEvmHooks{}

// MultiEvmHooks combine multiple evm hooks, all hook functions are run in array sequence
type MultiEvmHooks []types.EvmHooks

// NewMultiEvmHooks combine multiple evm hooks
func NewMultiEvmHooks(hooks ...types.EvmHooks) MultiEvmHooks {
	return hooks
}

// PostTxProcessing delegate the call to underlying hooks
func (mh MultiEvmHooks) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	for i := range mh {
		if err := mh[i].PostTxProcessing(ctx, msg, receipt); err != nil {
			return sdkerrors.Wrapf(err, "EVM hook %T failed", mh[i])
		}
	}
	return nil
}
