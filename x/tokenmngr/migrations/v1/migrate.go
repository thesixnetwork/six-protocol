// Package v1
package v1

import (
	"cosmossdk.io/core/store"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func MigrateStore(ctx sdk.Context, storeService store.KVStoreService) error {
	store := runtime.KVStoreAdapter(storeService.OpenKVStore(ctx))
	burnStore := prefix.NewStore(store, types.KeyPrefix(types.BurnKey))
	return removeBurns(burnStore)
}

// removeBurns use to remove buns history from KV
// then use indexer to collect burn history instead
func removeBurns(store storetypes.KVStore) error {
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		store.Delete(iterator.Key())
	}

	// reset burn count
	byteKey := []byte(types.BurnCountKey)
	store.Delete(byteKey)
	return nil
}
