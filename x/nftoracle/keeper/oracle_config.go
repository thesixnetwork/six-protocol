package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

// SetOracleConfig set oracleConfig in the store
func (k Keeper) SetOracleConfig(ctx sdk.Context, oracleConfig types.OracleConfig) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleConfigKey))
	b := k.cdc.MustMarshal(&oracleConfig)
	store.Set([]byte{0}, b)
}

// GetOracleConfig returns oracleConfig
func (k Keeper) GetOracleConfig(ctx sdk.Context) (val types.OracleConfig, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleConfigKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOracleConfig removes oracleConfig from the store
func (k Keeper) RemoveOracleConfig(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OracleConfigKey))
	store.Delete([]byte{0})
}
