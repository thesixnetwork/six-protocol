package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetOracleConfig set oracleConfig in the store
func (k Keeper) SetOracleConfig(ctx context.Context, oracleConfig types.OracleConfig) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OracleConfigKey))
	b := k.cdc.MustMarshal(&oracleConfig)
	store.Set([]byte{0}, b)
}

// GetOracleConfig returns oracleConfig
func (k Keeper) GetOracleConfig(ctx context.Context) (val types.OracleConfig, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OracleConfigKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOracleConfig removes oracleConfig from the store
func (k Keeper) RemoveOracleConfig(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OracleConfigKey))
	store.Delete([]byte{0})
}
