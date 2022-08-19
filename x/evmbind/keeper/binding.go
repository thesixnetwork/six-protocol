package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/evmbind/types"
)

// SetBinding set a specific binding in the store from its index
func (k Keeper) SetBinding(ctx sdk.Context, binding types.Binding) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BindingKeyPrefix))
	b := k.cdc.MustMarshal(&binding)
	store.Set(types.BindingKey(
		binding.EthAddress,
	), b)
}

// GetBinding returns a binding from its index
func (k Keeper) GetBinding(
	ctx sdk.Context,
	ethAddress string,

) (val types.Binding, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BindingKeyPrefix))

	b := store.Get(types.BindingKey(
		ethAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBinding removes a binding from the store
func (k Keeper) RemoveBinding(
	ctx sdk.Context,
	ethAddress string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BindingKeyPrefix))
	store.Delete(types.BindingKey(
		ethAddress,
	))
}

// GetAllBinding returns all binding
func (k Keeper) GetAllBinding(ctx sdk.Context) (list []types.Binding) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BindingKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Binding
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
