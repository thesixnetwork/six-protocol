package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// UpdateToken is disabled on chain and always returns an error.
func TestTokenMsgServerUpdate(t *testing.T) {
	_, ms, ctx := setupMsgServer(t)

	_, err := ms.UpdateToken(ctx, &types.MsgUpdateToken{
		Creator: sample.AccAddress(),
		Name:    "anything",
	})
	require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	require.Contains(t, err.Error(), "operation not available")
}

func TestTokenMsgServerDelete(t *testing.T) {
	creator := sample.AccAddress()

	t.Run("KeyNotFound", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.DeleteToken(ctx, &types.MsgDeleteToken{
			Creator: creator,
			Name:    "missing",
		})
		require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetToken(ctx, types.Token{Creator: creator, Name: "mytoken"})
		_, err := ms.DeleteToken(ctx, &types.MsgDeleteToken{
			Creator: sample.AccAddress(),
			Name:    "mytoken",
		})
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	})

	t.Run("Completed", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetToken(ctx, types.Token{Creator: creator, Name: "mytoken"})
		_, err := ms.DeleteToken(ctx, &types.MsgDeleteToken{
			Creator: creator,
			Name:    "mytoken",
		})
		require.NoError(t, err)
		_, found := k.GetToken(ctx, "mytoken")
		require.False(t, found)
	})
}
