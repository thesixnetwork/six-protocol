package app

import (
	"fmt"

	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	// nftmngrtypes "github.com/thesixnetwork/sixnft/x/nftmngr/types"
	tokenmngrmoduletypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
	nftoraclemoduletypes "github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

const UpgradeName = "v3.0.0"

func (app *App) VersionTrigger() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{}
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}

func (app *App) RegisterUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler(UpgradeName, func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		action_requests := app.NftoracleKeeper.GetAllActionRequestV603(ctx)

		// bug from previous version
		for _, action_request := range action_requests {
			var action_request_v703 nftoraclemoduletypes.ActionOracleRequest
			action_request_v703 =  nftoraclemoduletypes.ActionOracleRequest{
				Id:          action_request.Id,
				NftSchemaCode	: action_request.NftSchemaCode,
				TokenId		: action_request.TokenId,
				Action		: action_request.Action,
				Params		: make([]*nftoraclemoduletypes.ActionParameter, 0),
				Caller		: action_request.Caller,
				RefId		: action_request.RefId,
				RequiredConfirm	: action_request.RequiredConfirm,
				Status		: action_request.Status,
				CurrentConfirm	: action_request.CurrentConfirm,
				Confirmers	: action_request.Confirmers,
				CreatedAt	: action_request.CreatedAt,
				ValidUntil	: action_request.ValidUntil,
				DataHashes	: action_request.DataHashes,
				ExpiredHeight	: action_request.ExpiredHeight,
				ExecutionErrorMessage: action_request.ExecutionErrorMessage,
			}
			app.NftoracleKeeper.SetActionRequest(ctx, action_request_v703)
		}

		// Token-burn tokens
		token_burn := app.TokenmngrKeeper.GetAllTokenBurnV202(ctx)
		for _, burn := range token_burn {
			var burn_v203 tokenmngrmoduletypes.TokenBurn
			
			burn_v203 = tokenmngrmoduletypes.TokenBurn{
				Amount: sdk.NewCoin(burn.Token, sdk.NewInt(int64(burn.Amount))),
			}
			app.TokenmngrKeeper.SetTokenBurn(ctx, burn_v203)
		}

		// Burns
		burns := app.TokenmngrKeeper.GetAllBurnV202(ctx)
		for _, burn := range burns {
			var new_ver_burns tokenmngrmoduletypes.Burn
			new_ver_burns = tokenmngrmoduletypes.Burn{
				Id: burn.Id,
				Creator: burn.Creator,
				Amount: sdk.NewCoin(burn.Token, sdk.NewInt(int64(burn.Amount))),
			}
			app.TokenmngrKeeper.UpdateBurn(ctx, new_ver_burns)
		}

		return app.mm.RunMigrations(ctx, app.configurator, vm)
	})
}
