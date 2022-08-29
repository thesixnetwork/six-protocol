package keeper

import (
	"encoding/binary"
	
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/consume/types"
)

func (k Keeper) UpdateUseNft(ctx sdk.Context, use_nft types.UseNft) uint64 {
	// Get the current number of use_nfts in the store
	count := k.GetUseNftCount(ctx)
	// Assign an ID to the use_nft based on the number of use_nfts in the store
	use_nft.Id = count
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.UseNftKey))
	// Convert the use_nft ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, use_nft.Id)
	// Marshal the use_nft into bytes
	appendedValue := k.cdc.MustMarshal(&use_nft)
	// Insert the use_nft bytes using use_nft ID as a key
	store.Set(byteKey, appendedValue)
	// Update the use_nft count
	k.SetUseNftCount(ctx, count+1)
	return count
}

func (k Keeper) GetUseNftCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and UseNftCountKey (which is "UseNft-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.UseNftCountKey))
	// Convert the UseNftCountKey to bytes
	byteKey := []byte(types.UseNftCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found (for example, it's the first use_nft)
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetUseNftCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "blog") and UseNftCountKey (which is "UseNft-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.UseNftCountKey))
	// Convert the UseNftCountKey to bytes
	byteKey := []byte(types.UseNftCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the value of UseNft-count- to count
	store.Set(byteKey, bz)
}