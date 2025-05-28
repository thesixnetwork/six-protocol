package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftadmin/types"

	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// revoke permission
// sixnftd tx nftadmin revoke-permission minter 6nft1q3566qhn4hnjf8l0zt90daew2ade2yc6l5usaq --from alice
func (k msgServer) RevokePermission(goCtx context.Context, msg *types.MsgRevokePermission) (*types.MsgRevokePermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auth, found := k.Keeper.GetAuthorization(ctx)
	if !found {
		return nil, types.ErrAuthorizationNotFound
	}

	if msg.Creator != auth.RootAdmin {
		return nil, types.ErrUnauthorized
	}

	if auth.Permissions == nil {
		return nil, types.ErrNoPermissions
	}

	_, err := sdk.AccAddressFromBech32(msg.Revokee)
	if err != nil {
		return nil, types.ErrInvalidGrantee
	}

	list := auth.Permissions.GetPermissionAddressByKey(msg.Name)
	if list == nil {
		return nil, errormod.Wrapf(types.ErrNoPermissionsForName, "no permissions for name %s", msg.Name)
	}

	if list.Addresses == nil || len(list.Addresses) == 0 {
		return nil, errormod.Wrapf(types.ErrNoPermissionsForName, "no permissions for name %s", msg.Name)
	}

	removed := false

	for i, addr := range list.Addresses {
		if addr == msg.Revokee {
			list.Addresses = append(list.Addresses[:i], list.Addresses[i+1:]...)
			removed = true
			break
		}
	}

	if !removed {
		return nil, errormod.Wrapf(types.ErrGranteeNotFoundForName, "grantee %s not found for name %s", msg.Revokee, msg.Name)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRevokePermission,
			sdk.NewAttribute(types.AttributeKeyRevokePermissionType, msg.Name),
			sdk.NewAttribute(types.AttributeKeyRevokePermissionAddress, msg.Revokee),
			sdk.NewAttribute(types.AttributeKeyRevokePermissionStatus, "success"),
		),
	})

	k.Keeper.SetAuthorization(ctx, auth)

	return &types.MsgRevokePermissionResponse{
		Revokee: msg.Revokee,
	}, nil
}
