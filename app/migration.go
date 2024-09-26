package app

import (
	"math/big"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (app *App) MigrateParam(ctx sdk.Context) {
	currentParam := app.EVMKeeper.GetParams(ctx)
	migrateParam := currentParam
	migrateParam.AllowUnprotectedTxs = true
	app.EVMKeeper.SetParams(ctx, migrateParam)

}
