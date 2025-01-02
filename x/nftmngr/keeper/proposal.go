package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetVirtualSchemaProposal set a specific virtualSchemaProposal in the store from its index
func (k Keeper) SetVirtualSchemaProposal(ctx sdk.Context, virtualSchemaProposal types.VirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&virtualSchemaProposal)
	store.Set(types.VirtualSchemaProposalKey(
		virtualSchemaProposal.Id,
	), b)
}

// GetVirtualSchemaProposal returns a virtualSchemaProposal from its index
func (k Keeper) GetVirtualSchemaProposal(
	ctx sdk.Context,
	id string,
) (val types.VirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))

	b := store.Get(types.VirtualSchemaProposalKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVirtualSchemaProposal removes a virtualSchemaProposal from the store
func (k Keeper) RemoveVirtualSchemaProposal(
	ctx sdk.Context,
	id string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))
	store.Delete(types.VirtualSchemaProposalKey(
		id,
	))
}

// GetAllVirtualSchemaProposal returns all virtualSchemaProposal
func (k Keeper) GetAllVirtualSchemaProposal(ctx sdk.Context) (list []types.VirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetActiveVirtualSchemaProposal set a specific activeVirtualSchemaProposal in the store from its index
func (k Keeper) SetActiveVirtualSchemaProposal(ctx sdk.Context, activeVirtualSchemaProposal types.ActiveVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&activeVirtualSchemaProposal)
	store.Set(types.ActiveVirtualSchemaProposalKey(
		activeVirtualSchemaProposal.Id,
	), b)
}

// GetActiveVirtualSchemaProposal returns a activeVirtualSchemaProposal from its index
func (k Keeper) GetActiveVirtualSchemaProposal(
	ctx sdk.Context,
	index string,
) (val types.ActiveVirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))

	b := store.Get(types.ActiveVirtualSchemaProposalKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActiveVirtualSchemaProposal removes a activeVirtualSchemaProposal from the store
func (k Keeper) RemoveActiveVirtualSchemaProposal(
	ctx sdk.Context,
	index string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))
	store.Delete(types.ActiveVirtualSchemaProposalKey(
		index,
	))
}

// GetAllActiveVirtualSchemaProposal returns all activeVirtualSchemaProposal
func (k Keeper) GetAllActiveVirtualSchemaProposal(ctx sdk.Context) (list []types.ActiveVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) IterateInactiveProposal(ctx sdk.Context, endTime time.Time, cb func(proposal types.VirtualSchemaProposal) (stop bool)) {
	iterator := k.InactiveProposalCreateVirtualSchemaQueryIterator(ctx, endTime)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InactiveVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		proposal, found := k.GetVirtualSchemaProposal(ctx, val.Id)
		if !found {
			panic(fmt.Sprintf("proposal %d does not exist", &val.Id))
		}

		if cb(proposal) {
			break
		}
	}
}

func (k Keeper) IterateActiveProposalCreateVirtualSchema(ctx sdk.Context, endTime time.Time, cb func(proposal types.VirtualSchemaProposal) (stop bool)) {
	iterator := k.ActiveProposalCreateVirtualSchemaQueryIterator(ctx, endTime)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		proposal, found := k.GetVirtualSchemaProposal(ctx, val.Id)
		if !found {
			panic(fmt.Sprintf("proposal %d does not exist", &val.Id))
		}

		if cb(proposal) {
			break
		}
	}
}

// SetInactiveVirtualSchemaProposal set a specific inactiveVirtualSchemaProposal in the store from its index
func (k Keeper) SetInactiveVirtualSchemaProposal(ctx sdk.Context, inactiveVirtualSchemaProposal types.InactiveVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&inactiveVirtualSchemaProposal)
	store.Set(types.InactiveVirtualSchemaProposalKey(
		inactiveVirtualSchemaProposal.Id,
	), b)
}

// GetInactiveVirtualSchemaProposal returns a inactiveVirtualSchemaProposal from its index
func (k Keeper) GetInactiveVirtualSchemaProposal(
	ctx sdk.Context,
	index string,
) (val types.InactiveVirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))

	b := store.Get(types.InactiveVirtualSchemaProposalKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInactiveVirtualSchemaProposal removes a inactiveVirtualSchemaProposal from the store
func (k Keeper) RemoveInactiveVirtualSchemaProposal(
	ctx sdk.Context,
	index string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))
	store.Delete(types.InactiveVirtualSchemaProposalKey(
		index,
	))
}

// GetAllInactiveVirtualSchemaProposal returns all inactiveVirtualSchemaProposal
func (k Keeper) GetAllInactiveVirtualSchemaProposal(ctx sdk.Context) (list []types.InactiveVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InactiveVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ActiveProposalCreateVirtualSchemaQueryIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	return iterator
}

func (k Keeper) InactiveProposalCreateVirtualSchemaQueryIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	return iterator
}

func (k Keeper) IsProposalActive(ctx sdk.Context, proposal types.VirtualSchemaProposal) bool {
	return ctx.BlockTime().Before(proposal.VotingEndTime)
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

