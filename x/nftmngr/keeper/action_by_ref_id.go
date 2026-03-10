package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetActionByRefId set a specific actionByRefId in the store from its index
func (k Keeper) SetActionByRefId(ctx context.Context, actionByRefId types.ActionByRefId) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionByRefIdKeyPrefix))
	b := k.cdc.MustMarshal(&actionByRefId)
	store.Set(types.ActionByRefIdKey(
		actionByRefId.RefId,
	), b)
}

// GetActionByRefId returns a actionByRefId from its index
func (k Keeper) GetActionByRefId(
	ctx context.Context,
	refId string,
) (val types.ActionByRefId, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionByRefIdKeyPrefix))

	b := store.Get(types.ActionByRefIdKey(
		refId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActionByRefId removes a actionByRefId from the store
func (k Keeper) RemoveActionByRefId(
	ctx context.Context,
	refId string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionByRefIdKeyPrefix))
	store.Delete(types.ActionByRefIdKey(
		refId,
	))
}

// GetAllActionByRefId returns all actionByRefId
func (k Keeper) GetAllActionByRefId(ctx context.Context) (list []types.ActionByRefId) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionByRefIdKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionByRefId
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
