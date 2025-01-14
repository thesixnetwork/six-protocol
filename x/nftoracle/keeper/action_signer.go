package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

// SetActionSigner set a specific actionSigner in the store from its index
func (k Keeper) SetActionSigner(ctx sdk.Context, actionSigner types.ActionSigner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerKeyPrefix))
	b := k.cdc.MustMarshal(&actionSigner)
	store.Set(types.ActionSignerKey(
		actionSigner.ActorAddress,
		actionSigner.OwnerAddress,
	), b)
}

// GetActionSigner returns a actionSigner from its index
func (k Keeper) GetActionSigner(
	ctx sdk.Context,
	actorAddress string,
	ownerAddress string,
) (val types.ActionSigner, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerKeyPrefix))

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
	ctx sdk.Context,
	actorAddress string,
	ownerAddress string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerKeyPrefix))
	store.Delete(types.ActionSignerKey(
		actorAddress,
		ownerAddress,
	))
}

// GetAllActionSigner returns all actionSigner
func (k Keeper) GetAllActionSigner(ctx sdk.Context) (list []types.ActionSigner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionSignerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionSigner
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
