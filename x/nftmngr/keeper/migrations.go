package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
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

func (m Migrator) NoMigration(ctx sdk.Context) error {
	return nil
}

func (m Migrator) MigrateStore(ctx sdk.Context) error {
	m.keeper.SetParams(ctx, types.DefaultParams())
	return nil
}