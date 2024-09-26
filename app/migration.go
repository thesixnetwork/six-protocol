package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/precompiles"
)

func (app *App) MigrateParam(ctx sdk.Context) {
	currentParam := app.EVMKeeper.GetParams(ctx)
	migrateParam := currentParam
	migrateParam.AllowUnprotectedTxs = true
	app.EVMKeeper.SetParams(ctx, migrateParam)

	if err := precompiles.InitializePrecompiles(
		false,
		app.appCodec,
		app.BankKeeper,
		app.AccountKeeper,
		app.NftmngrKeeper,
	); err != nil {
		panic(err)
	}

}
