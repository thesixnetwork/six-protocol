package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	OracleLastVoteHeightPrefix = []byte{0x20} // Prefix for oracle last vote height tracking
)

// SetOracleLastVoteHeight sets the last block height when an oracle voted
func (k Keeper) SetOracleLastVoteHeight(ctx context.Context, oracle sdk.AccAddress, height int64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, OracleLastVoteHeightPrefix)

	key := oracle.Bytes()
	heightBytes := sdk.Uint64ToBigEndian(uint64(height))
	store.Set(key, heightBytes)
}

// GetOracleLastVoteHeight gets the last block height when an oracle voted
func (k Keeper) GetOracleLastVoteHeight(ctx context.Context, oracle sdk.AccAddress) int64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, OracleLastVoteHeightPrefix)

	key := oracle.Bytes()
	bz := store.Get(key)
	if bz == nil {
		return 0 // Never voted before
	}

	height := sdk.BigEndianToUint64(bz)
	return int64(height)
}

// DeleteOracleLastVoteHeight removes the last vote height for an oracle
func (k Keeper) DeleteOracleLastVoteHeight(ctx context.Context, oracle sdk.AccAddress) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, OracleLastVoteHeightPrefix)

	key := oracle.Bytes()
	store.Delete(key)
}

// IsOracleGaslessEnabled returns true if gasless oracle voting is enabled
func (k Keeper) IsOracleGaslessEnabled(ctx context.Context) bool {

	return true
}

// CleanupOldVoteHeights removes old vote height records to prevent store bloat
// Should be called periodically (e.g., in EndBlocker or BeginBlocker)
func (k Keeper) CleanupOldVoteHeights(ctx context.Context, maxAge int64) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, OracleLastVoteHeightPrefix)

	currentHeight := sdkCtx.BlockHeight()
	cutoffHeight := currentHeight - maxAge

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var keysToDelete [][]byte
	for ; iterator.Valid(); iterator.Next() {
		heightBytes := iterator.Value()
		if len(heightBytes) == 8 {
			height := int64(sdk.BigEndianToUint64(heightBytes))
			if height < cutoffHeight {
				keysToDelete = append(keysToDelete, iterator.Key())
			}
		}
	}

	// Delete old entries
	for _, key := range keysToDelete {
		store.Delete(key)
	}
}
