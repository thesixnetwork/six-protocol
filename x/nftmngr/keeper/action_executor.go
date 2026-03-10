package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetActionExecutor set a specific actionExecutor in the store from its index
func (k Keeper) SetActionExecutor(ctx context.Context, actionExecutor types.ActionExecutor) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionExecutorKeyPrefix))
	b := k.cdc.MustMarshal(&actionExecutor)
	store.Set(types.ActionExecutorKey(
		actionExecutor.NftSchemaCode,
		actionExecutor.ExecutorAddress,
	), b)
}

// RemoveActionExecutor removes a actionExecutor from the store
func (k Keeper) RemoveActionExecutor(
	ctx context.Context,
	nftSchemaCode string,
	executorAddress string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionExecutorKeyPrefix))
	store.Delete(types.ActionExecutorKey(
		nftSchemaCode,
		executorAddress,
	))
}

// GetActionExecutor returns a actionExecutor from its index
func (k Keeper) GetActionExecutor(
	ctx context.Context,
	nftSchemaCode string,
	executorAddress string,
) (val types.ActionExecutor, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionExecutorKeyPrefix))

	b := store.Get(types.ActionExecutorKey(
		nftSchemaCode,
		executorAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllActionExecutor returns all actionExecutor
func (k Keeper) GetAllActionExecutor(ctx context.Context) (list []types.ActionExecutor) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionExecutorKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionExecutor
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
