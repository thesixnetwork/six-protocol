package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1 "github.com/thesixnetwork/six-protocol/x/nftadmin/migrations/v1"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v1.MigrateStore(ctx, m.keeper.storeService, m.keeper.cdc)
}
