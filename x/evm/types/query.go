// Copyright (C) 2024 SIX Network
// This file is part of the modified EVM module from Ethermint, 
// and is licensed under the terms of the GNU Lesser General Public License v3
package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// UnpackInterfaces implements UnpackInterfacesMesssage.UnpackInterfaces
func (m QueryTraceTxRequest) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, msg := range m.Predecessors {
		if err := msg.UnpackInterfaces(unpacker); err != nil {
			return err
		}
	}
	return m.Msg.UnpackInterfaces(unpacker)
}

func (m QueryTraceBlockRequest) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, msg := range m.Txs {
		if err := msg.UnpackInterfaces(unpacker); err != nil {
			return err
		}
	}
	return nil
}
