package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetLockSchemaFee set a specific lockSchemaFee in the store from its index
func (k Keeper) SetLockSchemaFee(ctx sdk.Context, lockSchemaFee types.LockSchemaFee) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LockSchemaFeeKeyPrefix))
	b := k.cdc.MustMarshal(&lockSchemaFee)
	store.Set(types.LockSchemaFeeKey(
		lockSchemaFee.Id,
	), b)
}

// GetLockSchemaFee returns a lockSchemaFee from its index
func (k Keeper) GetLockSchemaFee(
	ctx sdk.Context,
	id string,
) (val types.LockSchemaFee, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LockSchemaFeeKeyPrefix))

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
	ctx sdk.Context,
	id string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LockSchemaFeeKeyPrefix))
	store.Delete(types.LockSchemaFeeKey(
		id,
	))
}

// GetAllLockSchemaFee returns all lockSchemaFee
func (k Keeper) GetAllLockSchemaFee(ctx sdk.Context) (list []types.LockSchemaFee) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LockSchemaFeeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LockSchemaFee
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
