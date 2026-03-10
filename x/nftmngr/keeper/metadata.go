package keeper

import (
	"context"
	"encoding/binary"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetMetadataCreator set a specific metadataCreator in the store from its index
func (k Keeper) SetMetadataCreator(ctx context.Context, metadataCreator types.MetadataCreator) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MetadataCreatorKeyPrefix))
	b := k.cdc.MustMarshal(&metadataCreator)
	store.Set(types.MetadataCreatorKey(
		metadataCreator.NftSchemaCode,
	), b)
}

// GetMetadataCreator returns a metadataCreator from its index
func (k Keeper) GetMetadataCreator(
	ctx context.Context,
	nftSchemaCode string,
) (val types.MetadataCreator, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MetadataCreatorKeyPrefix))

	b := store.Get(types.MetadataCreatorKey(
		nftSchemaCode,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMetadataCreator removes a metadataCreator from the store
func (k Keeper) RemoveMetadataCreator(
	ctx context.Context,
	nftSchemaCode string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MetadataCreatorKeyPrefix))
	store.Delete(types.MetadataCreatorKey(
		nftSchemaCode,
	))
}

// GetAllMetadataCreator returns all metadataCreator
func (k Keeper) GetAllMetadataCreator(ctx context.Context) (list []types.MetadataCreator) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MetadataCreatorKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MetadataCreator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) AddMetadataToCollection(ctx context.Context, data *types.NftData) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.CollectionkeyPrefix(data.NftSchemaCode))
	b := k.cdc.MustMarshal(data)
	store.Set([]byte(data.TokenId), b) // set the value of data to store of SchemaCode + "/" + TokenId as key (value is marshalled data) it will be => SchemaCode/TokenId: data
}

// SetNftCollection set a specific nftCollection in the store from its index
func (k Keeper) SetNftCollection(ctx context.Context, nftCollection types.NftCollection) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftCollectionKeyPrefix))
	b := k.cdc.MustMarshal(&nftCollection)
	store.Set(types.NftCollectionKey(
		nftCollection.NftSchemaCode,
	), b)
}

// SetNftCollectionDataCount set the total number of nftCollection
func (k Keeper) SetNftCollectionDataCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftCollectionDataCountKey))
	byteKey := []byte(types.NftCollectionDataCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetNftCollectionDataCount get the total number of nftCollection
func (k Keeper) GetNftCollectionDataCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftCollectionDataCountKey))
	byteKey := []byte(types.NftCollectionDataCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// GetNftCollection returns a nftCollection from its index
func (k Keeper) GetNftCollection(
	ctx context.Context,
	nftSchemaCode string,
) (val types.NftCollection, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftCollectionKeyPrefix))

	b := store.Get(types.NftCollectionKey(
		nftSchemaCode,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNftCollection removes a nftCollection from the store
func (k Keeper) RemoveNftCollection(
	ctx context.Context,
	nftSchemaCode string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftCollectionKeyPrefix))
	store.Delete(types.NftCollectionKey(
		nftSchemaCode,
	))
}

// GetAllNftCollection returns all nftCollection
func (k Keeper) GetAllNftCollection(ctx context.Context) (list []types.NftCollection) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftCollectionKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NftCollection
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
