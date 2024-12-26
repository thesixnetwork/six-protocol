package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetMetadataCreator set a specific metadataCreator in the store from its index
func (k Keeper) SetMetadataCreator(ctx sdk.Context, metadataCreator types.MetadataCreator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetadataCreatorKeyPrefix))
	b := k.cdc.MustMarshal(&metadataCreator)
	store.Set(types.MetadataCreatorKey(
		metadataCreator.NftSchemaCode,
	), b)
}

// GetMetadataCreator returns a metadataCreator from its index
func (k Keeper) GetMetadataCreator(
	ctx sdk.Context,
	nftSchemaCode string,
) (val types.MetadataCreator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetadataCreatorKeyPrefix))

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
	ctx sdk.Context,
	nftSchemaCode string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetadataCreatorKeyPrefix))
	store.Delete(types.MetadataCreatorKey(
		nftSchemaCode,
	))
}

// GetAllMetadataCreator returns all metadataCreator
func (k Keeper) GetAllMetadataCreator(ctx sdk.Context) (list []types.MetadataCreator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetadataCreatorKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MetadataCreator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetNftData set a specific nftData in the store from its index
func (k Keeper) SetNftData(ctx sdk.Context, nftData types.NftData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftDataKeyPrefix))
	b := k.cdc.MustMarshal(&nftData)
	store.Set(types.NftDataKey(
		nftData.NftSchemaCode,
		nftData.TokenId,
	), b)
}

// GetNftData returns a nftData from its index
func (k Keeper) GetNftData(
	ctx sdk.Context,
	nftSchemaCode string,
	tokenId string,
) (val types.NftData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftDataKeyPrefix))

	b := store.Get(types.NftDataKey(
		nftSchemaCode,
		tokenId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNftData removes a nftData from the store
func (k Keeper) RemoveNftData(
	ctx sdk.Context,
	nftSchemaCode string,
	tokenId string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftDataKeyPrefix))
	store.Delete(types.NftDataKey(
		nftSchemaCode,
		tokenId,
	))
}

// GetAllNftData returns all nftData
func (k Keeper) GetAllNftData(ctx sdk.Context) (list []types.NftData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftDataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NftData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) AddMetadataToCollection(ctx sdk.Context, data *types.NftData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CollectionkeyPrefix(data.NftSchemaCode)) // returns KVStore with prefix of SchemaCode + "/"
	b := k.cdc.MustMarshal(data)
	store.Set([]byte(data.TokenId), b) // set the value of data to store of SchemaCode + "/" + TokenId as key (value is marshalled data) it will be => SchemaCode/TokenId: data
}

// SetNftCollection set a specific nftCollection in the store from its index
func (k Keeper) SetNftCollection(ctx sdk.Context, nftCollection types.NftCollection) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftCollectionKeyPrefix))
	b := k.cdc.MustMarshal(&nftCollection)
	store.Set(types.NftCollectionKey(
		nftCollection.NftSchemaCode,
	), b)
}

// SetNftCollectionDataCount set the total number of nftCollection
func (k Keeper) SetNftCollectionDataCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftCollectionDataCountKey))
	byteKey := []byte(types.NftCollectionDataCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetNftCollectionDataCount get the total number of nftCollection
func (k Keeper) GetNftCollectionDataCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftCollectionDataCountKey))
	byteKey := []byte(types.NftCollectionDataCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// GetNftCollection returns a nftCollection from its index
func (k Keeper) GetNftCollection(
	ctx sdk.Context,
	nftSchemaCode string,
) (val types.NftCollection, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftCollectionKeyPrefix))

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
	ctx sdk.Context,
	nftSchemaCode string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftCollectionKeyPrefix))
	store.Delete(types.NftCollectionKey(
		nftSchemaCode,
	))
}

// GetAllNftCollection returns all nftCollection
func (k Keeper) GetAllNftCollection(ctx sdk.Context) (list []types.NftCollection) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftCollectionKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NftCollection
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
