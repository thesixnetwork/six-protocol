package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetSchemaAttribute set a specific schemaAttribute in the store from its index
func (k Keeper) SetSchemaAttribute(ctx context.Context, schemaAttribute types.SchemaAttribute) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SchemaAttributeKeyPrefix))
	b := k.cdc.MustMarshal(&schemaAttribute)
	store.Set(types.SchemaAttributeKey(
		schemaAttribute.NftSchemaCode,
		schemaAttribute.Name,
	), b)
}

// GetSchemaAttribute returns a schemaAttribute from its index
func (k Keeper) GetSchemaAttribute(
	ctx context.Context,
	nftSchemaCode string,
	name string,
) (val types.SchemaAttribute, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SchemaAttributeKeyPrefix))

	b := store.Get(types.SchemaAttributeKey(
		nftSchemaCode,
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSchemaAttribute removes a schemaAttribute from the store
func (k Keeper) RemoveSchemaAttribute(
	ctx context.Context,
	nftSchemaCode string,
	name string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SchemaAttributeKeyPrefix))
	store.Delete(types.SchemaAttributeKey(
		nftSchemaCode,
		name,
	))
}

// GetAllSchemaAttribute returns all schemaAttribute
func (k Keeper) GetAllSchemaAttribute(ctx context.Context) (list []types.SchemaAttribute) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SchemaAttributeKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SchemaAttribute
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAttributeOfSchema returns a attributeOfSchema from its index
func (k Keeper) GetAttributeOfSchema(
	ctx context.Context,
	nftSchemaCode string,
) (val types.AttributeOfSchema, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SchemaAttributeKeyPrefix))

	b := store.Get(types.AttributeOfSchemaKey(
		nftSchemaCode,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAttributeOfSchema removes a attributeOfSchema from the store
func (k Keeper) RemoveAttributeOfSchema(
	ctx context.Context,
	nftSchemaCode string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SchemaAttributeKeyPrefix))
	store.Delete(types.AttributeOfSchemaKey(
		nftSchemaCode,
	))
}

// GetAllAttributeOfSchema returns all attributeOfSchema
func (k Keeper) GetAllAttributeOfSchema(ctx context.Context) (list []types.AttributeOfSchema) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SchemaAttributeKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AttributeOfSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
