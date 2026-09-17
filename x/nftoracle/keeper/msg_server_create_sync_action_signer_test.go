package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"
)

func TestCreateSyncActionSigner(t *testing.T) {
	creator := sample.AccAddress()
	chain := "ethereum"
	actorAddress := "0x0000000000000000000000000000000000000001"
	ownerAddress := "0x0000000000000000000000000000000000000002"

	t.Run("oracle config not found", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.CreateSyncActionSigner(ctx, types.NewMsgCreateSyncActionSigner(creator, chain, actorAddress, ownerAddress, 1))
		require.ErrorIs(t, err, types.ErrOracleConfigNotFound)
	})

	t.Run("required confirm too less", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetOracleConfig(ctx, types.OracleConfig{MinimumConfirmation: 2})
		_, err := ms.CreateSyncActionSigner(ctx, types.NewMsgCreateSyncActionSigner(creator, chain, actorAddress, ownerAddress, 1))
		require.ErrorIs(t, err, types.ErrRequiredConfirmTooLess)
	})

	t.Run("chain config not found", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetOracleConfig(ctx, types.OracleConfig{MinimumConfirmation: 1})
		_, err := ms.CreateSyncActionSigner(ctx, types.NewMsgCreateSyncActionSigner(creator, chain, actorAddress, ownerAddress, 1))
		require.ErrorIs(t, err, types.ErrActionSignerConfigNotFound)
	})

	t.Run("success", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetOracleConfig(ctx, types.OracleConfig{MinimumConfirmation: 1})
		k.SetActionSignerConfig(ctx, types.ActionSignerConfig{
			Chain:           chain,
			Creator:         creator,
			ContractAddress: "0x0000000000000000000000000000000000000003",
		})

		resp, err := ms.CreateSyncActionSigner(ctx, types.NewMsgCreateSyncActionSigner(creator, chain, actorAddress, ownerAddress, 2))
		require.NoError(t, err)
		require.NotNil(t, resp)

		require.Equal(t, uint64(1), k.GetSyncActionSignerCount(ctx))
		syncSigner, found := k.GetSyncActionSigner(ctx, 0)
		require.True(t, found)
		require.Equal(t, chain, syncSigner.Chain)
		require.Equal(t, actorAddress, syncSigner.ActorAddress)
		require.Equal(t, ownerAddress, syncSigner.OwnerAddress)
		require.Equal(t, creator, syncSigner.Caller)
		require.Equal(t, uint64(2), syncSigner.RequiredConfirm)
		require.Equal(t, types.RequestStatus_PENDING, syncSigner.Status)
		require.Equal(t, uint64(0), syncSigner.CurrentConfirm)
	})
}
