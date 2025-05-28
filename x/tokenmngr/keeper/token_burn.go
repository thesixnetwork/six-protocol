package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetTokenBurn set a specific tokenBurn in the store from its index
func (k Keeper) SetTokenBurn(ctx context.Context, tokenBurn types.TokenBurn) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TokenBurnKeyPrefix))
	b := k.cdc.MustMarshal(&tokenBurn)
	store.Set(types.TokenBurnKey(
		tokenBurn.Amount.Denom,
	), b)
}

// GetTokenBurn returns a tokenBurn from its index
func (k Keeper) GetTokenBurn(
	ctx context.Context,
	index string,
) (val types.TokenBurn, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TokenBurnKeyPrefix))

	b := store.Get(types.TokenBurnKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTokenBurn removes a tokenBurn from the store
func (k Keeper) RemoveTokenBurn(
	ctx context.Context,
	index string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TokenBurnKeyPrefix))
	store.Delete(types.TokenBurnKey(
		index,
	))
}

// GetAllTokenBurn returns all tokenBurn
func (k Keeper) GetAllTokenBurn(ctx context.Context) (list []types.TokenBurn) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TokenBurnKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TokenBurn
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
