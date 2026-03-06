package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetAdmin set a specific admin in the store from its index
func (k Keeper) SetAdmin(ctx context.Context, admin types.Admin) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AdminKeyPrefix))
	b := k.cdc.MustMarshal(&admin)
	store.Set(types.AdminKey(
		admin.Group,
		admin.Admin,
	), b)
}

// GetAdmin returns a admin from its index
func (k Keeper) GetAdmin(
	ctx context.Context,
	group string,
	admin string,
) (val types.Admin, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AdminKeyPrefix))

	b := store.Get(types.AdminKey(
		group,
		admin,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAdmin removes a admin from the store
func (k Keeper) RemoveAdmin(
	ctx context.Context,
	group string,
	admin string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AdminKeyPrefix))
	store.Delete(types.AdminKey(
		group,
		admin,
	))
}

// GetAllAdmin returns all admin
func (k Keeper) GetAllAdmin(ctx context.Context) (list []types.Admin) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AdminKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Admin
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) Authenticate(ctx context.Context, group string, address string) bool {
	_, foundSuperAdmin := k.GetAdmin(ctx, SUPER_ADMIN, address)
	_, foundGroupAdmin := k.GetAdmin(ctx, group, address)

	return foundGroupAdmin || foundSuperAdmin
}
