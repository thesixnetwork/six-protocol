package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetExecutorOfSchema set a specific executorOfSchema in the store from its index
func (k Keeper) SetExecutorOfSchema(ctx context.Context, executorOfSchema types.ExecutorOfSchema) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExecutorOfSchemaKeyPrefix))

	b := k.cdc.MustMarshal(&executorOfSchema)
	store.Set(types.ExecutorOfSchemaKey(
		executorOfSchema.NftSchemaCode,
	), b)
}

// GetExecutorOfSchema returns a executorOfSchema from its index
func (k Keeper) GetExecutorOfSchema(
	ctx context.Context,
	nftSchemaCode string,
) (val types.ExecutorOfSchema, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExecutorOfSchemaKeyPrefix))

	b := store.Get(types.ExecutorOfSchemaKey(
		nftSchemaCode,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveExecutorOfSchema removes a executorOfSchema from the store
func (k Keeper) RemoveExecutorOfSchema(
	ctx context.Context,
	nftSchemaCode string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExecutorOfSchemaKeyPrefix))
	store.Delete(types.ExecutorOfSchemaKey(
		nftSchemaCode,
	))
}

// GetAllExecutorOfSchema returns all executorOfSchema
func (k Keeper) GetAllExecutorOfSchema(ctx context.Context) (list []types.ExecutorOfSchema) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExecutorOfSchemaKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ExecutorOfSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
