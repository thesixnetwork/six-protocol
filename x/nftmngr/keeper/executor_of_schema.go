package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetExecutorOfSchema set a specific executorOfSchema in the store from its index
func (k Keeper) SetExecutorOfSchema(ctx sdk.Context, executorOfSchema types.ExecutorOfSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutorOfSchemaKeyPrefix))
	b := k.cdc.MustMarshal(&executorOfSchema)
	store.Set(types.ExecutorOfSchemaKey(
		executorOfSchema.NftSchemaCode,
	), b)
}

// GetExecutorOfSchema returns a executorOfSchema from its index
func (k Keeper) GetExecutorOfSchema(
	ctx sdk.Context,
	nftSchemaCode string,
) (val types.ExecutorOfSchema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutorOfSchemaKeyPrefix))

	b := store.Get(types.ExecutorOfSchemaKey(
		nftSchemaCode,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveExecutorOfSchema removes a executorOfSchema from the store
func (k Keeper) RemoveExecutorOfSchema(
	ctx sdk.Context,
	nftSchemaCode string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutorOfSchemaKeyPrefix))
	store.Delete(types.ExecutorOfSchemaKey(
		nftSchemaCode,
	))
}

// GetAllExecutorOfSchema returns all executorOfSchema
func (k Keeper) GetAllExecutorOfSchema(ctx sdk.Context) (list []types.ExecutorOfSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutorOfSchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ExecutorOfSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
