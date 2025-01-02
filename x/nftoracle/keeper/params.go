package keeper

import (
	"time"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MintRequestActiveDuration(ctx),
		k.ActionRequestActiveDuration(ctx),
		k.VerifyRequestActiveDuration(ctx),
		k.ActionSignerActiveDuration(ctx),
		k.SyncActionSignerActiveDuration(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) MintRequestActiveDuration(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyMintRequestActiveDuration, &res)
	return
}

func (k Keeper) ActionRequestActiveDuration(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyActionRequestActiveDuration, &res)
	return
}

func (k Keeper) VerifyRequestActiveDuration(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyVerifyRequestActiveDuration, &res)
	return
}

func (k Keeper) ActionSignerActiveDuration(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyActionSignerActiveDuration, &res)
	return
}

func (k Keeper) SyncActionSignerActiveDuration(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyActionSignerActiveDuration, &res)
	return
}
