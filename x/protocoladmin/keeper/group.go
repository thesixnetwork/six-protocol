package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

const SUPER_ADMIN = "super.admin"

// SetGroup set a specific group in the store from its index
func (k Keeper) SetGroup(ctx context.Context, group types.Group) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GroupKeyPrefix))
	b := k.cdc.MustMarshal(&group)
	store.Set(types.GroupKey(
		group.Name,
	), b)
}

// GetGroup returns a group from its index
func (k Keeper) GetGroup(
	ctx context.Context,
	name string,
) (val types.Group, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GroupKeyPrefix))

	b := store.Get(types.GroupKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGroup removes a group from the store
func (k Keeper) RemoveGroup(
	ctx context.Context,
	name string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GroupKeyPrefix))
	store.Delete(types.GroupKey(
		name,
	))
}

// GetAllGroup returns all group
func (k Keeper) GetAllGroup(ctx context.Context) (list []types.Group) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GroupKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Group
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
