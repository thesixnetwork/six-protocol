package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

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
