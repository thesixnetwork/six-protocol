package keeper

import (
	"slices"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/thesixnetwork/six-protocol/v4/x/erc20/types"
)

var isTrue = []byte("0x01")

const addressLength = 42

// GetParams returns the total set of erc20 parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	enableErc20 := k.IsERC20Enabled(ctx)
	dynamicPrecompiles := k.getDynamicPrecompiles(ctx)
	nativePrecompiles := k.getNativePrecompiles(ctx)
	return types.NewParams(enableErc20, nativePrecompiles, dynamicPrecompiles)
}

// SetParams sets the erc20 parameters to the store.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	// sort param slices to ensure determinism
	slices.Sort(params.DynamicPrecompiles)
	slices.Sort(params.NativePrecompiles)

	if err := params.Validate(); err != nil {
		return err
	}

	// update the codehash for enabled or disabled dynamic precompiles
	if err := k.UpdateCodeHash(ctx, params.DynamicPrecompiles); err != nil {
		return err
	}

	k.setERC20Enabled(ctx, params.EnableErc20)
	k.setDynamicPrecompiles(ctx, params.DynamicPrecompiles)
	k.setNativePrecompiles(ctx, params.NativePrecompiles)
	return nil
}

// UpdateCodeHash updates the code hash for enabled/disabled dynamic precompiles.
func (k Keeper) UpdateCodeHash(ctx sdk.Context, updatedDynamicPrecompiles []string) error {
	oldDynamicPrecompiles := k.getDynamicPrecompiles(ctx)
	disabledPrecompiles, enabledPrecompiles := types.GetDisabledAndEnabledPrecompiles(oldDynamicPrecompiles, updatedDynamicPrecompiles)

	for _, precompile := range disabledPrecompiles {
		if err := k.UnRegisterERC20CodeHash(ctx, precompile); err != nil {
			return err
		}
	}

	for _, precompile := range enabledPrecompiles {
		if err := k.RegisterERC20CodeHash(ctx, common.HexToAddress(precompile)); err != nil {
			return err
		}
	}
	return nil
}

// IsERC20Enabled returns true if the module logic is enabled
func (k Keeper) IsERC20Enabled(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ParamStoreKeyEnableErc20)
}

// setERC20Enabled sets the EnableERC20 param in the store
func (k Keeper) setERC20Enabled(ctx sdk.Context, enable bool) {
	store := ctx.KVStore(k.storeKey)
	if enable {
		store.Set(types.ParamStoreKeyEnableErc20, isTrue)
		return
	}
	store.Delete(types.ParamStoreKeyEnableErc20)
}

// setDynamicPrecompiles sets the DynamicPrecompiles param in the store
func (k Keeper) setDynamicPrecompiles(ctx sdk.Context, dynamicPrecompiles []string) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 0, addressLength*len(dynamicPrecompiles))
	for _, str := range dynamicPrecompiles {
		bz = append(bz, []byte(str)...)
	}
	store.Set(types.ParamStoreKeyDynamicPrecompiles, bz)
}

// getDynamicPrecompiles returns the DynamicPrecompiles param from the store
func (k Keeper) getDynamicPrecompiles(ctx sdk.Context) (dynamicPrecompiles []string) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamStoreKeyDynamicPrecompiles)
	for i := 0; i < len(bz); i += addressLength {
		dynamicPrecompiles = append(dynamicPrecompiles, string(bz[i:i+addressLength]))
	}
	return dynamicPrecompiles
}

// setNativePrecompiles sets the NativePrecompiles param in the store
func (k Keeper) setNativePrecompiles(ctx sdk.Context, nativePrecompiles []string) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 0, addressLength*len(nativePrecompiles))
	for _, str := range nativePrecompiles {
		bz = append(bz, []byte(str)...)
	}
	store.Set(types.ParamStoreKeyNativePrecompiles, bz)
}

// getNativePrecompiles returns the NativePrecompiles param from the store
func (k Keeper) getNativePrecompiles(ctx sdk.Context) (nativePrecompiles []string) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamStoreKeyNativePrecompiles)
	for i := 0; i < len(bz); i += addressLength {
		nativePrecompiles = append(nativePrecompiles, string(bz[i:i+addressLength]))
	}
	return nativePrecompiles
}
