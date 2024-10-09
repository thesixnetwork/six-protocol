// Copyright (C) 2024 SIX Network
// This file is part of the modified FeeMarket module from Ethermint, 
// and is licensed under the terms of the GNU Lesser General Public License v3
package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v010 "github.com/thesixnetwork/six-protocol/x/feemarket/migrations/v010"
	v011 "github.com/thesixnetwork/six-protocol/x/feemarket/migrations/v011"
	v012 "github.com/thesixnetwork/six-protocol/x/feemarket/migrations/v012"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}

// Migrate1to2 migrates the store from consensus version v1 to v2
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v010.MigrateStore(ctx, &m.keeper.paramSpace, m.keeper.storeKey)
}

// Migrate2to3 migrates the store from consensus version v2 to v3
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v011.MigrateStore(ctx, &m.keeper.paramSpace)
}

// Migrate3to4 migrates the store from consensus version v2 to v3
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	return v012.MigrateStore(ctx, &m.keeper.paramSpace)
}
