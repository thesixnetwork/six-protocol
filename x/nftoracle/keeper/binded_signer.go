package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetBindedSigner set a specific bindedSigner in the store from its index
func (k Keeper) SetBindedSigner(ctx context.Context, bindedSigner types.BindedSigner) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BindedSignerKeyPrefix))
	b := k.cdc.MustMarshal(&bindedSigner)
	store.Set(types.BindedSignerKey(
		bindedSigner.OwnerAddress,
	), b)
}

// GetBindedSigner returns a bindedSigner from its index
func (k Keeper) GetBindedSigner(
	ctx context.Context,
	ownerAddress string,
) (val types.BindedSigner, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BindedSignerKeyPrefix))

	b := store.Get(types.BindedSignerKey(
		ownerAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBindedSigner removes a bindedSigner from the store
func (k Keeper) RemoveBindedSigner(
	ctx context.Context,
	ownerAddress string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BindedSignerKeyPrefix))
	store.Delete(types.BindedSignerKey(
		ownerAddress,
	))
}

// Remove Sepecific Signer fron bindedSignerList
func (k Keeper) RemoveSignerFromBindedSignerList(
	ctx context.Context,
	ownerAddress string,
	signerAddress string,
) {
	bindedSigner, found := k.GetBindedSigner(ctx, ownerAddress)
	if !found {
		return
	}
	for i, signer := range bindedSigner.Signers {
		if signer.ActorAddress == signerAddress {
			bindedSigner.Signers = append(bindedSigner.Signers[:i], bindedSigner.Signers[i+1:]...)
			break
		}
	}

	k.SetBindedSigner(ctx, bindedSigner)
}

// GetAllBindedSigner returns all bindedSigner
func (k Keeper) GetAllBindedSigner(ctx context.Context) (list []types.BindedSigner) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BindedSignerKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BindedSigner
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
