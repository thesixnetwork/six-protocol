package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

// SetActionSignerConfig set a specific actionSignerConfig in the store from its index
func (k Keeper) SetActionSignerConfig(ctx sdk.Context, actionSignerConfig types.ActionSignerConfig) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerConfigKeyPrefix))
	b := k.cdc.MustMarshal(&actionSignerConfig)
	store.Set(types.ActionSignerConfigKey(
		actionSignerConfig.Chain,
	), b)
}

// GetActionSignerConfig returns a actionSignerConfig from its index
func (k Keeper) GetActionSignerConfig(
	ctx sdk.Context,
	chain string,
) (val types.ActionSignerConfig, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerConfigKeyPrefix))

	b := store.Get(types.ActionSignerConfigKey(
		chain,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActionSignerConfig removes a actionSignerConfig from the store
func (k Keeper) RemoveActionSignerConfig(
	ctx sdk.Context,
	chain string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerConfigKeyPrefix))
	store.Delete(types.ActionSignerConfigKey(
		chain,
	))
}

// GetAllActionSignerConfig returns all actionSignerConfig
func (k Keeper) GetAllActionSignerConfig(ctx sdk.Context) (list []types.ActionSignerConfig) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerConfigKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionSignerConfig
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
