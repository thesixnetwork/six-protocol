package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetNftData set a specific nftData in the store from its index
func (k Keeper) SetNftData(ctx context.Context, nftData types.NftData) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftDataKeyPrefix))
	b := k.cdc.MustMarshal(&nftData)
	store.Set(types.NftDataKey(
		nftData.NftSchemaCode,
		nftData.TokenId,
	), b)
}

// GetNftData returns a nftData from its index
func (k Keeper) GetNftData(
	ctx context.Context,
	nftSchemaCode string,
	tokenId string,
) (val types.NftData, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftDataKeyPrefix))

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
	ctx context.Context,
	nftSchemaCode string,
	tokenId string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftDataKeyPrefix))
	store.Delete(types.NftDataKey(
		nftSchemaCode,
		tokenId,
	))
}

// GetAllNftData returns all nftData
func (k Keeper) GetAllNftData(ctx context.Context) (list []types.NftData) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftDataKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NftData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
