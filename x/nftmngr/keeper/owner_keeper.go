package keeper

import (
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) GetSchemaOwner(ctx sdk.Context, nftSchemaName string) (string, error) {
	var ownerAddress string

	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return "", sdkerrors.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}
	ownerAddress = schema.Owner
	return ownerAddress, nil
}

func (k Keeper) IsSchemaOwner(ctx sdk.Context, nftSchemaName, inputAddress string) (bool, error) {
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return false, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if schema.Owner != inputAddress {
		return false, nil
	}

	return true, nil
}

func (k Keeper) ChangeOrgOwner(ctx sdk.Context, creator, newOwner, orgName string) error {
	// get the organization
	organization, found := k.GetOrganization(ctx, orgName)
	if !found {
		return sdkerrors.Wrap(types.ErrOrganizationNotFound, orgName)
	}

	if organization.Owner != creator {
		return sdkerrors.Wrap(types.ErrOrganizationOwner, creator)
	}

	// change the owner
	organization.Owner = newOwner

	// save the organization
	k.SetOrganization(ctx, organization)
	return nil
}

func (k Keeper) ChangeSchemaOwner(ctx sdk.Context, creator, newOwner, nftSchemaName string) error {
	// Retreive schema data
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, creator)
	}

	// Check if the creator is the same as the current owner
	if creator != schema.Owner {
		return sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	// Change the owner
	schema.Owner = newOwner

	// Save the schema
	k.SetNFTSchema(ctx, schema)

	return nil
}
