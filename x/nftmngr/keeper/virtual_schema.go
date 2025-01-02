package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetVirtualSchema set a specific virSchema in the store from its index
func (k Keeper) SetVirtualSchema(ctx sdk.Context, virSchema types.VirtualSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaKeyPrefix))
	b := k.cdc.MustMarshal(&virSchema)
	store.Set(types.VirtualSchemaKey(
		virSchema.VirtualNftSchemaCode,
	), b)
}

// GetVirtualSchema returns a virSchema from its index
func (k Keeper) GetVirtualSchema(
	ctx sdk.Context,
	schemaCode string,
) (val types.VirtualSchema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaKeyPrefix))

	b := store.Get(types.VirtualSchemaKey(
		schemaCode,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVirtualSchema removes a virSchema from the store
func (k Keeper) RemoveVirtualSchema(
	ctx sdk.Context,
	schemaCode string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaKeyPrefix))
	store.Delete(types.VirtualSchemaKey(
		schemaCode,
	))
}

// GetAllVirtualSchema returns all virSchema
func (k Keeper) GetAllVirtualSchema(ctx sdk.Context) (list []types.VirtualSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VirtualSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetDisableVirtualSchema set a specific disableVirtualSchema in the store from its index
func (k Keeper) SetDisableVirtualSchemaProposal(ctx sdk.Context, disableVirtualSchema types.DisableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisableVirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&disableVirtualSchema)
	store.Set(types.DisableVirtualSchemaKey(
		disableVirtualSchema.Id,
	), b)
}

// GetDisableVirtualSchema returns a disableVirtualSchema from its index
func (k Keeper) GetDisableVirtualSchemaProposal(
	ctx sdk.Context,
	id string,
) (val types.DisableVirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisableVirtualSchemaProposalKeyPrefix))

	b := store.Get(types.DisableVirtualSchemaKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDisableVirtualSchema removes a disableVirtualSchema from the store
func (k Keeper) RemoveDisableVirtualSchemaProposal(
	ctx sdk.Context,
	id string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisableVirtualSchemaProposalKeyPrefix))
	store.Delete(types.DisableVirtualSchemaKey(
		id,
	))
}

// GetAllDisableVirtualSchema returns all disableVirtualSchema
func (k Keeper) GetAllDisableVirtualSchemaProposal(ctx sdk.Context) (list []types.DisableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisableVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DisableVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
