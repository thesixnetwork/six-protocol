package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/consume/types"
)

// SetNftUsed set a specific nftUsed in the store from its index
func (k Keeper) SetNftUsed(ctx sdk.Context, nftUsed types.NftUsed) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftUsedKeyPrefix))
	b := k.cdc.MustMarshal(&nftUsed)
	store.Set(types.NftUsedKey(
		nftUsed.Token,
	), b)
}

// GetNftUsed returns a nftUsed from its index
func (k Keeper) GetNftUsed(
	ctx sdk.Context,
	token string,
	creator string,

) (val types.NftUsed, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftUsedKeyPrefix))

	b := store.Get(types.NftUsedKey(
		token,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNftUsed removes a nftUsed from the store
func (k Keeper) RemoveNftUsed(
	ctx sdk.Context,
	token string,
	creator string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftUsedKeyPrefix))
	store.Delete(types.NftUsedKey(
		token,
	))
}

// GetAllNftUsed returns all nftUsed
func (k Keeper) GetAllNftUsed(ctx sdk.Context) (list []types.NftUsed) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftUsedKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NftUsed
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
