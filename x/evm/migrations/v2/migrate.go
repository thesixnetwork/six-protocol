// Copyright (C) 2024 SIX Network
// This file is part of the modified EVM module from Ethermint, 
// and is licensed under the terms of the GNU Lesser General Public License v3
package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/thesixnetwork/six-protocol/x/evm/types"
)

// MigrateStore sets the default AllowUnprotectedTxs parameter.
func MigrateStore(ctx sdk.Context, paramstore *paramtypes.Subspace) error {
	if !paramstore.HasKeyTable() {
		ps := paramstore.WithKeyTable(types.ParamKeyTable())
		paramstore = &ps
	}

	// add RejectUnprotected
	paramstore.Set(ctx, types.ParamStoreKeyAllowUnprotectedTxs, types.DefaultAllowUnprotectedTxs)
	return nil
}
