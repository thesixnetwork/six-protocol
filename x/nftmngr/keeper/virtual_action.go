package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// TODO:: Feat(VirtualSchema)
// SetVirtualAction set a specific virtual in the store from its index
func (k Keeper) SetVirtualAction(ctx sdk.Context, virtual types.VirtualAction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualActionKeyPrefix))
	b := k.cdc.MustMarshal(&virtual)
	store.Set(types.VirtualActionKey(
		virtual.NftSchemaCode,
		virtual.Name,
	), b)
}

// TODO:: Feat(VirtualSchema)
// GetVirtualAction returns a virtual from its index
func (k Keeper) GetVirtualAction(
	ctx sdk.Context,
	code string,
	actionName string,
) (val types.VirtualAction, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualActionKeyPrefix))

	b := store.Get(types.VirtualActionKey(
		code,
		actionName,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// TODO:: Feat(VirtualSchema)
// RemoveVirtualAction removes a virtual from the store
func (k Keeper) RemoveVirtualAction(
	ctx sdk.Context,
	code string,
	actionName string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualActionKeyPrefix))
	store.Delete(types.VirtualActionKey(
		code,
		actionName,
	))
}

// TODO:: Feat(VirtualSchema)
// GetAllVirtualAction returns all virtual
func (k Keeper) GetAllVirtualAction(ctx sdk.Context) (list []types.VirtualAction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualActionKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VirtualAction
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
