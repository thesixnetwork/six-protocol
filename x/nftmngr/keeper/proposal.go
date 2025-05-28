package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetVirtualSchemaProposal set a specific virtualSchemaProposal in the store from its index
func (k Keeper) SetVirtualSchemaProposal(ctx context.Context, virtualSchemaProposal types.VirtualSchemaProposal) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&virtualSchemaProposal)
	store.Set(types.VirtualSchemaProposalKey(
		virtualSchemaProposal.Id,
	), b)
}

// GetVirtualSchemaProposal returns a virtualSchemaProposal from its index
func (k Keeper) GetVirtualSchemaProposal(
	ctx context.Context,
	id string,
) (val types.VirtualSchemaProposal, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))

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
	ctx context.Context,
	id string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))
	store.Delete(types.VirtualSchemaProposalKey(
		id,
	))
}

// GetAllVirtualSchemaProposal returns all virtualSchemaProposal
func (k Keeper) GetAllVirtualSchemaProposal(ctx context.Context) (list []types.VirtualSchemaProposal) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetActiveVirtualSchemaProposal set a specific activeVirtualSchemaProposal in the store from its index
func (k Keeper) SetActiveVirtualSchemaProposal(ctx context.Context, activeVirtualSchemaProposal types.ActiveVirtualSchemaProposal) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))

	b := k.cdc.MustMarshal(&activeVirtualSchemaProposal)
	store.Set(types.ActiveVirtualSchemaProposalKey(
		activeVirtualSchemaProposal.Id,
	), b)
}

// GetActiveVirtualSchemaProposal returns a activeVirtualSchemaProposal from its index
func (k Keeper) GetActiveVirtualSchemaProposal(
	ctx context.Context,
	index string,
) (val types.ActiveVirtualSchemaProposal, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))

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
	ctx context.Context,
	index string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))
	store.Delete(types.ActiveVirtualSchemaProposalKey(
		index,
	))
}

// GetAllActiveVirtualSchemaProposal returns all activeVirtualSchemaProposal
func (k Keeper) GetAllActiveVirtualSchemaProposal(ctx context.Context) (list []types.ActiveVirtualSchemaProposal) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetInactiveVirtualSchemaProposal set a specific inactiveVirtualSchemaProposal in the store from its index
func (k Keeper) SetInactiveVirtualSchemaProposal(ctx context.Context, inactiveVirtualSchemaProposal types.InactiveVirtualSchemaProposal) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&inactiveVirtualSchemaProposal)
	store.Set(types.InactiveVirtualSchemaProposalKey(
		inactiveVirtualSchemaProposal.Id,
	), b)
}

// GetInactiveVirtualSchemaProposal returns a inactiveVirtualSchemaProposal from its index
func (k Keeper) GetInactiveVirtualSchemaProposal(
	ctx context.Context,
	index string,
) (val types.InactiveVirtualSchemaProposal, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))
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
	ctx context.Context,
	index string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))
	store.Delete(types.InactiveVirtualSchemaProposalKey(
		index,
	))
}

// GetAllInactiveVirtualSchemaProposal returns all inactiveVirtualSchemaProposal
func (k Keeper) GetAllInactiveVirtualSchemaProposal(ctx context.Context) (list []types.InactiveVirtualSchemaProposal) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InactiveVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

/*
	###########################################################################################################
	###########################################################################################################
	############################################## SUPPORT METHOD #############################################
	###########################################################################################################
	###########################################################################################################
*/

func (k Keeper) VirtualSchemaActiveProposalQueryIterator(ctx context.Context, endTime time.Time) storetypes.Iterator {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	return iterator
}

func (k Keeper) InactiveProposalCreateVirtualSchemaQueryIterator(ctx context.Context, endTime time.Time) storetypes.Iterator {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	return iterator
}

func (k Keeper) IsProposalActive(goCtx context.Context, proposal types.VirtualSchemaProposal) bool {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return ctx.BlockTime().Before(proposal.VotingEndTime)
}

func (k Keeper) IterateInactiveProposal(ctx context.Context, endTime time.Time, cb func(proposal types.VirtualSchemaProposal) (stop bool)) {
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

func (k Keeper) IterateActiveVirtualSchemaProposal(ctx context.Context, endTime time.Time, cb func(proposal types.VirtualSchemaProposal) (stop bool)) {
	iterator := k.VirtualSchemaActiveProposalQueryIterator(ctx, endTime)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		proposal, found := k.GetVirtualSchemaProposal(ctx, val.Id)
		if !found {
			panic(fmt.Sprintf("proposal %d does not exist", &val.Id))
		}

		// check time to end proposal
		if !k.IsProposalActive(ctx, proposal) {
			k.RemoveActiveVirtualSchemaProposal(ctx, val.Id)
			k.SetInactiveVirtualSchemaProposal(ctx, types.InactiveVirtualSchemaProposal(val))
			if err := k.processSchemaFee(ctx, proposal, false); err != nil {
				k.Logger().Error("failed to process schema fee", "error", err)
			}
		}

		if cb(proposal) {
			break
		}
	}
}
