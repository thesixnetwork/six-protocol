package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetActionOfSchema set a specific actionOfSchema in the store from its index
func (k Keeper) SetActionOfSchema(ctx context.Context, actionOfSchema types.ActionOfSchema) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionOfSchemaKeyPrefix))
	b := k.cdc.MustMarshal(&actionOfSchema)
	store.Set(types.ActionOfSchemaKey(
		actionOfSchema.NftSchemaCode,
		actionOfSchema.Name,
	), b)
}

// GetActionOfSchema returns a actionOfSchema from its index
func (k Keeper) GetActionOfSchema(
	ctx context.Context,
	nftSchemaCode string,
	name string,
) (val types.ActionOfSchema, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionOfSchemaKeyPrefix))

	b := store.Get(types.ActionOfSchemaKey(
		nftSchemaCode,
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActionOfSchema removes a actionOfSchema from the store
func (k Keeper) RemoveActionOfSchema(
	ctx context.Context,
	nftSchemaCode string,
	name string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionOfSchemaKeyPrefix))
	store.Delete(types.ActionOfSchemaKey(
		nftSchemaCode,
		name,
	))
}

// GetAllActionOfSchema returns all actionOfSchema
func (k Keeper) GetAllActionOfSchema(ctx context.Context) (list []types.ActionOfSchema) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionOfSchemaKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionOfSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
