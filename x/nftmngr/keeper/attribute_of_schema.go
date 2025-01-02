package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// GetAttributeOfSchema returns a attributeOfSchema from its index
func (k Keeper) GetAttributeOfSchema(
	ctx sdk.Context,
	nftSchemaCode string,

) (val types.AttributeOfSchema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AttributeOfSchemaKeyPrefix))

	b := store.Get(types.AttributeOfSchemaKey(
		nftSchemaCode,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAttributeOfSchema removes a attributeOfSchema from the store
func (k Keeper) RemoveAttributeOfSchema(
	ctx sdk.Context,
	nftSchemaCode string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AttributeOfSchemaKeyPrefix))
	store.Delete(types.AttributeOfSchemaKey(
		nftSchemaCode,
	))
}

// GetAllAttributeOfSchema returns all attributeOfSchema
func (k Keeper) GetAllAttributeOfSchema(ctx sdk.Context) (list []types.AttributeOfSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AttributeOfSchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AttributeOfSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
