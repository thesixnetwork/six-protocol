package app

import (
	"fmt"
	"strings"
	"time"

	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	banktype "github.com/cosmos/cosmos-sdk/x/bank/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	evmtypes "github.com/evmos/ethermint/x/evm/types"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
	erc20types "github.com/evmos/evmos/v6/x/erc20/types"
	tokenmngrmoduletypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
	tokenmngrtypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
	nftmngrtypes "github.com/thesixnetwork/sixnft/x/nftmngr/types"
	nftoraclemoduletypes "github.com/thesixnetwork/sixnft/x/nftoracle/types"
)

const UpgradeName = "v3.1.0"
// CHAIN ID will save pointer of string
var CHAINID string

func (app *App) RegisterUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler(UpgradeName, func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		if ctx.ChainID() == "fivenet" {
			return app.mm.RunMigrations(ctx, app.configurator, vm)
		}

		// ** Migrate store keys from v2.1.0 to v2.2.0 **
		// * Module Tokenmngr *

		// Change token struct
		tokens := app.TokenmngrKeeper.GetAllTokenV202(ctx)
		for _, token := range tokens {
			var token_v300 tokenmngrmoduletypes.Token
			max_supply_v202_to_v300 := sdk.NewIntFromUint64(token.MaxSupply)
			new_coin := sdk.NewCoin(token.Name, max_supply_v202_to_v300)
			token_v300 = tokenmngrmoduletypes.Token{
				Name:      token.Name,
				Base:      token.Base,
				MaxSupply: new_coin,
				Mintee:    token.Mintee,
				Creator:   token.Creator,
			}
			app.TokenmngrKeeper.SetToken(ctx, token_v300)
		}

		// Token-burn tokens
		token_burn := app.TokenmngrKeeper.GetAllTokenBurnV202(ctx)
		for _, burn := range token_burn {
			var burn_v300 tokenmngrmoduletypes.TokenBurn
			burn_v300 = tokenmngrmoduletypes.TokenBurn{
				Amount: sdk.NewCoin(burn.Token, sdk.NewInt(int64(burn.Amount))),
			}
			app.TokenmngrKeeper.SetTokenBurn(ctx, burn_v300)
		}

		// Burns
		burns := app.TokenmngrKeeper.GetAllBurnV202(ctx)
		list_burns := make([]tokenmngrmoduletypes.Burn, 0)
		for _, burn := range burns {
			var new_ver_burns tokenmngrmoduletypes.Burn
			new_ver_burns = tokenmngrmoduletypes.Burn{
				Id:      burn.Id,
				Creator: burn.Creator,
				Amount:  sdk.NewCoin(burn.Token, sdk.NewInt(int64(burn.Amount))),
			}
			list_burns = append(list_burns, new_ver_burns)
			app.TokenmngrKeeper.UpdateBurn(ctx, new_ver_burns)
		}
		// Set list burns
		app.TokenmngrKeeper.SetBurns(ctx, list_burns)

		// * Module NFTOracle *

		// bug from previous version
		action_requests, err := app.NftoracleKeeper.GetAllActionRequestV603(ctx)
		if err != nil {
			action_requests := app.NftoracleKeeper.GetAllActionRequest(ctx)
			for _, action_request := range action_requests {
				var action_request_v703 nftoraclemoduletypes.ActionOracleRequest
				action_request_v703 = nftoraclemoduletypes.ActionOracleRequest{
					Id:                    action_request.Id,
					NftSchemaCode:         action_request.NftSchemaCode,
					TokenId:               action_request.TokenId,
					Action:                action_request.Action,
					Params:                make([]*nftoraclemoduletypes.ActionParameter, 0),
					Caller:                action_request.Caller,
					RefId:                 action_request.RefId,
					RequiredConfirm:       action_request.RequiredConfirm,
					Status:                action_request.Status,
					CurrentConfirm:        action_request.CurrentConfirm,
					Confirmers:            action_request.Confirmers,
					CreatedAt:             action_request.CreatedAt,
					ValidUntil:            action_request.ValidUntil,
					DataHashes:            action_request.DataHashes,
					ExpiredHeight:         action_request.ExpiredHeight,
					ExecutionErrorMessage: action_request.ExecutionErrorMessage,
				}
				app.NftoracleKeeper.SetActionRequest(ctx, action_request_v703)
			}
		} else {
			for _, action_request := range action_requests {
				var action_request_v703 nftoraclemoduletypes.ActionOracleRequest
				action_request_v703 = nftoraclemoduletypes.ActionOracleRequest{
					Id:                    action_request.Id,
					NftSchemaCode:         action_request.NftSchemaCode,
					TokenId:               action_request.TokenId,
					Action:                action_request.Action,
					Params:                make([]*nftoraclemoduletypes.ActionParameter, 0),
					Caller:                action_request.Caller,
					RefId:                 action_request.RefId,
					RequiredConfirm:       action_request.RequiredConfirm,
					Status:                action_request.Status,
					CurrentConfirm:        action_request.CurrentConfirm,
					Confirmers:            action_request.Confirmers,
					CreatedAt:             action_request.CreatedAt,
					ValidUntil:            action_request.ValidUntil,
					DataHashes:            action_request.DataHashes,
					ExpiredHeight:         action_request.ExpiredHeight,
					ExecutionErrorMessage: action_request.ExecutionErrorMessage,
				}
				app.NftoracleKeeper.SetActionRequest(ctx, action_request_v703)
			}
		}

		// ActionSigners => Append to new store
		action_signers := app.NftoracleKeeper.GetAllActionSigner(ctx)
		for _, action_signer := range action_signers {

			// query binded signer
			binded_signer, found := app.NftoracleKeeper.GetBindedSigner(ctx, action_signer.OwnerAddress)
			if !found {
				var list_signer []*nftoraclemoduletypes.XSetSignerParams
				// var binded_signer nftoraclemoduletypes.BindedSigner
				binded_signer = nftoraclemoduletypes.BindedSigner{
					OwnerAddress: action_signer.OwnerAddress,
					Signers: append(list_signer, &nftoraclemoduletypes.XSetSignerParams{
						ActorAddress: action_signer.ActorAddress,
						ExpiredAt:    action_signer.ExpiredAt,
					}),
				}
				app.NftoracleKeeper.SetBindedSigner(ctx, binded_signer)
			} else {
				binded_signer = nftoraclemoduletypes.BindedSigner{
					OwnerAddress: action_signer.OwnerAddress,
					Signers: append(binded_signer.Signers, &nftoraclemoduletypes.XSetSignerParams{
						ActorAddress: action_signer.ActorAddress,
						ExpiredAt:    action_signer.ExpiredAt,
					}),
				}
				app.NftoracleKeeper.SetBindedSigner(ctx, binded_signer)
			}

		}

		// * Module NFTManager *

		// NFTSchema and NFTCollection
		schema_list := app.NftmngrKeeper.GetAllNFTSchemaV072(ctx)
		for _, schema := range schema_list {

			var new_nftattribute []*nftmngrtypes.AttributeDefinition
			for _, nftattribute := range schema.OnchainData.NftAttributes {
				new_nftattribute = append(new_nftattribute, &nftmngrtypes.AttributeDefinition{
					Name:                nftattribute.Name,
					DataType:            nftattribute.DataType,
					Required:            nftattribute.Required,
					DisplayValueField:   nftattribute.DisplayValueField,
					DisplayOption:       nftattribute.DisplayOption,
					DefaultMintValue:    nftattribute.DefaultMintValue,
					HiddenOveride:       true,
					HiddenToMarketplace: nftattribute.HiddenToMarketplace,
					Index:               nftattribute.Index,
				})
			}

			var new_tokenattribute []*nftmngrtypes.AttributeDefinition
			for _, tokenattribute := range schema.OnchainData.TokenAttributes {
				new_tokenattribute = append(new_tokenattribute, &nftmngrtypes.AttributeDefinition{
					Name:                tokenattribute.Name,
					DataType:            tokenattribute.DataType,
					Required:            tokenattribute.Required,
					DisplayValueField:   tokenattribute.DisplayValueField,
					DisplayOption:       tokenattribute.DisplayOption,
					DefaultMintValue:    tokenattribute.DefaultMintValue,
					HiddenOveride:       true,
					HiddenToMarketplace: tokenattribute.HiddenToMarketplace,
					Index:               tokenattribute.Index,
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
					NftAttributes:   new_nftattribute,
					TokenAttributes: new_tokenattribute,
					Actions:         schema.OnchainData.Actions,
				},
				IsVerified:        schema.IsVerified,
				MintAuthorization: schema.MintAuthorization,
			}

			// replace string in action[num].then[num]
			for i, action := range new_schema.OnchainData.Actions {
				for j, then := range action.Then {
					new_schema.OnchainData.Actions[i].Then[j] = strings.ReplaceAll(then, "Arribute", "Attribute")
				}
			}

			app.NftmngrKeeper.SetNFTSchema(ctx, new_schema)

			// NFT Collection
			// query all nft data
			nft_list := app.NftmngrKeeper.GetAllNftData(ctx)
			for _, nft := range nft_list {
				// append to list of nft collection if schema is match
				if nft.NftSchemaCode == schema.Code {
					app.NftmngrKeeper.AddMetadataToCollection(ctx, &nft)
				}
			}

		}

		//* Module Bank *
		// register new denom metadata for using with evm
		// the usix denom cannot be used with evm because it is a decimal denom
		// then we need to create a new denom for evm

		// set new denom metadata
		app.BankKeeper.SetDenomMetaData(ctx, banktype.Metadata{
			Description: "The native evm token of the SIX Protocol",
			Base:        "asix",
			Display:     "asix",
			Symbol:      "asix",
			Name:        "evm six token",
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

		// * Module Protocaladmin *
		// add existing address to super.admin group
		// super admin to set token admin
		// token admin to set mintperm
		// mintperm to mint token
		token_admin, _ := app.ProtocoladminKeeper.GetGroup(ctx, "token.admin")
		super_admin, _ := app.ProtocoladminKeeper.GetGroup(ctx, "super.admin")
		asix_maxsupply := sdk.NewIntFromUint64(0)
		asix_coin := sdk.NewCoin("asix", asix_maxsupply)
		app.TokenmngrKeeper.SetToken(ctx, tokenmngrtypes.Token{
			Name:      "asix",
			Base:      "asix",
			MaxSupply: asix_coin,
			Mintee:    token_admin.Owner,
			Creator:   super_admin.Owner,
		})
		// add mintperm
		app.TokenmngrKeeper.SetMintperm(ctx, tokenmngrtypes.Mintperm{
			Token:   "asix",
			Address: token_admin.Owner,
			Creator: super_admin.Owner,
		})

		// * Module NFT ORACLE *
		// set nft duration
		var oracle_params nftoraclemoduletypes.Params
		oracle_params.MintRequestActiveDuration = 120 * time.Second
		oracle_params.ActionRequestActiveDuration = 120 * time.Second
		oracle_params.VerifyRequestActiveDuration = 120 * time.Second
		oracle_params.ActionSignerActiveDuration = 30 * (24 * time.Hour)
		oracle_params.SyncActionSignerActiveDuration = 300 * time.Second // five minutes
		app.NftoracleKeeper.SetParams(ctx, oracle_params)

		// migrate action Signer
		action_signersV8 := app.NftoracleKeeper.GetAllActionSigner(ctx)
		for _, action_signer := range action_signersV8 {
			action_signer.Creator = action_signer.OwnerAddress
			action_signer.CreationFlow = nftoraclemoduletypes.CreationFlow_INTERNAL_OWNER
			app.NftoracleKeeper.SetActionSigner(ctx, action_signer)
		}

		CHAINID = ctx.ChainID()
		return app.mm.RunMigrations(ctx, app.configurator, vm)
	})
}


func (app *App) VersionTrigger() {
	// sixnetChainID := big.NewInt(etherminttypes.CHAINID_NUMBER_MAINNET)
	// fivnetChainID := big.NewInt(etherminttypes.CHAINID_NUMBER_TESTNET)
	fmt.Println("########################## CHAINID ##########################", CHAINID)
	if CHAINID == "fivenet" {
		fmt.Println("########################## FIVENET ##########################")
		upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
		if err != nil {
			panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
		}

		if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
			storeUpgrades := store.StoreUpgrades{}
			// configure store loader that checks if version == upgradeHeight and applies store upgrades
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
		}
	} else {
		fmt.Println("########################## SIXNET ##########################")
		upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
		if err != nil {
			panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
		}
		if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
			storeUpgrades := store.StoreUpgrades{
				Added:   []string{evmtypes.ModuleName, feemarkettypes.ModuleName, erc20types.ModuleName},
				Deleted: []string{"wasm"},
			}
			// configure store loader that checks if version == upgradeHeight and applies store upgrades
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
		}
	}
}