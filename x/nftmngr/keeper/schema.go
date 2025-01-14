package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v2types "github.com/thesixnetwork/six-protocol/x/nftmngr/migrations/v2/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetNFTSchema set a specific nFTSchema in the store from its index
func (k Keeper) SetNFTSchema(ctx sdk.Context, nFTSchema types.NFTSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))
	b := k.cdc.MustMarshal(&nFTSchema)
	store.Set(types.NFTSchemaKey(
		nFTSchema.Code,
	), b)
}

// GetNFTSchema returns a nFTSchema from its index
func (k Keeper) GetNFTSchema(ctx sdk.Context, code string) (val types.NFTSchema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))

	b := store.Get(types.NFTSchemaKey(
		code,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNFTSchema removes a nFTSchema from the store
func (k Keeper) RemoveNFTSchema(ctx sdk.Context, code string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))
	store.Delete(types.NFTSchemaKey(
		code,
	))
}

// GetAllNFTSchema returns all nFTSchema
func (k Keeper) GetAllNFTSchema(ctx sdk.Context) (list []types.NFTSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NFTSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetAllNFTSchemaLegacy(ctx sdk.Context) (list []v2types.NFTSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v2types.NFTSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetNFTSchemaByContract set a specific nFTSchemaByContract in the store from its index
func (k Keeper) SetNFTSchemaByContract(ctx sdk.Context, nFTSchemaByContract types.NFTSchemaByContract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaByContractKeyPrefix))
	b := k.cdc.MustMarshal(&nFTSchemaByContract)
	store.Set(types.NFTSchemaByContractKey(
		nFTSchemaByContract.OriginContractAddress,
	), b)
}

// GetNFTSchemaByContract returns a nFTSchemaByContract from its index
func (k Keeper) GetNFTSchemaByContract(
	ctx sdk.Context,
	originContractAddress string,
) (val types.NFTSchemaByContract, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaByContractKeyPrefix))

	b := store.Get(types.NFTSchemaByContractKey(
		originContractAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNFTSchemaByContract removes a nFTSchemaByContract from the store
func (k Keeper) RemoveNFTSchemaByContract(
	ctx sdk.Context,
	originContractAddress string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaByContractKeyPrefix))
	store.Delete(types.NFTSchemaByContractKey(
		originContractAddress,
	))
}

// GetAllNFTSchemaByContract returns all nFTSchemaByContract
func (k Keeper) GetAllNFTSchemaByContract(ctx sdk.Context) (list []types.NFTSchemaByContract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTSchemaByContractKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NFTSchemaByContract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
