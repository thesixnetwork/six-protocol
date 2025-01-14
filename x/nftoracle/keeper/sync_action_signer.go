package keeper

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

var ActiveSyncActionSignerQueuePrefix = []byte{0x05}

// GetSyncActionSignerCount get the total number of syncActionSigner
func (k Keeper) GetSyncActionSignerCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SyncActionSignerCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSyncActionSignerCount set the total number of syncActionSigner
func (k Keeper) SetSyncActionSignerCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SyncActionSignerCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSyncActionSigner appends a syncActionSigner in the store with a new id and update the count
func (k Keeper) AppendSyncActionSigner(
	ctx sdk.Context,
	syncActionSigner types.SyncActionSigner,
) uint64 {
	// Create the syncActionSigner
	count := k.GetSyncActionSignerCount(ctx)

	// Set the ID of the appended value
	syncActionSigner.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SyncActionSignerKey))
	appendedValue := k.cdc.MustMarshal(&syncActionSigner)
	store.Set(GetSyncActionSignerIDBytes(syncActionSigner.Id), appendedValue)

	// Update syncActionSigner count
	k.SetSyncActionSignerCount(ctx, count+1)

	return count
}

// SetSyncActionSigner set a specific syncActionSigner in the store
func (k Keeper) SetSyncActionSigner(ctx sdk.Context, syncActionSigner types.SyncActionSigner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SyncActionSignerKey))
	b := k.cdc.MustMarshal(&syncActionSigner)
	store.Set(GetSyncActionSignerIDBytes(syncActionSigner.Id), b)
}

// GetSyncActionSigner returns a syncActionSigner from its id
func (k Keeper) GetSyncActionSigner(ctx sdk.Context, id uint64) (val types.SyncActionSigner, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SyncActionSignerKey))
	b := store.Get(GetSyncActionSignerIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSyncActionSigner removes a syncActionSigner from the store
func (k Keeper) RemoveSyncActionSigner(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SyncActionSignerKey))
	store.Delete(GetSyncActionSignerIDBytes(id))
}

// GetAllSyncActionSigner returns all syncActionSigner
func (k Keeper) GetAllSyncActionSigner(ctx sdk.Context) (list []types.SyncActionSigner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SyncActionSignerKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

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

func (k Keeper) InsertActiveSyncActionSignerQueue(ctx sdk.Context, sync_id uint64, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	bz := GetSyncActionSignerIDBytes(sync_id)
	store.Set(ActiveSyncActionSignerQueueKey(sync_id, endTime), bz)
}

func (k Keeper) RemoveFromActiveSyncActionSignerQueue(ctx sdk.Context, sync_id uint64, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(ActiveSyncActionSignerQueueKey(sync_id, endTime))
}

func (k Keeper) IterateActiveSyncActionSignerQueue(ctx sdk.Context, endTime time.Time, cb func(syncRequest types.SyncActionSigner) (stop bool)) {
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

func (k Keeper) ActiveSyncActionSignerQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(ActiveSyncActionSignerQueuePrefix, sdk.PrefixEndBytes(ActiveSyncActionSignerByTimeKey(endTime)))
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
