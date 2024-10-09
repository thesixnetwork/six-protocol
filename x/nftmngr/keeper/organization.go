package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetOrganization set a specific organization in the store from its index
func (k Keeper) SetOrganization(ctx sdk.Context, organization types.Organization) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrganizationKeyPrefix))
	b := k.cdc.MustMarshal(&organization)
	store.Set(types.OrganizationKey(
		organization.Name,
	), b)
}

// GetOrganization returns a organization from its index
func (k Keeper) GetOrganization(
	ctx sdk.Context,
	name string,

) (val types.Organization, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrganizationKeyPrefix))

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
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrganizationKeyPrefix))
	store.Delete(types.OrganizationKey(
		name,
	))
}

// GetAllOrganization returns all organization
func (k Keeper) GetAllOrganization(ctx sdk.Context) (list []types.Organization) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrganizationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Organization
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
