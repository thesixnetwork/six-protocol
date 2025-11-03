package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetActionSignerConfig set a specific actionSignerConfig in the store from its index
func (k Keeper) SetActionSignerConfig(ctx context.Context, actionSignerConfig types.ActionSignerConfig) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionSignerConfigKeyPrefix))
	b := k.cdc.MustMarshal(&actionSignerConfig)
	store.Set(types.ActionSignerConfigKey(
		actionSignerConfig.Chain,
	), b)
}

// GetActionSignerConfig returns a actionSignerConfig from its index
func (k Keeper) GetActionSignerConfig(
	ctx context.Context,
	chain string,
) (val types.ActionSignerConfig, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionSignerConfigKeyPrefix))

	b := store.Get(types.ActionSignerConfigKey(
		chain,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActionSignerConfig removes a actionSignerConfig from the store
func (k Keeper) RemoveActionSignerConfig(
	ctx context.Context,
	chain string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionSignerConfigKeyPrefix))
	store.Delete(types.ActionSignerConfigKey(
		chain,
	))
}

// GetAllActionSignerConfig returns all actionSignerConfig
func (k Keeper) GetAllActionSignerConfig(ctx context.Context) (list []types.ActionSignerConfig) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionSignerConfigKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionSignerConfig
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
