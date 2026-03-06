package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	// v2types "github.com/thesixnetwork/six-protocol/v4/x/nftmngr/migrations/v2/types"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetNFTSchema set a specific nFTSchema in the store from its index
func (k Keeper) SetNFTSchema(ctx context.Context, nFTSchema types.NFTSchema) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTSchemaKeyPrefix))
	b := k.cdc.MustMarshal(&nFTSchema)
	store.Set(types.NFTSchemaKey(
		nFTSchema.Code,
	), b)
}

// GetNFTSchema returns a nFTSchema from its index
func (k Keeper) GetNFTSchema(ctx context.Context, code string) (val types.NFTSchema, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTSchemaKeyPrefix))

	b := store.Get(types.NFTSchemaKey(
		code,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNFTSchema removes a nFTSchema from the store
func (k Keeper) RemoveNFTSchema(ctx context.Context, code string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTSchemaKeyPrefix))
	store.Delete(types.NFTSchemaKey(
		code,
	))
}

// GetAllNFTSchema returns all nFTSchema
func (k Keeper) GetAllNFTSchema(ctx context.Context) (list []types.NFTSchema) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTSchemaKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NFTSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetNFTSchemaByContract set a specific nFTSchemaByContract in the store from its index
func (k Keeper) SetNFTSchemaByContract(ctx context.Context, nFTSchemaByContract types.NFTSchemaByContract) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTSchemaByContractKeyPrefix))
	b := k.cdc.MustMarshal(&nFTSchemaByContract)
	store.Set(types.NFTSchemaByContractKey(
		nFTSchemaByContract.OriginContractAddress,
	), b)
}

// GetNFTSchemaByContract returns a nFTSchemaByContract from its index
func (k Keeper) GetNFTSchemaByContract(
	ctx context.Context,
	originContractAddress string,
) (val types.NFTSchemaByContract, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTSchemaByContractKeyPrefix))

	b := store.Get(types.NFTSchemaByContractKey(
		originContractAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNFTSchemaByContract removes a nFTSchemaByContract from the store
func (k Keeper) RemoveNFTSchemaByContract(
	ctx context.Context,
	originContractAddress string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTSchemaByContractKeyPrefix))
	store.Delete(types.NFTSchemaByContractKey(
		originContractAddress,
	))
}

// GetAllNFTSchemaByContract returns all nFTSchemaByContract
func (k Keeper) GetAllNFTSchemaByContract(ctx context.Context) (list []types.NFTSchemaByContract) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTSchemaByContractKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NFTSchemaByContract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
