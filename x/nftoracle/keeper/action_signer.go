package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetActionSigner set a specific actionSigner in the store from its index
func (k Keeper) SetActionSigner(ctx context.Context, actionSigner types.ActionSigner) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionSignerKeyPrefix))
	b := k.cdc.MustMarshal(&actionSigner)
	store.Set(types.ActionSignerKey(
		actionSigner.ActorAddress,
		actionSigner.OwnerAddress,
	), b)
}

// GetActionSigner returns a actionSigner from its index
func (k Keeper) GetActionSigner(
	ctx context.Context,
	actorAddress string,
	ownerAddress string,
) (val types.ActionSigner, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionSignerKeyPrefix))

	b := store.Get(types.ActionSignerKey(
		actorAddress,
		ownerAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActionSigner removes a actionSigner from the store
func (k Keeper) RemoveActionSigner(
	ctx context.Context,
	actorAddress string,
	ownerAddress string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionSignerKeyPrefix))
	store.Delete(types.ActionSignerKey(
		actorAddress,
		ownerAddress,
	))
}

// GetAllActionSigner returns all actionSigner
func (k Keeper) GetAllActionSigner(ctx context.Context) (list []types.ActionSigner) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActionSignerKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionSigner
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
