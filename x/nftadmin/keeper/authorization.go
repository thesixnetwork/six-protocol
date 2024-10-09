package keeper

import (
	"github.com/thesixnetwork/six-protocol/x/nftadmin/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAuthorization set authorization in the store
func (k Keeper) SetAuthorization(ctx sdk.Context, authorization types.Authorization) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorizationKey))
	b := k.cdc.MustMarshal(&authorization)
	store.Set([]byte{0}, b)
}

// GetAuthorization returns authorization
func (k Keeper) GetAuthorization(ctx sdk.Context) (val types.Authorization, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorizationKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAuthorization removes authorization from the store
func (k Keeper) RemoveAuthorization(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorizationKey))
	store.Delete([]byte{0})
}

// HasPermission returns true if the address has permission on a given name
func (k Keeper) HasPermission(ctx sdk.Context, name string, addr sdk.AccAddress) bool {
	auth, found := k.GetAuthorization(ctx)
	if !found {
		return false
	}

	if auth.Permissions == nil {
		return false
	}

	addressList := auth.Permissions.GetPermissionAddressByKey(name)
	if addressList == nil {
		return false
	}

	mapAll := make(map[string]string)
	for _, addr := range addressList.Addresses {
		mapAll[addr] = addr
	}

	if _, found := mapAll[addr.String()]; !found {
		return false
	}

	return true
}
