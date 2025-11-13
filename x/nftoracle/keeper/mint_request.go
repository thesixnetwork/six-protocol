package keeper

import (
	"context"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
)

var ActiveMintRequestQueuePrefix = []byte{0x01}

var lenTime = len(sdk.FormatTimeBytes(time.Now()))

// GetMintRequestCount get the total number of mintRequest
func (k Keeper) GetMintRequestCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	byteKey := types.KeyPrefix(types.MintRequestCountKey)
	store := prefix.NewStore(storeAdapter, byteKey)
	bz := store.Get(byteKey)
	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMintRequestCount set the total number of mintRequest
func (k Keeper) SetMintRequestCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	byteKey := types.KeyPrefix(types.MintRequestCountKey)
	store := prefix.NewStore(storeAdapter, byteKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMintRequest appends a mintRequest in the store with a new id and update the count
func (k Keeper) AppendMintRequest(
	ctx context.Context,
	mintRequest types.MintRequest,
) uint64 {
	// Create the mintRequest
	count := k.GetMintRequestCount(ctx)

	// Set the ID of the appended value
	mintRequest.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintRequestKey))

	appendedValue := k.cdc.MustMarshal(&mintRequest)
	store.Set(GetMintRequestIDBytes(mintRequest.Id), appendedValue)

	// Update mintRequest count
	k.SetMintRequestCount(ctx, count+1)

	return count
}

// SetMintRequest set a specific mintRequest in the store
func (k Keeper) SetMintRequest(ctx context.Context, mintRequest types.MintRequest) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintRequestKey))
	b := k.cdc.MustMarshal(&mintRequest)
	store.Set(GetMintRequestIDBytes(mintRequest.Id), b)
}

// GetMintRequest returns a mintRequest from its id
func (k Keeper) GetMintRequest(ctx context.Context, id uint64) (val types.MintRequest, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintRequestKey))
	b := store.Get(GetMintRequestIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMintRequest removes a mintRequest from the store
func (k Keeper) RemoveMintRequest(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintRequestKey))
	store.Delete(GetMintRequestIDBytes(id))
}

// GetAllMintRequest returns all mintRequest
func (k Keeper) GetAllMintRequest(ctx context.Context) (list []types.MintRequest) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MintRequestKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MintRequest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMintRequestIDBytes returns the byte representation of the ID
func GetMintRequestIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMintRequestIDFromBytes returns ID in uint64 format from a byte array
func GetMintRequestIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) InsertActiveMintRequestQueue(ctx context.Context, requestID uint64, endTime time.Time) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := GetMintRequestIDBytes(requestID)
	storeAdapter.Set(ActiveMintRequestQueueKey(requestID, endTime), bz)
}

func (k Keeper) RemoveFromActiveMintRequestQueue(ctx context.Context, requestID uint64, endTime time.Time) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	storeAdapter.Delete(ActiveMintRequestQueueKey(requestID, endTime))
}

func (k Keeper) IterateActiveMintRequestsQueue(ctx context.Context, endTime time.Time, cb func(mintRequest types.MintRequest) (stop bool)) {
	iterator := k.ActiveMintRequestQueueIterator(ctx, endTime)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		requestID, _ := SplitActiveMintRequestQueueKey(iterator.Key())
		mintRequest, found := k.GetMintRequest(ctx, requestID)
		if !found {
			panic(fmt.Sprintf("mintRequest %d does not exist", requestID))
		}

		if cb(mintRequest) {
			break
		}
	}
}

func (k Keeper) ActiveMintRequestQueueIterator(ctx context.Context, endTime time.Time) storetypes.Iterator {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return storeAdapter.Iterator(ActiveMintRequestQueuePrefix, storetypes.PrefixEndBytes(ActiveMintRequestByTimeKey(endTime)))
}

func ActiveMintRequestQueueKey(requestID uint64, endTime time.Time) []byte {
	return append(ActiveMintRequestByTimeKey(endTime), GetMintRequestIDBytes(requestID)...)
}

func ActiveMintRequestByTimeKey(endTime time.Time) []byte {
	return append(ActiveMintRequestQueuePrefix, sdk.FormatTimeBytes(endTime)...)
}

func SplitActiveMintRequestQueueKey(key []byte) (requestID uint64, endTime time.Time) {
	return splitKeyWithTime(key)
}

func splitKeyWithTime(key []byte) (requestID uint64, endTime time.Time) {
	kv.AssertKeyLength(key[1:], 8+lenTime)

	endTime, err := sdk.ParseTimeBytes(key[1 : 1+lenTime])
	if err != nil {
		panic(err)
	}

	requestID = GetMintRequestIDFromBytes(key[1+lenTime:])
	return
}
