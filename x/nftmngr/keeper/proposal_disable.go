package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetDisableVirtualSchema set a specific disableVirtualSchema in the store from its index
func (k Keeper) SetDisableVirtualSchemaProposal(ctx sdk.Context, disableVirtualSchema types.DisableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisableVirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&disableVirtualSchema)
	store.Set(types.DisableVirtualSchemaKey(
		disableVirtualSchema.Id,
	), b)
}

// GetDisableVirtualSchema returns a disableVirtualSchema from its index
func (k Keeper) GetDisableVirtualSchemaProposal(
	ctx sdk.Context,
	id string,
) (val types.DisableVirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisableVirtualSchemaProposalKeyPrefix))

	b := store.Get(types.DisableVirtualSchemaKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDisableVirtualSchema removes a disableVirtualSchema from the store
func (k Keeper) RemoveDisableVirtualSchemaProposal(
	ctx sdk.Context,
	id string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisableVirtualSchemaProposalKeyPrefix))
	store.Delete(types.DisableVirtualSchemaKey(
		id,
	))
}

// GetAllDisableVirtualSchema returns all disableVirtualSchema
func (k Keeper) GetAllDisableVirtualSchemaProposal(ctx sdk.Context) (list []types.DisableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisableVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DisableVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}


// SetActiveDisableVirtualSchemaProposal set a specific activeDisableVirtualSchemaProposal in the store from its index
func (k Keeper) SetActiveDisableVirtualSchemaProposal(ctx sdk.Context, activeDisableVirtualSchemaProposal types.ActiveDisableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDisableVirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&activeDisableVirtualSchemaProposal)
	store.Set(types.ActiveDisableVirtualSchemaProposalKey(
		activeDisableVirtualSchemaProposal.Id,
	), b)
}

// GetActiveDisableVirtualSchemaProposal returns a activeDisableVirtualSchemaProposal from its index
func (k Keeper) GetActiveDisableVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) (val types.ActiveDisableVirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDisableVirtualSchemaProposalKeyPrefix))

	b := store.Get(types.ActiveDisableVirtualSchemaProposalKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActiveDisableVirtualSchemaProposal removes a activeDisableVirtualSchemaProposal from the store
func (k Keeper) RemoveActiveDisableVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDisableVirtualSchemaProposalKeyPrefix))
	store.Delete(types.ActiveDisableVirtualSchemaProposalKey(
		index,
	))
}

// GetAllActiveDisableVirtualSchemaProposal returns all activeDisableVirtualSchemaProposal
func (k Keeper) GetAllActiveDisableVirtualSchemaProposal(ctx sdk.Context) (list []types.ActiveDisableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDisableVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveDisableVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetInactiveDisableVirtualSchemaProposal set a specific inactiveDisableVirtualSchemaProposal in the store from its index
func (k Keeper) SetInactiveDisableVirtualSchemaProposal(ctx sdk.Context, inactiveDisableVirtualSchemaProposal types.InactiveDisableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveDisableVirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&inactiveDisableVirtualSchemaProposal)
	store.Set(types.InactiveDisableVirtualSchemaProposalKey(
		inactiveDisableVirtualSchemaProposal.Id,
	), b)
}

// GetInactiveDisableVirtualSchemaProposal returns a inactiveDisableVirtualSchemaProposal from its index
func (k Keeper) GetInactiveDisableVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) (val types.InactiveDisableVirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveDisableVirtualSchemaProposalKeyPrefix))

	b := store.Get(types.InactiveDisableVirtualSchemaProposalKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}


// RemoveInactiveDisableVirtualSchemaProposal removes a inactiveDisableVirtualSchemaProposal from the store
func (k Keeper) RemoveInactiveDisableVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveDisableVirtualSchemaProposalKeyPrefix))
	store.Delete(types.InactiveDisableVirtualSchemaProposalKey(
		index,
	))
}

// GetAllInactiveDisableVirtualSchemaProposal returns all inactiveDisableVirtualSchemaProposal
func (k Keeper) GetAllInactiveDisableVirtualSchemaProposal(ctx sdk.Context) (list []types.InactiveDisableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveDisableVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InactiveDisableVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) DisableVirtualSchemaIterateActiveProposal(ctx sdk.Context, endTime time.Time, cb func(proposal types.DisableVirtualSchemaProposal) (stop bool)) {
	iterator := k.DisableVirtualSchemaActiveProposalQueryIterator(ctx, endTime)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		proposal, found := k.GetDisableVirtualSchemaProposal(ctx, val.Id)
		if !found {
			panic(fmt.Sprintf("proposal %d does not exist", &val.Id))
		}

		if cb(proposal) {
			break
		}
	}
}

func (k Keeper) DisableVirtualSchemaActiveProposalQueryIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDisableVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	return iterator
}
