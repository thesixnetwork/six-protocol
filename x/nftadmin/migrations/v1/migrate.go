// Package v1
package v1

import (
	"cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftadmin/types"
	v1 "github.com/thesixnetwork/six-protocol/x/nftadmin/types/v1"
)

func MigrateStore(ctx sdk.Context, storeService store.KVStoreService, cdc codec.BinaryCodec) error {
	store := runtime.KVStoreAdapter(storeService.OpenKVStore(ctx))
	burnStore := prefix.NewStore(store, types.KeyPrefix(types.AuthorizationKey))
	return restructAutorization(burnStore, cdc)
}

// restruct authorization
func restructAutorization(store storetypes.KVStore, cdc codec.BinaryCodec) error {
	var v1Auth v1.Authorization
	b := store.Get([]byte{0})
	if b == nil {
		return errorsmod.Wrapf(types.ErrAuthorizationNotFound, "invalid key to query autorization")
	}

	cdc.MustUnmarshal(b, &v1Auth)

	var auth types.Authorization

	b = cdc.MustMarshal(&auth)

	store.Set([]byte{0}, b)

	return nil
}
