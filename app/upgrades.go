package app

import (
	"fmt"
	"time"

	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	nftmngrtypes "github.com/thesixnetwork/sixnft/x/nftmngr/types"
	nftoraclemoduletypes "github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

const UpgradeName = "v2.1.1"

func (app *App) VersionTrigger() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Deleted: []string{"evmsupport"},
		}
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}

func (app *App) RegisterUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler(UpgradeName, func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {

		schema_list := app.NftmngrKeeper.GetAllNFTSchemaV063(ctx)
		for _, schema := range schema_list {
			// get actions
			var v063_actions []*nftmngrtypes.Action
			for _, action := range schema.OnchainData.Actions {
				v063_actions = append(v063_actions, &nftmngrtypes.Action{
					Name:            action.Name,
					Desc:            action.Desc,
					Disable:         action.Disable,
					When:            action.When,
					Then:            action.Then,
					AllowedActioner: action.AllowedActioner,
					Params:          make([]*nftmngrtypes.ActionParams, 0),
				})
			}

			// build new schema verion
			new_schema := nftmngrtypes.NFTSchema{
				Code:            schema.Code,
				Name:            schema.Name,
				Owner:           schema.Owner,
				SystemActioners: schema.SystemActioners,
				OriginData:      schema.OriginData,
				OnchainData: &nftmngrtypes.OnChainData{
					RevealRequired:  schema.OnchainData.RevealRequired,
					RevealSecret:    schema.OnchainData.RevealSecret,
					NftAttributes:   schema.OnchainData.NftAttributes,
					TokenAttributes: schema.OnchainData.TokenAttributes,
					Actions:         v063_actions,
				},
				IsVerified:        schema.IsVerified,
				MintAuthorization: schema.MintAuthorization,
			}

			app.NftmngrKeeper.SetNFTSchema(ctx, new_schema)
		}

		// set nft duration
		var oracle_params nftoraclemoduletypes.Params
		oracle_params.MintRequestActiveDuration = 120 * time.Second
		oracle_params.ActionRequestActiveDuration = 120 * time.Second
		oracle_params.VerifyRequestActiveDuration = 120 * time.Second
		oracle_params.ActionSignerActiveDuration = 30 * (24 * time.Hour)
		app.NftoracleKeeper.SetParams(ctx, oracle_params)

		return app.mm.RunMigrations(ctx, app.configurator, vm)
	})
}
