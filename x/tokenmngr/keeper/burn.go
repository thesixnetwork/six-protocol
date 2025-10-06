// Package keeper
//
//	deprecate
package keeper

import (
	"encoding/binary"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// UpdateBurn History burn
func (k Keeper) UpdateBurn(ctx sdk.Context, burn types.Burn) uint64 {
	count := k.GetBurnCount(ctx)
	burn.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnKey))
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, burn.Id)
	appendedValue := k.cdc.MustMarshal(&burn)
	store.Set(byteKey, appendedValue)
	k.SetBurnCount(ctx, count+1)
	return count
}

// GetBurnIDBytes returns the byte representation of the ID
func GetBurnIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// SetBurns is a special function used by upgrade module to set burns after upgrade
func (k Keeper) SetBurns(ctx sdk.Context, burns []types.Burn) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnKey))
	for _, burn := range burns {
		b := k.cdc.MustMarshal(&burn)
		store.Set(GetBurnIDBytes(burn.Id), b)
	}
}

func (k Keeper) GetBurnCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnKey))
	byteKey := []byte(types.BurnCountKey)
	bz := store.Get(byteKey)
	// Return zero if the count value is not found (for example, it's the first burn)
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetBurnCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BurnKey))
	byteKey := []byte(types.BurnCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}
