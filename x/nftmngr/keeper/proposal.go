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
	iterator := k.InactiveProposalQueryIterator(ctx, endTime)
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

func (k Keeper) IterateActiveProposal(ctx sdk.Context, endTime time.Time, cb func(proposal types.VirtualSchemaProposal) (stop bool)) {

	iterator := k.ActiveProposalQueryIterator(ctx, endTime)
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

func (k Keeper) ActiveProposalQueryIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	return iterator
}

func (k Keeper) InactiveProposalQueryIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	return iterator
}

func (k Keeper) IsProposalActive(ctx sdk.Context, proposal types.VirtualSchemaProposal) bool {
	return ctx.BlockTime().Before(proposal.VotingEndTime)
}
