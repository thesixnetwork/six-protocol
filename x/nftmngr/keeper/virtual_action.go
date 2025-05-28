package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	errormod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetVirtualAction set a specific virtual in the store from its index
func (k Keeper) SetVirtualAction(ctx context.Context, virtual types.VirtualAction) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualActionKeyPrefix))
	b := k.cdc.MustMarshal(&virtual)
	store.Set(types.VirtualActionKey(
		virtual.VirtualNftSchemaCode,
		virtual.Name,
	), b)
}

// GetVirtualAction returns a virtual from its index
func (k Keeper) GetVirtualAction(
	ctx context.Context,
	code string,
	actionName string,
) (val types.VirtualAction, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualActionKeyPrefix))

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

// RemoveVirtualAction removes a virtual from the store
func (k Keeper) RemoveVirtualAction(
	ctx context.Context,
	code string,
	actionName string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualActionKeyPrefix))
	store.Delete(types.VirtualActionKey(
		code,
		actionName,
	))
}

// GetAllVirtualAction returns all virtual
func (k Keeper) GetAllVirtualAction(ctx context.Context) (list []types.VirtualAction) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualActionKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VirtualAction
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) UpdateVirtualActionKeeper(ctx sdk.Context, nftSchemaName string, updateAction types.Action) error {
	_, found := k.GetVirtualSchema(ctx, nftSchemaName)
	if !found {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	// validate Action data
	err := ValidateVirutualAction(&updateAction)
	if err != nil {
		return errormod.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	k.SetVirtualAction(ctx, types.VirtualAction{
		VirtualNftSchemaCode: nftSchemaName,
		Name:                 updateAction.Name,
		Desc:                 updateAction.Desc,
		When:                 updateAction.When,
		Then:                 updateAction.Then,
		Params:               updateAction.Params,
		Disable:              updateAction.Disable,
		AllowedActioner:      updateAction.AllowedActioner,
	})

	return nil
}
