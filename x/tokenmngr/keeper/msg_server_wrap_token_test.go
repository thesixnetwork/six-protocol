package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Only the pre-bank error paths are exercised here: the wrap path itself
// requires a bank keeper, which is not wired in the test keeper.
func TestWrapTokenMsgServer(t *testing.T) {
	creator := sample.AccAddress()
	receiver := sample.AccAddress()

	t.Run("InvalidCreator", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.WrapToken(ctx, &types.MsgWrapToken{
			Creator:  "invalid_address",
			Receiver: receiver,
			Amount:   sdk.NewCoin("usix", sdkmath.NewInt(100)),
		})
		require.Error(t, err)
	})

	t.Run("InvalidReceiver", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.WrapToken(ctx, &types.MsgWrapToken{
			Creator:  creator,
			Receiver: "invalid_address",
			Amount:   sdk.NewCoin("usix", sdkmath.NewInt(100)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidAddress)
	})

	t.Run("WrongDenom", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.WrapToken(ctx, &types.MsgWrapToken{
			Creator:  creator,
			Receiver: receiver,
			Amount:   sdk.NewCoin("foo", sdkmath.NewInt(100)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	})

	t.Run("TokenNotFound", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.WrapToken(ctx, &types.MsgWrapToken{
			Creator:  creator,
			Receiver: receiver,
			Amount:   sdk.NewCoin("usix", sdkmath.NewInt(100)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
	})

	t.Run("TokenBaseNotUsix", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetToken(ctx, types.Token{Creator: creator, Name: "usix", Base: "foo"})
		_, err := ms.WrapToken(ctx, &types.MsgWrapToken{
			Creator:  creator,
			Receiver: receiver,
			Amount:   sdk.NewCoin("usix", sdkmath.NewInt(100)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	})

	t.Run("ZeroAmount", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetToken(ctx, types.Token{Creator: creator, Name: "usix", Base: "usix"})
		_, err := ms.WrapToken(ctx, &types.MsgWrapToken{
			Creator:  creator,
			Receiver: receiver,
			Amount:   sdk.NewCoin("usix", sdkmath.NewInt(0)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	})
}
