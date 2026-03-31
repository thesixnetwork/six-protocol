package keeper

import (
	"context"
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/v4/x/precisebank/types"
)

// GetFractionalBalance returns the fractional balance for an address.
func (k Keeper) GetFractionalBalance(ctx context.Context, address sdk.AccAddress) sdkmath.Int {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.FractionalBalancePrefix)

	bz := store.Get(types.FractionalBalanceKey(address))
	if bz == nil {
		return sdkmath.ZeroInt()
	}

	var bal sdkmath.Int
	if err := bal.Unmarshal(bz); err != nil {
		panic(fmt.Errorf("failed to unmarshal fractional balance: %w", err))
	}

	return bal
}

// SetFractionalBalance sets the fractional balance for an address.
func (k Keeper) SetFractionalBalance(ctx context.Context, address sdk.AccAddress, amount sdkmath.Int) {
	if address.Empty() {
		panic(errors.New("address cannot be empty"))
	}

	if amount.IsZero() {
		k.DeleteFractionalBalance(ctx, address)
		return
	}

	if err := types.ValidateFractionalAmount(amount); err != nil {
		panic(fmt.Errorf("amount is invalid: %w", err))
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.FractionalBalancePrefix)

	amountBytes, err := amount.Marshal()
	if err != nil {
		panic(fmt.Errorf("failed to marshal fractional balance: %w", err))
	}

	store.Set(types.FractionalBalanceKey(address), amountBytes)
}

// DeleteFractionalBalance deletes the fractional balance for an address.
func (k Keeper) DeleteFractionalBalance(ctx context.Context, address sdk.AccAddress) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.FractionalBalancePrefix)
	store.Delete(types.FractionalBalanceKey(address))
}

// IterateFractionalBalances iterates over all fractional balances in the store.
func (k Keeper) IterateFractionalBalances(
	ctx context.Context,
	cb func(address sdk.AccAddress, amount sdkmath.Int) (stop bool),
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.FractionalBalancePrefix)

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		address := sdk.AccAddress(iterator.Key())

		var amount sdkmath.Int
		if err := amount.Unmarshal(iterator.Value()); err != nil {
			panic(fmt.Errorf("failed to unmarshal fractional balance: %w", err))
		}

		if cb(address, amount) {
			break
		}
	}
}

// GetTotalSumFractionalBalances returns the sum of all fractional balances.
func (k Keeper) GetTotalSumFractionalBalances(ctx context.Context) sdkmath.Int {
	sum := sdkmath.ZeroInt()

	k.IterateFractionalBalances(ctx, func(_ sdk.AccAddress, amount sdkmath.Int) bool {
		sum = sum.Add(amount)
		return false
	})

	return sum
}

// GetRemainderAmount returns the remainder amount from the store.
func (k Keeper) GetRemainderAmount(ctx context.Context) sdkmath.Int {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	bz := storeAdapter.Get(types.RemainderBalanceKey)
	if bz == nil {
		return sdkmath.ZeroInt()
	}

	var remainder sdkmath.Int
	if err := remainder.Unmarshal(bz); err != nil {
		panic(fmt.Errorf("failed to unmarshal remainder amount: %w", err))
	}

	return remainder
}

// SetRemainderAmount sets the remainder amount in the store.
func (k Keeper) SetRemainderAmount(ctx context.Context, remainder sdkmath.Int) {
	if remainder.IsZero() {
		k.DeleteRemainderAmount(ctx)
		return
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	remainderBytes, err := remainder.Marshal()
	if err != nil {
		panic(fmt.Errorf("failed to marshal remainder amount: %w", err))
	}

	storeAdapter.Set(types.RemainderBalanceKey, remainderBytes)
}

// DeleteRemainderAmount deletes the remainder amount from the store.
func (k Keeper) DeleteRemainderAmount(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	storeAdapter.Delete(types.RemainderBalanceKey)
}
