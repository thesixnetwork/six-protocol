package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetVirtualSchema set a specific virSchema in the store from its index
func (k Keeper) SetVirtualSchema(ctx sdk.Context, virSchema types.VirtualSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaKeyPrefix))
	b := k.cdc.MustMarshal(&virSchema)
	store.Set(types.VirtualSchemaKey(
		virSchema.VirtualNftSchemaCode,
	), b)
}

// GetVirtualSchema returns a virSchema from its index
func (k Keeper) GetVirtualSchema(
	ctx sdk.Context,
	index string,

) (val types.VirtualSchema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaKeyPrefix))

	b := store.Get(types.VirtualSchemaKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVirtualSchema removes a virSchema from the store
func (k Keeper) RemoveVirtualSchema(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaKeyPrefix))
	store.Delete(types.VirtualSchemaKey(
		index,
	))
}

// GetAllVirtualSchema returns all virSchema
func (k Keeper) GetAllVirtualSchema(ctx sdk.Context) (list []types.VirtualSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VirtualSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
