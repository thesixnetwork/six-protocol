package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

func TestMsgServerCreateActionExecutor(t *testing.T) {
	owner := sample.AccAddress()
	executor := sample.AccAddress()

	t.Run("invalid creator address", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.CreateActionExecutor(ctx, &types.MsgCreateActionExecutor{
			Creator:         "invalid_address",
			NftSchemaCode:   "schema",
			ExecutorAddress: executor,
		})
		require.Error(t, err)
	})

	t.Run("invalid executor address", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.CreateActionExecutor(ctx, &types.MsgCreateActionExecutor{
			Creator:         owner,
			NftSchemaCode:   "schema",
			ExecutorAddress: "invalid_address",
		})
		require.Error(t, err)
	})

	t.Run("schema not found", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.CreateActionExecutor(ctx, &types.MsgCreateActionExecutor{
			Creator:         owner,
			NftSchemaCode:   "missing",
			ExecutorAddress: executor,
		})
		require.ErrorIs(t, err, types.ErrSchemaDoesNotExists)
	})

	t.Run("creator is not the schema owner", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetNFTSchema(sdk.UnwrapSDKContext(ctx), types.NFTSchema{
			Code:  "schema",
			Owner: sample.AccAddress(),
		})
		_, err := ms.CreateActionExecutor(ctx, &types.MsgCreateActionExecutor{
			Creator:         owner,
			NftSchemaCode:   "schema",
			ExecutorAddress: executor,
		})
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	})

	t.Run("success and duplicate", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		k.SetNFTSchema(sdkCtx, types.NFTSchema{
			Code:  "schema",
			Owner: owner,
		})

		resp, err := ms.CreateActionExecutor(ctx, &types.MsgCreateActionExecutor{
			Creator:         owner,
			NftSchemaCode:   "schema",
			ExecutorAddress: executor,
		})
		require.NoError(t, err)
		require.Equal(t, "schema", resp.NftSchemaCode)
		require.Equal(t, executor, resp.ExecutorAddress)

		actionExecutor, found := k.GetActionExecutor(sdkCtx, "schema", executor)
		require.True(t, found)
		require.Equal(t, owner, actionExecutor.Creator)

		executorOfSchema, found := k.GetExecutorOfSchema(sdkCtx, "schema")
		require.True(t, found)
		require.Contains(t, executorOfSchema.ExecutorAddress, executor)

		// creating the same executor again fails
		_, err = ms.CreateActionExecutor(ctx, &types.MsgCreateActionExecutor{
			Creator:         owner,
			NftSchemaCode:   "schema",
			ExecutorAddress: executor,
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	})
}

func TestMsgServerDeleteActionExecutor(t *testing.T) {
	owner := sample.AccAddress()
	executor := sample.AccAddress()

	t.Run("schema not found", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.DeleteActionExecutor(ctx, &types.MsgDeleteActionExecutor{
			Creator:         owner,
			NftSchemaCode:   "missing",
			ExecutorAddress: executor,
		})
		require.ErrorIs(t, err, types.ErrSchemaDoesNotExists)
	})

	t.Run("executor not found", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetNFTSchema(sdk.UnwrapSDKContext(ctx), types.NFTSchema{
			Code:  "schema",
			Owner: owner,
		})
		_, err := ms.DeleteActionExecutor(ctx, &types.MsgDeleteActionExecutor{
			Creator:         owner,
			NftSchemaCode:   "schema",
			ExecutorAddress: executor,
		})
		require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
	})

	t.Run("success", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		k.SetNFTSchema(sdkCtx, types.NFTSchema{
			Code:  "schema",
			Owner: owner,
		})
		_, err := ms.CreateActionExecutor(ctx, &types.MsgCreateActionExecutor{
			Creator:         owner,
			NftSchemaCode:   "schema",
			ExecutorAddress: executor,
		})
		require.NoError(t, err)

		_, err = ms.DeleteActionExecutor(ctx, &types.MsgDeleteActionExecutor{
			Creator:         owner,
			NftSchemaCode:   "schema",
			ExecutorAddress: executor,
		})
		require.NoError(t, err)

		_, found := k.GetActionExecutor(sdkCtx, "schema", executor)
		require.False(t, found)
	})
}
