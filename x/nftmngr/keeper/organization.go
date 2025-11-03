package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetOrganization set a specific organization in the store from its index
func (k Keeper) SetOrganization(ctx context.Context, organization types.Organization) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OrganizationKeyPrefix))
	b := k.cdc.MustMarshal(&organization)
	store.Set(types.OrganizationKey(
		organization.Name,
	), b)
}

// GetOrganization returns a organization from its index
func (k Keeper) GetOrganization(
	ctx context.Context,
	name string,
) (val types.Organization, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OrganizationKeyPrefix))

	b := store.Get(types.OrganizationKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrganization removes a organization from the store
func (k Keeper) RemoveOrganization(
	ctx context.Context,
	name string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OrganizationKeyPrefix))
	store.Delete(types.OrganizationKey(
		name,
	))
}

// GetAllOrganization returns all organization
func (k Keeper) GetAllOrganization(ctx context.Context) (list []types.Organization) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OrganizationKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Organization
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
