// Package v1
package v1

import (
	"cosmossdk.io/core/store"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftadmin/types"
	"github.com/thesixnetwork/six-protocol/x/nftadmin/types/v1"
)

func MigrateStore(ctx sdk.Context, storeService store.KVStoreService) error {
	store := runtime.KVStoreAdapter(storeService.OpenKVStore(ctx))
	burnStore := prefix.NewStore(store, types.KeyPrefix(types.AuthorizationKey))
	return restructAutorization(burnStore)
}

// restruct authorization
func restructAutorization(store storetypes.KVStore) error {
	var v1Auth = v1.Authorization
	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}
}
