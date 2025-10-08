package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetLockSchemaFee set a specific lockSchemaFee in the store from its index
func (k Keeper) SetLockSchemaFee(ctx context.Context, lockSchemaFee types.LockSchemaFee) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LockSchemaFeeKeyPrefix))

	b := k.cdc.MustMarshal(&lockSchemaFee)
	store.Set(types.LockSchemaFeeKey(
		lockSchemaFee.Id,
	), b)
}

// GetLockSchemaFee returns a lockSchemaFee from its index
func (k Keeper) GetLockSchemaFee(
	ctx context.Context,
	id string,
) (val types.LockSchemaFee, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LockSchemaFeeKeyPrefix))

	b := store.Get(types.LockSchemaFeeKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLockSchemaFee removes a lockSchemaFee from the store
func (k Keeper) RemoveLockSchemaFee(
	ctx context.Context,
	id string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LockSchemaFeeKeyPrefix))
	store.Delete(types.LockSchemaFeeKey(
		id,
	))
}

// GetAllLockSchemaFee returns all lockSchemaFee
func (k Keeper) GetAllLockSchemaFee(ctx context.Context) (list []types.LockSchemaFee) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LockSchemaFeeKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LockSchemaFee
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
