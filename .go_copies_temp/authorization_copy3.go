package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/v4/x/nftadmin/types"

	"cosmossdk.io/store/prefix"

	// storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAuthorization set authorization in the store
func (k Keeper) SetAuthorization(ctx context.Context, authorization types.Authorization) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AuthorizationKey))
	b := k.cdc.MustMarshal(&authorization)
	store.Set([]byte{0}, b)
}

// GetAuthorization returns authorization
func (k Keeper) GetAuthorization(ctx context.Context) (val types.Authorization, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AuthorizationKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAuthorization removes authorization from the store
func (k Keeper) RemoveAuthorization(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AuthorizationKey))
	store.Delete([]byte{0})
}

// HasPermission returns true if the address has permission on a given name
func (k Keeper) HasPermission(ctx context.Context, name string, addr sdk.AccAddress) bool {
	auth, found := k.GetAuthorization(ctx)
	if !found {
		return false
	}

	if auth.Permissions == nil {
		return false
	}

	addressList := auth.GetPermissionAddressByKey(name)
	if addressList == nil {
		return false
	}

	mapAll := make(map[string]string)
	for _, addr := range addressList {
		mapAll[addr] = addr
	}

	if _, found := mapAll[addr.String()]; !found {
		return false
	}

	return true
}
