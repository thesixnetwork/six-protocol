package app

import (
	"fmt"
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	evmsupportmoduletypes "github.com/thesixnetwork/sixnft/x/evmsupport/types"

	evmtypes "github.com/evmos/ethermint/x/evm/types"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
)

const UpgradeName = "v3.0.0"

func (app *App) VersionTrigger() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}
	fmt.Println("##########upgradeInfo", upgradeInfo)
	if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Added: []string{evmtypes.StoreKey, feemarkettypes.StoreKey},
			Deleted: []string{evmsupportmoduletypes.StoreKey},
		}
		fmt.Println("##########storeUpgrades", storeUpgrades)
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
	fmt.Println("########## end of if")
}

func (app *App) RegisterUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler(UpgradeName, func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {

		var evm_params evmtypes.Params
		evm_params.EvmDenom = "usix"
		app.EVMKeeper.SetParams(ctx, evm_params)

		return app.mm.RunMigrations(ctx, app.configurator, vm)
	})
}
