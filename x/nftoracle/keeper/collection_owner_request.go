package keeper

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

var (
	ActiveVerifyCollectionRequestQueuePrefix = []byte{0x03}
)

// GetCollectionOwnerRequestCount get the total number of collectionOwnerRequest
func (k Keeper) GetCollectionOwnerRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CollectionOwnerRequestCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCollectionOwnerRequestCount set the total number of collectionOwnerRequest
func (k Keeper) SetCollectionOwnerRequestCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CollectionOwnerRequestCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendCollectionOwnerRequest appends a collectionOwnerRequest in the store with a new id and update the count
func (k Keeper) AppendCollectionOwnerRequest(
	ctx sdk.Context,
	collectionOwnerRequest types.CollectionOwnerRequest,
) uint64 {
	// Create the collectionOwnerRequest
	count := k.GetCollectionOwnerRequestCount(ctx)

	// Set the ID of the appended value
	collectionOwnerRequest.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionOwnerRequestKey))
	appendedValue := k.cdc.MustMarshal(&collectionOwnerRequest)
	store.Set(GetCollectionOwnerRequestIDBytes(collectionOwnerRequest.Id), appendedValue)

	// Update collectionOwnerRequest count
	k.SetCollectionOwnerRequestCount(ctx, count+1)

	return count
}

// SetCollectionOwnerRequest set a specific collectionOwnerRequest in the store
func (k Keeper) SetCollectionOwnerRequest(ctx sdk.Context, collectionOwnerRequest types.CollectionOwnerRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionOwnerRequestKey))
	b := k.cdc.MustMarshal(&collectionOwnerRequest)
	store.Set(GetCollectionOwnerRequestIDBytes(collectionOwnerRequest.Id), b)
}

// GetCollectionOwnerRequest returns a collectionOwnerRequest from its id
func (k Keeper) GetCollectionOwnerRequest(ctx sdk.Context, id uint64) (val types.CollectionOwnerRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionOwnerRequestKey))
	b := store.Get(GetCollectionOwnerRequestIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCollectionOwnerRequest removes a collectionOwnerRequest from the store
func (k Keeper) RemoveCollectionOwnerRequest(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionOwnerRequestKey))
	store.Delete(GetCollectionOwnerRequestIDBytes(id))
}

// GetAllCollectionOwnerRequest returns all collectionOwnerRequest
func (k Keeper) GetAllCollectionOwnerRequest(ctx sdk.Context) (list []types.CollectionOwnerRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionOwnerRequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CollectionOwnerRequest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCollectionOwnerRequestIDBytes returns the byte representation of the ID
func GetCollectionOwnerRequestIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCollectionOwnerRequestIDFromBytes returns ID in uint64 format from a byte array
func GetCollectionOwnerRequestIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) InsertActiveVerifyCollectionOwnerRequestQueue(ctx sdk.Context, requestID uint64, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	bz := GetCollectionOwnerRequestIDBytes(requestID)
	store.Set(ActiveVerifyCollectionOwnerQueueKey(requestID, endTime), bz)
}

func (k Keeper) RemoveFromActiveVerifyCollectionOwnerQueue(ctx sdk.Context, requestID uint64, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(ActiveVerifyCollectionOwnerQueueKey(requestID, endTime))
}

func (k Keeper) IterateActiveVerifyCollectionOwnersQueue(ctx sdk.Context, endTime time.Time, cb func(mintRequest types.CollectionOwnerRequest) (stop bool)) {
	iterator := k.ActiveVerifyCollectionOwnerQueueIterator(ctx, endTime)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		requestID, _ := SplitActiveVerifyCollectionOwnerQueueKey(iterator.Key())
		verifyRequest, found := k.GetCollectionOwnerRequest(ctx, requestID)
		if !found {
			panic(fmt.Sprintf("verifyRequest %d does not exist", requestID))
		}

		if cb(verifyRequest) {
			break
		}
	}
}

func (k Keeper) ActiveVerifyCollectionOwnerQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(ActiveVerifyCollectionRequestQueuePrefix, sdk.PrefixEndBytes(ActiveVerifyCollectionOwnerByTimeKey(endTime)))
}

func ActiveVerifyCollectionOwnerQueueKey(requestID uint64, endTime time.Time) []byte {
	return append(ActiveVerifyCollectionOwnerByTimeKey(endTime), GetCollectionOwnerRequestIDBytes(requestID)...)
}
func ActiveVerifyCollectionOwnerByTimeKey(endTime time.Time) []byte {
	return append(ActiveVerifyCollectionRequestQueuePrefix, sdk.FormatTimeBytes(endTime)...)
}

func SplitActiveVerifyCollectionOwnerQueueKey(key []byte) (requestID uint64, endTime time.Time) {
	return splitKeyWithTime(key)
}
