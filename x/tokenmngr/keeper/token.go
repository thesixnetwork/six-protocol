package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetToken set a specific token in the store from its index
func (k Keeper) SetToken(ctx context.Context, token types.Token) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TokenKeyPrefix))
	b := k.cdc.MustMarshal(&token)
	store.Set(types.TokenKey(
		token.Name,
	), b)
}

// GetToken returns a token from its index
func (k Keeper) GetToken(
	ctx context.Context,
	name string,
) (val types.Token, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TokenKeyPrefix))

	b := store.Get(types.TokenKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveToken removes a token from the store
func (k Keeper) RemoveToken(
	ctx context.Context,
	name string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TokenKeyPrefix))
	store.Delete(types.TokenKey(
		name,
	))
}

// GetAllToken returns all token
func (k Keeper) GetAllToken(ctx context.Context) (list []types.Token) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TokenKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Token
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
