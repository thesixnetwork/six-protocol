package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetExecutorOfSchema set a specific executorOfSchema in the store from its index
func (k Keeper) SetExecutorOfSchema(ctx sdk.Context, executorOfSchema types.ExecutorOfSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutorOfSchemaKeyPrefix))
	b := k.cdc.MustMarshal(&executorOfSchema)
	store.Set(types.ExecutorOfSchemaKey(
		executorOfSchema.NftSchemaCode,
	), b)
}

// GetExecutorOfSchema returns a executorOfSchema from its index
func (k Keeper) GetExecutorOfSchema(
	ctx sdk.Context,
	nftSchemaCode string,
) (val types.ExecutorOfSchema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutorOfSchemaKeyPrefix))

	b := store.Get(types.ExecutorOfSchemaKey(
		nftSchemaCode,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveExecutorOfSchema removes a executorOfSchema from the store
func (k Keeper) RemoveExecutorOfSchema(
	ctx sdk.Context,
	nftSchemaCode string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutorOfSchemaKeyPrefix))
	store.Delete(types.ExecutorOfSchemaKey(
		nftSchemaCode,
	))
}

// GetAllExecutorOfSchema returns all executorOfSchema
func (k Keeper) GetAllExecutorOfSchema(ctx sdk.Context) (list []types.ExecutorOfSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutorOfSchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ExecutorOfSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetActionExecutor set a specific actionExecutor in the store from its index
func (k Keeper) SetActionExecutor(ctx sdk.Context, actionExecutor types.ActionExecutor) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionExecutorKeyPrefix))
	b := k.cdc.MustMarshal(&actionExecutor)
	store.Set(types.ActionExecutorKey(
		actionExecutor.NftSchemaCode,
		actionExecutor.ExecutorAddress,
	), b)
}

// GetActionExecutor returns a actionExecutor from its index
func (k Keeper) GetActionExecutor(
	ctx sdk.Context,
	nftSchemaCode string,
	executorAddress string,
) (val types.ActionExecutor, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionExecutorKeyPrefix))

	b := store.Get(types.ActionExecutorKey(
		nftSchemaCode,
		executorAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActionExecutor removes a actionExecutor from the store
func (k Keeper) RemoveActionExecutor(
	ctx sdk.Context,
	nftSchemaCode string,
	executorAddress string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionExecutorKeyPrefix))
	store.Delete(types.ActionExecutorKey(
		nftSchemaCode,
		executorAddress,
	))
}

// GetAllActionExecutor returns all actionExecutor
func (k Keeper) GetAllActionExecutor(ctx sdk.Context) (list []types.ActionExecutor) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionExecutorKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionExecutor
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) AddActionExecutor(ctx sdk.Context, creator, nftSchemaName, executorAddress string) error {
	schema, foundNftSchema := k.GetNFTSchema(ctx, nftSchemaName)
	_, foundVirtualSchema := k.GetVirtualSchema(ctx, nftSchemaName)

	if !foundNftSchema && !foundVirtualSchema {
		return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if foundNftSchema && creator != schema.Owner {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if foundVirtualSchema && creator != k.GetModuleAddress().String() {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only module account can do this process")
	}

	// Check if the value already exists
	_, isFound := k.GetActionExecutor(
		ctx,
		nftSchemaName,
		executorAddress,
	)

	if isFound {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Action Executor already exists")
	}

	actionExecutor := types.ActionExecutor{
		Creator:         creator,
		NftSchemaCode:   nftSchemaName,
		ExecutorAddress: executorAddress,
	}

	val, found := k.GetExecutorOfSchema(ctx, nftSchemaName)
	if !found {
		val = types.ExecutorOfSchema{
			NftSchemaCode:   nftSchemaName,
			ExecutorAddress: []string{},
		}
	}

	// set executorOfSchema
	val.ExecutorAddress = append(val.ExecutorAddress, executorAddress)

	k.SetExecutorOfSchema(ctx, types.ExecutorOfSchema{
		NftSchemaCode:   val.NftSchemaCode,
		ExecutorAddress: val.ExecutorAddress,
	})

	k.SetActionExecutor(ctx, actionExecutor)

	return nil
}

// RemoveActionExecutor removes a actionExecutor from the store
func (k Keeper) DelActionExecutor(ctx sdk.Context, creator, nftSchemaName, executorAddress string) error {
	// Retrieve the schema
	schema, foundNftSchema := k.GetNFTSchema(ctx, nftSchemaName)
	_, foundVirtualSchema := k.GetVirtualSchema(ctx, nftSchemaName)

	if !foundNftSchema && !foundVirtualSchema {
		return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if foundNftSchema && creator != schema.Owner {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if foundVirtualSchema && creator != k.GetModuleAddress().String() {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only module account can do this process")
	}

	// Check if the value exists
	_, isFound := k.GetActionExecutor(
		ctx,
		nftSchemaName,
		executorAddress,
	)
	if !isFound {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	k.RemoveActionExecutor(
		ctx,
		nftSchemaName,
		executorAddress,
	)

	val, found := k.GetExecutorOfSchema(ctx, nftSchemaName)
	if !found {
		val = types.ExecutorOfSchema{
			NftSchemaCode:   nftSchemaName,
			ExecutorAddress: []string{},
		}
	}

	// remove executorOfSchema
	for i, executor := range val.ExecutorAddress {
		if executor == executorAddress {
			val.ExecutorAddress = append(val.ExecutorAddress[:i], val.ExecutorAddress[i+1:]...)
			break
		}
	}

	k.SetExecutorOfSchema(ctx, types.ExecutorOfSchema{
		NftSchemaCode:   val.NftSchemaCode,
		ExecutorAddress: val.ExecutorAddress,
	})

	return nil
}
