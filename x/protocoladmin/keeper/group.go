package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

const SUPER_ADMIN = "super.admin"

// SetGroup set a specific group in the store from its index
func (k Keeper) SetGroup(ctx sdk.Context, group types.Group) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKeyPrefix))
	b := k.cdc.MustMarshal(&group)
	store.Set(types.GroupKey(
		group.Name,
	), b)
}

// GetGroup returns a group from its index
func (k Keeper) GetGroup(
	ctx sdk.Context,
	name string,
) (val types.Group, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKeyPrefix))

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
	ctx sdk.Context,
	name string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKeyPrefix))
	store.Delete(types.GroupKey(
		name,
	))
}

// GetAllGroup returns all group
func (k Keeper) GetAllGroup(ctx sdk.Context) (list []types.Group) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Group
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
