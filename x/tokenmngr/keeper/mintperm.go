package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetMintperm set a specific mintperm in the store from its index
func (k Keeper) SetMintperm(ctx context.Context, mintperm types.Mintperm) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintpermKeyPrefix))
	b := k.cdc.MustMarshal(&mintperm)
	store.Set(types.MintpermKey(
		mintperm.Token,
		mintperm.Address,
	), b)
}

// GetMintperm returns a mintperm from its index
func (k Keeper) GetMintperm(
	ctx context.Context,
	token string,
	address string,
) (val types.Mintperm, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintpermKeyPrefix))

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
	ctx context.Context,
	token string,
	address string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintpermKeyPrefix))
	store.Delete(types.MintpermKey(
		token,
		address,
	))
}

// GetAllMintperm returns all mintperm
func (k Keeper) GetAllMintperm(ctx context.Context) (list []types.Mintperm) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintpermKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mintperm
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
