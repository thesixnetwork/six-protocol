package keeper

import (
	"context"
	"time"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx context.Context) (params types.Params) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return params
	}

	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx context.Context, params types.Params) error {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	store.Set(types.ParamsKey, bz)

	return nil
}

func (k Keeper) MintRequestActiveDuration(ctx context.Context) (res time.Duration) {
	return k.GetParams(ctx).MintRequestActiveDuration
}

func (k Keeper) ActionRequestActiveDuration(ctx context.Context) (res time.Duration) {
	return k.GetParams(ctx).ActionRequestActiveDuration
}

func (k Keeper) VerifyRequestActiveDuration(ctx context.Context) (res time.Duration) {
	return k.GetParams(ctx).VerifyRequestActiveDuration
}

func (k Keeper) ActionSignerActiveDuration(ctx context.Context) (res time.Duration) {
	return k.GetParams(ctx).ActionSignerActiveDuration
}

func (k Keeper) SyncActionSignerActiveDuration(ctx context.Context) (res time.Duration) {
	return k.GetParams(ctx).SyncActionSignerActiveDuration
}
