package app

import (
	"context"
	"fmt"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	storetypes "cosmossdk.io/store/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	evidencetypes "cosmossdk.io/x/evidence/types"
	ratelimittypes "github.com/cosmos/ibc-apps/modules/rate-limiting/v8/types"
	circuittypes "cosmossdk.io/x/circuit/types"
)

const UpgradeName = "v4.0.0"

func (app *App) RegisterUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler(UpgradeName, func(ctx context.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		return app.ModuleManager.RunMigrations(ctx, app.configurator, vm)
	})


	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}
	if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added:   []string{
				crisistypes.StoreKey,
				consensusparamkeeper.StoreKey,
				circuittypes.StoreKey,
				evidencetypes.StoreKey,

				ibcexported.StoreKey,
				icahosttypes.StoreKey,
				ibcfeetypes.StoreKey,
				icacontrollertypes.StoreKey,

				ratelimittypes.StoreKey,
			},
			Deleted: []string{},
		}
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
