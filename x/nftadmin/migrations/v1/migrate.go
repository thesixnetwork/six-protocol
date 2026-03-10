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

	"github.com/thesixnetwork/six-protocol/v4/x/nftadmin/types"
	v1 "github.com/thesixnetwork/six-protocol/v4/x/nftadmin/types/v1"
)

const FIVENT_ROOT_ADMIN = "6x1w4h88d93rqezzyqpvdhfe08xp6732m5f3e9evf"

func MigrateStore(ctx sdk.Context, storeService store.KVStoreService, cdc codec.BinaryCodec) error {
	store := runtime.KVStoreAdapter(storeService.OpenKVStore(ctx))
	authStore := prefix.NewStore(store, types.KeyPrefix(types.AuthorizationKey))

	chainID := ctx.ChainID()

	if chainID == "testnet" || chainID == "fivenet" {
		return fivenetHotfix(authStore, cdc)
	}

	return restructAuthorization(authStore, cdc)
}

// restructAuthorization migrates v1 Authorization structure to v2
func restructAuthorization(store storetypes.KVStore, cdc codec.BinaryCodec) error {
	var v1Auth v1.Authorization
	b := store.Get([]byte{0})
	if b == nil {
		return errorsmod.Wrapf(types.ErrAuthorizationNotFound, "invalid key to query authorization")
	}

	cdc.MustUnmarshal(b, &v1Auth)

	var v2Auth types.Authorization

	if v1Auth.Permissions != nil {
		for _, perm := range v1Auth.Permissions.Permissions {
			var addresses []string
			if perm.Addresses != nil {
				addresses = perm.Addresses.Addresses
			}

			v2Auth.Permissions = append(v2Auth.Permissions, &types.Permission{
				Name:      perm.Name,
				Addresses: addresses,
			})
		}
	}

	v2Auth.RootAdmin = v1Auth.RootAdmin

	b = cdc.MustMarshal(&v2Auth)

	store.Set([]byte{0}, b)

	return nil
}

func fivenetHotfix(store storetypes.KVStore, cdc codec.BinaryCodec) error {
	var v2Auth types.Authorization
	v2Auth.RootAdmin = FIVENT_ROOT_ADMIN
	v2Auth.Permissions = []*types.Permission{
		{
			Name:      "nft_fee_admin",
			Addresses: []string{FIVENT_ROOT_ADMIN},
		},
	}
	b := cdc.MustMarshal(&v2Auth)
	store.Set([]byte{0}, b)
	return nil
}
