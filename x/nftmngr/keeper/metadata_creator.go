package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetMetadataCreator set a specific metadataCreator in the store from its index
func (k Keeper) SetMetadataCreator(ctx sdk.Context, metadataCreator types.MetadataCreator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetadataCreatorKeyPrefix))
	b := k.cdc.MustMarshal(&metadataCreator)
	store.Set(types.MetadataCreatorKey(
		metadataCreator.NftSchemaCode,
	), b)
}

// GetMetadataCreator returns a metadataCreator from its index
func (k Keeper) GetMetadataCreator(
	ctx sdk.Context,
	nftSchemaCode string,

) (val types.MetadataCreator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetadataCreatorKeyPrefix))

	b := store.Get(types.MetadataCreatorKey(
		nftSchemaCode,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMetadataCreator removes a metadataCreator from the store
func (k Keeper) RemoveMetadataCreator(
	ctx sdk.Context,
	nftSchemaCode string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetadataCreatorKeyPrefix))
	store.Delete(types.MetadataCreatorKey(
		nftSchemaCode,
	))
}

// GetAllMetadataCreator returns all metadataCreator
func (k Keeper) GetAllMetadataCreator(ctx sdk.Context) (list []types.MetadataCreator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetadataCreatorKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MetadataCreator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
