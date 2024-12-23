package keeper

import (
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNftData set a specific nftData in the store from its index
func (k Keeper) SetNftData(ctx sdk.Context, nftData types.NftData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftDataKeyPrefix))
	b := k.cdc.MustMarshal(&nftData)
	store.Set(types.NftDataKey(
		nftData.NftSchemaCode,
		nftData.TokenId,
	), b)
}

// GetNftData returns a nftData from its index
func (k Keeper) GetNftData(
	ctx sdk.Context,
	nftSchemaCode string,
	tokenId string,
) (val types.NftData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftDataKeyPrefix))

	b := store.Get(types.NftDataKey(
		nftSchemaCode,
		tokenId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNftData removes a nftData from the store
func (k Keeper) RemoveNftData(
	ctx sdk.Context,
	nftSchemaCode string,
	tokenId string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftDataKeyPrefix))
	store.Delete(types.NftDataKey(
		nftSchemaCode,
		tokenId,
	))
}

// GetAllNftData returns all nftData
func (k Keeper) GetAllNftData(ctx sdk.Context) (list []types.NftData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftDataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NftData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
