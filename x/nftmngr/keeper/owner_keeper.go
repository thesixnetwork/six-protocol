package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	errormod "cosmossdk.io/errors"
)

func (k Keeper) GetSchemaOwner(ctx context.Context, nftSchemaName string) (string, error) {
	var ownerAddress string

	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return "", errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}
	ownerAddress = schema.Owner
	return ownerAddress, nil
}

func (k Keeper) IsSchemaOwner(ctx context.Context, nftSchemaName, inputAddress string) (bool, error) {
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return false, errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if schema.Owner != inputAddress {
		return false, nil
	}

	return true, nil
}

func (k Keeper) ChangeOrgOwner(ctx context.Context, creator, newOwner, orgName string) error {
	// get the organization
	organization, found := k.GetOrganization(ctx, orgName)
	if !found {
		return errormod.Wrap(types.ErrOrganizationNotFound, orgName)
	}

	if organization.Owner != creator {
		return errormod.Wrap(types.ErrOrganizationOwner, creator)
	}

	// change the owner
	organization.Owner = newOwner

	// save the organization
	k.SetOrganization(ctx, organization)
	return nil
}

func (k Keeper) ChangeSchemaOwner(ctx context.Context, creator, newOwner, nftSchemaName string) error {
	// Retrieve schema data
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, creator)
	}

	// Check if the creator is the same as the current owner
	if creator != schema.Owner {
		return errormod.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	// Change the owner
	schema.Owner = newOwner

	// Save the schema
	k.SetNFTSchema(ctx, schema)

	return nil
}
