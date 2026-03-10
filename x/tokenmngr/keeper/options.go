package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetOptions set options in the store
func (k Keeper) SetOptions(ctx context.Context, options types.Options) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OptionsKey))
	b := k.cdc.MustMarshal(&options)
	store.Set([]byte{0}, b)
}

// GetOptions returns options
func (k Keeper) GetOptions(ctx context.Context) (val types.Options, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OptionsKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOptions removes options from the store
func (k Keeper) RemoveOptions(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OptionsKey))
	store.Delete([]byte{0})
}
