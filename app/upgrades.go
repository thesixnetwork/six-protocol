package app

import (
	"fmt"
	"time"

	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authz "github.com/cosmos/cosmos-sdk/x/authz"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	evmsupportmoduletypes "github.com/thesixnetwork/sixnft/x/evmsupport/types"
	nftadminmoduletypes "github.com/thesixnetwork/sixnft/x/nftadmin/types"
	nftmngrmoduletypes "github.com/thesixnetwork/sixnft/x/nftmngr/types"
	nftoraclemoduletypes "github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

// const UpgradeName = "v2.0.0"

func (app *App) VersionTrigger() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}
	fmt.Println("##########upgradeInfo", upgradeInfo)
	if upgradeInfo.Name == "v2.0.0" && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Added: []string{nftmngrmoduletypes.StoreKey, evmsupportmoduletypes.StoreKey, nftoraclemoduletypes.StoreKey, nftadminmoduletypes.StoreKey, authz.ModuleName},
		}
		fmt.Println("##########storeUpgrades", storeUpgrades)
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
	fmt.Println("########## end of if")
}

func (app *App) RegisterUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler("v2.0.0", func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		// set state root admin
		var admin nftadminmoduletypes.Authorization
		super_admin, _ := app.ProtocoladminKeeper.GetGroup(ctx, "super.admin")
		admin.RootAdmin = super_admin.GetOwner()
		app.NftadminKeeper.SetAuthorization(ctx, admin)

		// set nftngr nft_fee_config
		var nft_fee_config nftmngrmoduletypes.NFTFeeConfig
		var fee_config nftmngrmoduletypes.FeeConfig
		fee_config.FeeAmount = "20000000000usix"
		fee_config.FeeDistributions = make([]*nftmngrmoduletypes.FeeDistribution, 0)
		fee_config.FeeDistributions = append(fee_config.FeeDistributions, &nftmngrmoduletypes.FeeDistribution{
			Method:  nftmngrmoduletypes.FeeDistributionMethod_BURN,
			Portion: 0.1,
		})
		fee_config.FeeDistributions = append(fee_config.FeeDistributions, &nftmngrmoduletypes.FeeDistribution{
			Method:  nftmngrmoduletypes.FeeDistributionMethod_REWARD_POOL,
			Portion: 0.9,
		})

		nft_fee_config.SchemaFee = &fee_config
		app.NftmngrKeeper.SetNFTFeeConfig(ctx, nft_fee_config)

		// set nft duration
		var oracle_params nftoraclemoduletypes.Params
		oracle_params.MintRequestActiveDuration = 120 * time.Second
		oracle_params.ActionRequestActiveDuration = 120 * time.Second
		oracle_params.VerifyRequestActiveDuration = 120 * time.Second
		app.NftoracleKeeper.SetParams(ctx, oracle_params)

		// set oracle minimum_confirmations
		app.NftoracleKeeper.SetOracleConfig(ctx, nftoraclemoduletypes.OracleConfig{
			MinimumConfirmation: 4,
		})

		return app.mm.RunMigrations(ctx, app.configurator, vm)
	})
}
