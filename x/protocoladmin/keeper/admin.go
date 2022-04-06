package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

// SetAdmin set a specific admin in the store from its index
func (k Keeper) SetAdmin(ctx sdk.Context, admin types.Admin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdminKeyPrefix))
	b := k.cdc.MustMarshal(&admin)
	store.Set(types.AdminKey(
		admin.Group,
		admin.Admin,
	), b)
}

// GetAdmin returns a admin from its index
func (k Keeper) GetAdmin(
	ctx sdk.Context,
	group string,
	admin string,

) (val types.Admin, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdminKeyPrefix))

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
	ctx sdk.Context,
	group string,
	admin string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdminKeyPrefix))
	store.Delete(types.AdminKey(
		group,
		admin,
	))
}

// GetAllAdmin returns all admin
func (k Keeper) GetAllAdmin(ctx sdk.Context) (list []types.Admin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdminKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Admin
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) Authenticate(ctx sdk.Context, group string, address string) bool {
	_, foundSuperAdmin := k.GetAdmin(ctx, SUPER_ADMIN, address)
	_, foundGroupAdmin := k.GetAdmin(ctx, group, address)

	return foundGroupAdmin || foundSuperAdmin
}
