package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nftadmin module sentinel errors
var (
	ErrAuthorizationNotFound  = sdkerrors.Register(ModuleName, 1, "authorization not found")
	ErrUnauthorized           = sdkerrors.Register(ModuleName, 2, "unauthorized")
	ErrNoPermissions          = sdkerrors.Register(ModuleName, 3, "no permissions")
	ErrNoPermissionsForName   = sdkerrors.Register(ModuleName, 4, "no permissions for name")
	ErrGranteeExists          = sdkerrors.Register(ModuleName, 5, "grantee exists")
	ErrGranteeNotFoundForName = sdkerrors.Register(ModuleName, 6, "grantee not found for name")
	ErrInvalidGrantee         = sdkerrors.Register(ModuleName, 7, "invalid grantee format")
	ErrInvalidRevokee         = sdkerrors.Register(ModuleName, 8, "invalid revokee format")
	ErrGranteeAlreadyExists   = sdkerrors.Register(ModuleName, 9, "grantee already exists")
)
