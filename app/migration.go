package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	nftmngrtypes "github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (app *App) SIXNFTMigrationFromV1ToV2Handlers(ctx sdk.Context) {

	// get all NFTSchema
	nftSchemasV1 := app.NftmngrKeeper.GetAllNFTSchemaV1(ctx)

	for _, nftSchemaV1 := range nftSchemasV1 {
		// migrate list of system actioners to new map
		for _, systemActioner := range nftSchemaV1.SystemActioners {
			// set system actioner to new map
			app.NftmngrKeeper.SetActionExecutor(ctx, nftmngrtypes.ActionExecutor{
				NftSchemaCode:   nftSchemaV1.Code,
				ExecutorAddress: systemActioner,
				Creator: nftSchemaV1.Owner,
			})
		}

		// migrate schema to new schema
		app.NftmngrKeeper.SetNFTSchema(ctx, nftmngrtypes.NFTSchema{
			Code: 		  nftSchemaV1.Code,
			Name: 		  nftSchemaV1.Name,
			Owner: 		  nftSchemaV1.Owner,
			OriginData:  nftSchemaV1.OriginData,
			OnchainData: nftSchemaV1.OnchainData,
			IsVerified: nftSchemaV1.IsVerified,
			MintAuthorization: nftSchemaV1.MintAuthorization,
		})
	}

}
