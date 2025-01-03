package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetEnableVirtualSchemaProposal set a specific enableVirtualSchemaProposal in the store from its index
func (k Keeper) SetEnableVirtualSchemaProposal(ctx sdk.Context, enableVirtualSchemaProposal types.EnableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnableVirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&enableVirtualSchemaProposal)
	store.Set(types.EnableVirtualSchemaProposalKey(
		enableVirtualSchemaProposal.Id,
	), b)
}

// GetEnableVirtualSchemaProposal returns a enableVirtualSchemaProposal from its index
func (k Keeper) GetEnableVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) (val types.EnableVirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnableVirtualSchemaProposalKeyPrefix))

	b := store.Get(types.EnableVirtualSchemaProposalKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEnableVirtualSchemaProposal removes a enableVirtualSchemaProposal from the store
func (k Keeper) RemoveEnableVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnableVirtualSchemaProposalKeyPrefix))
	store.Delete(types.EnableVirtualSchemaProposalKey(
		index,
	))
}

// GetAllEnableVirtualSchemaProposal returns all enableVirtualSchemaProposal
func (k Keeper) GetAllEnableVirtualSchemaProposal(ctx sdk.Context) (list []types.EnableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnableVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EnableVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) EnableVirtualSchemaIterateActiveProposal(ctx sdk.Context, endTime time.Time, cb func(proposal types.EnableVirtualSchemaProposal) (stop bool)) {
	iterator := k.EnableVirtualSchemaActiveProposalQueryIterator(ctx, endTime)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		proposal, found := k.GetEnableVirtualSchemaProposal(ctx, val.Id)
		if !found {
			panic(fmt.Sprintf("proposal %d does not exist", &val.Id))
		}

		if cb(proposal) {
			break
		}
	}
}

func (k Keeper) EnableVirtualSchemaActiveProposalQueryIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveEnableVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	return iterator
}


// SetInactiveEnableVirtualSchemaProposal set a specific inactiveEnableVirtualSchemaProposal in the store from its index
func (k Keeper) SetInactiveEnableVirtualSchemaProposal(ctx sdk.Context, inactiveEnableVirtualSchemaProposal types.InactiveEnableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveEnableVirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&inactiveEnableVirtualSchemaProposal)
	store.Set(types.InactiveEnableVirtualSchemaProposalKey(
		inactiveEnableVirtualSchemaProposal.Id,
	), b)
}

// GetInactiveEnableVirtualSchemaProposal returns a inactiveEnableVirtualSchemaProposal from its index
func (k Keeper) GetInactiveEnableVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) (val types.InactiveEnableVirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveEnableVirtualSchemaProposalKeyPrefix))

	b := store.Get(types.InactiveEnableVirtualSchemaProposalKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInactiveEnableVirtualSchemaProposal removes a inactiveEnableVirtualSchemaProposal from the store
func (k Keeper) RemoveInactiveEnableVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveEnableVirtualSchemaProposalKeyPrefix))
	store.Delete(types.InactiveEnableVirtualSchemaProposalKey(
		index,
	))
}

// GetAllInactiveEnableVirtualSchemaProposal returns all inactiveEnableVirtualSchemaProposal
func (k Keeper) GetAllInactiveEnableVirtualSchemaProposal(ctx sdk.Context) (list []types.InactiveEnableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveEnableVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InactiveEnableVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetActiveEnableVirtualSchemaProposal set a specific activeEnableVirtualSchemaProposal in the store from its index
func (k Keeper) SetActiveEnableVirtualSchemaProposal(ctx sdk.Context, activeEnableVirtualSchemaProposal types.ActiveEnableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveEnableVirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&activeEnableVirtualSchemaProposal)
	store.Set(types.ActiveEnableVirtualSchemaProposalKey(
		activeEnableVirtualSchemaProposal.Id,
	), b)
}

// GetActiveEnableVirtualSchemaProposal returns a activeEnableVirtualSchemaProposal from its index
func (k Keeper) GetActiveEnableVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) (val types.ActiveEnableVirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveEnableVirtualSchemaProposalKeyPrefix))

	b := store.Get(types.ActiveEnableVirtualSchemaProposalKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActiveEnableVirtualSchemaProposal removes a activeEnableVirtualSchemaProposal from the store
func (k Keeper) RemoveActiveEnableVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveEnableVirtualSchemaProposalKeyPrefix))
	store.Delete(types.ActiveEnableVirtualSchemaProposalKey(
		index,
	))
}

// GetAllActiveEnableVirtualSchemaProposal returns all activeEnableVirtualSchemaProposal
func (k Keeper) GetAllActiveEnableVirtualSchemaProposal(ctx sdk.Context) (list []types.ActiveEnableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveEnableVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveEnableVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
