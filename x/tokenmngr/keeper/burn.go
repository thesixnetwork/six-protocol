package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

// History burn
func (k Keeper) UpdateBurn(ctx sdk.Context, burn types.Burn) uint64 {
	// Get the current number of burns in the store
	count := k.GetBurnCount(ctx)
	// Assign an ID to the burn based on the number of burns in the store
	burn.Id = count
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BurnKey))
	// Convert the burn ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, burn.Id)
	// Marshal the burn into bytes
	appendedValue := k.cdc.MustMarshal(&burn)
	// Insert the burn bytes using burn ID as a key
	store.Set(byteKey, appendedValue)
	// Update the burn count
	k.SetBurnCount(ctx, count+1)
	return count
}

func (k Keeper) GetBurnCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and BurnCountKey (which is "Burn-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BurnCountKey))
	// Convert the BurnCountKey to bytes
	byteKey := []byte(types.BurnCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found (for example, it's the first burn)
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetBurnCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "burn") and BurnCountKey (which is "Burn-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BurnCountKey))
	// Convert the BurnCountKey to bytes
	byteKey := []byte(types.BurnCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the value of Burn-count- to count
	store.Set(byteKey, bz)
}
