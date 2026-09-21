package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

func TestMsgServerChangeOrgOwner(t *testing.T) {
	owner := sample.AccAddress()
	newOwner := sample.AccAddress()

	t.Run("invalid creator address", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.ChangeOrgOwner(ctx, &types.MsgChangeOrgOwner{
			Creator:    "invalid_address",
			ToNewOwner: newOwner,
			OrgName:    "org",
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidAddress)
	})

	t.Run("invalid new owner address", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.ChangeOrgOwner(ctx, &types.MsgChangeOrgOwner{
			Creator:    owner,
			ToNewOwner: "invalid_address",
			OrgName:    "org",
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidAddress)
	})

	t.Run("organization not found", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.ChangeOrgOwner(ctx, &types.MsgChangeOrgOwner{
			Creator:    owner,
			ToNewOwner: newOwner,
			OrgName:    "missing",
		})
		require.ErrorIs(t, err, types.ErrOrganizationNotFound)
	})

	t.Run("creator is not the organization owner", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetOrganization(sdk.UnwrapSDKContext(ctx), types.Organization{
			Name:  "org",
			Owner: sample.AccAddress(),
		})
		_, err := ms.ChangeOrgOwner(ctx, &types.MsgChangeOrgOwner{
			Creator:    owner,
			ToNewOwner: newOwner,
			OrgName:    "org",
		})
		require.ErrorIs(t, err, types.ErrOrganizationOwner)
	})

	t.Run("success", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		k.SetOrganization(sdkCtx, types.Organization{
			Name:  "org",
			Owner: owner,
		})
		resp, err := ms.ChangeOrgOwner(ctx, &types.MsgChangeOrgOwner{
			Creator:    owner,
			ToNewOwner: newOwner,
			OrgName:    "org",
		})
		require.NoError(t, err)
		require.Equal(t, "org", resp.OrgName)
		require.Equal(t, owner, resp.OldOwner)
		require.Equal(t, newOwner, resp.NewOwner)

		org, found := k.GetOrganization(sdkCtx, "org")
		require.True(t, found)
		require.Equal(t, newOwner, org.Owner)
	})
}

func TestMsgServerChangeSchemaOwner(t *testing.T) {
	owner := sample.AccAddress()
	newOwner := sample.AccAddress()

	t.Run("invalid creator address", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.ChangeSchemaOwner(ctx, &types.MsgChangeSchemaOwner{
			Creator:       "invalid_address",
			NewOwner:      newOwner,
			NftSchemaCode: "schema",
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidAddress)
	})

	t.Run("invalid new owner address", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.ChangeSchemaOwner(ctx, &types.MsgChangeSchemaOwner{
			Creator:       owner,
			NewOwner:      "invalid_address",
			NftSchemaCode: "schema",
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidAddress)
	})

	t.Run("schema not found", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.ChangeSchemaOwner(ctx, &types.MsgChangeSchemaOwner{
			Creator:       owner,
			NewOwner:      newOwner,
			NftSchemaCode: "missing",
		})
		require.ErrorIs(t, err, types.ErrSchemaDoesNotExists)
	})

	t.Run("creator is not the schema owner", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetNFTSchema(sdk.UnwrapSDKContext(ctx), types.NFTSchema{
			Code:  "schema",
			Owner: sample.AccAddress(),
		})
		_, err := ms.ChangeSchemaOwner(ctx, &types.MsgChangeSchemaOwner{
			Creator:       owner,
			NewOwner:      newOwner,
			NftSchemaCode: "schema",
		})
		require.ErrorIs(t, err, types.ErrCreatorDoesNotMatch)
	})

	t.Run("success", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		k.SetNFTSchema(sdkCtx, types.NFTSchema{
			Code:  "schema",
			Owner: owner,
		})
		resp, err := ms.ChangeSchemaOwner(ctx, &types.MsgChangeSchemaOwner{
			Creator:       owner,
			NewOwner:      newOwner,
			NftSchemaCode: "schema",
		})
		require.NoError(t, err)
		require.Equal(t, "schema", resp.NftSchemaCode)
		require.Equal(t, newOwner, resp.NewOwner)

		schema, found := k.GetNFTSchema(sdkCtx, "schema")
		require.True(t, found)
		require.Equal(t, newOwner, schema.Owner)
	})
}
