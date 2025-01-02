package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetActionOfSchema set a specific actionOfSchema in the store from its index
func (k Keeper) SetActionOfSchema(ctx sdk.Context, actionOfSchema types.ActionOfSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionOfSchemaKeyPrefix))
	b := k.cdc.MustMarshal(&actionOfSchema)
	store.Set(types.ActionOfSchemaKey(
		actionOfSchema.NftSchemaCode,
		actionOfSchema.Name,
	), b)
}

// GetActionOfSchema returns a actionOfSchema from its index
func (k Keeper) GetActionOfSchema(
	ctx sdk.Context,
	nftSchemaCode string,
	name string,

) (val types.ActionOfSchema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionOfSchemaKeyPrefix))

	b := store.Get(types.ActionOfSchemaKey(
		nftSchemaCode,
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActionOfSchema removes a actionOfSchema from the store
func (k Keeper) RemoveActionOfSchema(
	ctx sdk.Context,
	nftSchemaCode string,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionOfSchemaKeyPrefix))
	store.Delete(types.ActionOfSchemaKey(
		nftSchemaCode,
		name,
	))
}

// GetAllActionOfSchema returns all actionOfSchema
func (k Keeper) GetAllActionOfSchema(ctx sdk.Context) (list []types.ActionOfSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionOfSchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionOfSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
