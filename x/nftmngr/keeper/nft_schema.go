package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetNftschema set a specific nftschema in the store from its index
func (k Keeper) SetNftschema(ctx context.Context, nftschema types.NFTSchema) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftschemaKeyPrefix))
	b := k.cdc.MustMarshal(&nftschema)
	store.Set(types.NftschemaKey(
		nftschema.Code,
	), b)
}

// GetNftschema returns a nftschema from its index
func (k Keeper) GetNftschema(
	ctx context.Context,
	code string,
) (val types.NFTSchema, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftschemaKeyPrefix))

	b := store.Get(types.NftschemaKey(
		code,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNftschema removes a nftschema from the store
func (k Keeper) RemoveNftschema(
	ctx context.Context,
	code string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftschemaKeyPrefix))
	store.Delete(types.NftschemaKey(
		code,
	))
}

// GetAllNftschema returns all nftschema
func (k Keeper) GetAllNftschema(ctx context.Context) (list []types.NFTSchema) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NftschemaKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NFTSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
