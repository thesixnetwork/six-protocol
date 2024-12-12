package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetDisableVirtualSchema set a specific disableVirtualSchema in the store from its index
func (k Keeper) SetDisableVirtualSchema(ctx sdk.Context, disableVirtualSchema types.DisableVirtualSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisableVirtualSchemaKeyPrefix))
	b := k.cdc.MustMarshal(&disableVirtualSchema)
	store.Set(types.DisableVirtualSchemaKey(
		disableVirtualSchema.Id,
	), b)
}

// GetDisableVirtualSchema returns a disableVirtualSchema from its index
func (k Keeper) GetDisableVirtualSchema(
	ctx sdk.Context,
	id string,

) (val types.DisableVirtualSchema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisableVirtualSchemaKeyPrefix))

	b := store.Get(types.DisableVirtualSchemaKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDisableVirtualSchema removes a disableVirtualSchema from the store
func (k Keeper) RemoveDisableVirtualSchema(
	ctx sdk.Context,
	id string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisableVirtualSchemaKeyPrefix))
	store.Delete(types.DisableVirtualSchemaKey(
		id,
	))
}

// GetAllDisableVirtualSchema returns all disableVirtualSchema
func (k Keeper) GetAllDisableVirtualSchema(ctx sdk.Context) (list []types.DisableVirtualSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisableVirtualSchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DisableVirtualSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
