package app

import (
	"fmt"

	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktype "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	tokenmngrtypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
)

const UpgradeName = "v3.0.0"

func (app *App) VersionTrigger() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Added: []string{evmtypes.ModuleName,feemarkettypes.ModuleName},
		}
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}

func (app *App) RegisterUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler(UpgradeName, func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		// set new denom metadata
		app.BankKeeper.SetDenomMetaData(ctx, banktype.Metadata{
			Description: "The native evm token of the SIX Protocol",
			Base:        "asix",
			Display:     "asix",
			Symbol:      "asix",
			Name: "six evm token",
			DenomUnits: []*banktype.DenomUnit{
				{
					Denom:    "asix",
					Exponent: 0,
					Aliases:  []string{"attosix"},
				},
				{
					Denom:    "usix",
					Exponent: 12,
					Aliases:  []string{"microsix"},
				},
				{
					Denom:    "msix",
					Exponent: 15,
					Aliases:  []string{"millisix"},
				},
				{
					Denom:    "six",
					Exponent: 18,
					Aliases:  []string{"six"},
				},
			},

		})

		// super admin address
		token_admin, _ := app.ProtocoladminKeeper.GetGroup(ctx, "token.admin")
		super_admin, _ := app.ProtocoladminKeeper.GetGroup(ctx, "super.admin")
		app.TokenmngrKeeper.SetToken(ctx, tokenmngrtypes.Token{
			Name: "asix",
			Base: "asix",
			MaxSupply: 0,
			Mintee: token_admin.Owner,
			Creator: super_admin.Owner,

		})

		// add mintperm
		app.TokenmngrKeeper.SetMintperm(ctx, tokenmngrtypes.Mintperm{
			Token: "asix",
			Address: token_admin.Owner,
			Creator: super_admin.Owner,
		})

		evm_param := evmtypes.DefaultParams()
		evm_param.EvmDenom = "asix"
		app.EVMKeeper.SetParams(ctx, evm_param)
		return app.mm.RunMigrations(ctx, app.configurator, vm)
	})
}
