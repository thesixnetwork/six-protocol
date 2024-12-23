package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

// SetMintperm set a specific mintperm in the store from its index
func (k Keeper) SetMintperm(ctx sdk.Context, mintperm types.Mintperm) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintpermKeyPrefix))
	b := k.cdc.MustMarshal(&mintperm)
	store.Set(types.MintpermKey(
		mintperm.Token,
		mintperm.Address,
	), b)
}

// GetMintperm returns a mintperm from its index
func (k Keeper) GetMintperm(
	ctx sdk.Context,
	token string,
	address string,
) (val types.Mintperm, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintpermKeyPrefix))

	b := store.Get(types.MintpermKey(
		token,
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMintperm removes a mintperm from the store
func (k Keeper) RemoveMintperm(
	ctx sdk.Context,
	token string,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintpermKeyPrefix))
	store.Delete(types.MintpermKey(
		token,
		address,
	))
}

// GetAllMintperm returns all mintperm
func (k Keeper) GetAllMintperm(ctx sdk.Context) (list []types.Mintperm) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintpermKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mintperm
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
