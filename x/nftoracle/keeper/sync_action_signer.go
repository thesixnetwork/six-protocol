package keeper

import (
	"context"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var ActiveSyncActionSignerQueuePrefix = []byte{0x05}

// GetSyncActionSignerCount get the total number of syncActionSigner
func (k Keeper) GetSyncActionSignerCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	byteKey := types.KeyPrefix(types.SyncActionSignerCountKey)
	store := prefix.NewStore(storeAdapter, byteKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSyncActionSignerCount set the total number of syncActionSigner
func (k Keeper) SetSyncActionSignerCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	byteKey := types.KeyPrefix(types.SyncActionSignerCountKey)
	store := prefix.NewStore(storeAdapter, byteKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSyncActionSigner appends a syncActionSigner in the store with a new id and update the count
func (k Keeper) AppendSyncActionSigner(
	ctx context.Context,
	syncActionSigner types.SyncActionSigner,
) uint64 {
	// Create the syncActionSigner
	count := k.GetSyncActionSignerCount(ctx)

	// Set the ID of the appended value
	syncActionSigner.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SyncActionSignerKey))
	appendedValue := k.cdc.MustMarshal(&syncActionSigner)
	store.Set(GetSyncActionSignerIDBytes(syncActionSigner.Id), appendedValue)

	// Update syncActionSigner count
	k.SetSyncActionSignerCount(ctx, count+1)

	return count
}

// SetSyncActionSigner set a specific syncActionSigner in the store
func (k Keeper) SetSyncActionSigner(ctx context.Context, syncActionSigner types.SyncActionSigner) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SyncActionSignerKey))
	b := k.cdc.MustMarshal(&syncActionSigner)
	store.Set(GetSyncActionSignerIDBytes(syncActionSigner.Id), b)
}

// GetSyncActionSigner returns a syncActionSigner from its id
func (k Keeper) GetSyncActionSigner(ctx context.Context, id uint64) (val types.SyncActionSigner, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SyncActionSignerKey))
	b := store.Get(GetSyncActionSignerIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSyncActionSigner removes a syncActionSigner from the store
func (k Keeper) RemoveSyncActionSigner(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SyncActionSignerKey))
	store.Delete(GetSyncActionSignerIDBytes(id))
}

// GetAllSyncActionSigner returns all syncActionSigner
func (k Keeper) GetAllSyncActionSigner(ctx context.Context) (list []types.SyncActionSigner) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SyncActionSignerKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SyncActionSigner
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSyncActionSignerIDBytes returns the byte representation of the ID
func GetSyncActionSignerIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSyncActionSignerIDFromBytes returns ID in uint64 format from a byte array
func GetSyncActionSignerIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) InsertActiveSyncActionSignerQueue(ctx context.Context, sync_id uint64, endTime time.Time) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := GetSyncActionSignerIDBytes(sync_id)
	storeAdapter.Set(ActiveSyncActionSignerQueueKey(sync_id, endTime), bz)
}

func (k Keeper) RemoveFromActiveSyncActionSignerQueue(ctx context.Context, sync_id uint64, endTime time.Time) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	storeAdapter.Delete(ActiveSyncActionSignerQueueKey(sync_id, endTime))
}

func (k Keeper) IterateActiveSyncActionSignerQueue(ctx context.Context, endTime time.Time, cb func(syncRequest types.SyncActionSigner) (stop bool)) {
	iterator := k.ActiveSyncActionSignerQueueIterator(ctx, endTime)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		sync_id, _ := SplitActiveSyncActionSignerQueueKey(iterator.Key())
		syncRequest, found := k.GetSyncActionSigner(ctx, sync_id)
		if !found {
			panic(fmt.Sprintf("syncRequest %d does not exist", sync_id))
		}

		if cb(syncRequest) {
			break
		}
	}
}

func (k Keeper) ActiveSyncActionSignerQueueIterator(ctx context.Context, endTime time.Time) storetypes.Iterator {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return storeAdapter.Iterator(ActiveSyncActionSignerQueuePrefix, storetypes.PrefixEndBytes(ActiveSyncActionSignerByTimeKey(endTime)))
}

func ActiveSyncActionSignerQueueKey(sync_id uint64, endTime time.Time) []byte {
	return append(ActiveSyncActionSignerByTimeKey(endTime), GetSyncActionSignerIDBytes(sync_id)...)
}

func ActiveSyncActionSignerByTimeKey(endTime time.Time) []byte {
	return append(ActiveSyncActionSignerQueuePrefix, sdk.FormatTimeBytes(endTime)...)
}

func SplitActiveSyncActionSignerQueueKey(key []byte) (sync_id uint64, endTime time.Time) {
	return splitKeyWithTime(key)
}
