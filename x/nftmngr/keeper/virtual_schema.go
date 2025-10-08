package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetVirtualSchema set a specific virSchema in the store from its index
func (k Keeper) SetVirtualSchema(ctx context.Context, virSchema types.VirtualSchema) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualSchemaKeyPrefix))
	b := k.cdc.MustMarshal(&virSchema)
	store.Set(types.VirtualSchemaKey(
		virSchema.VirtualNftSchemaCode,
	), b)
}

// GetVirtualSchema returns a virSchema from its index
func (k Keeper) GetVirtualSchema(
	ctx context.Context,
	schemaCode string,
) (val types.VirtualSchema, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualSchemaKeyPrefix))

	b := store.Get(types.VirtualSchemaKey(
		schemaCode,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVirtualSchema removes a virSchema from the store
func (k Keeper) RemoveVirtualSchema(
	ctx context.Context,
	schemaCode string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualSchemaKeyPrefix))
	store.Delete(types.VirtualSchemaKey(
		schemaCode,
	))
}

// GetAllVirtualSchema returns all virSchema
func (k Keeper) GetAllVirtualSchema(ctx context.Context) (list []types.VirtualSchema) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualSchemaKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VirtualSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
