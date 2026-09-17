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

const attoToMicro = int64(1_000_000_000_000)

// Only the pre-bank error paths are exercised here: the unwrap path itself
// requires bank and account keepers, which are not wired in the test keeper.
func TestUnwrapTokenMsgServer(t *testing.T) {
	creator := sample.AccAddress()
	receiver := sample.AccAddress()

	t.Run("ZeroAmount", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.UnwrapToken(ctx, &types.MsgUnwrapToken{
			Creator:  creator,
			Receiver: receiver,
			Amount:   sdk.NewCoin("asix", sdkmath.NewInt(0)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	})

	t.Run("NonAttoMultipleAmount", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.UnwrapToken(ctx, &types.MsgUnwrapToken{
			Creator:  creator,
			Receiver: receiver,
			Amount:   sdk.NewCoin("asix", sdkmath.NewInt(5)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	})

	t.Run("InvalidCreator", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.UnwrapToken(ctx, &types.MsgUnwrapToken{
			Creator:  "invalid_address",
			Receiver: receiver,
			Amount:   sdk.NewCoin("asix", sdkmath.NewInt(attoToMicro)),
		})
		require.Error(t, err)
	})

	t.Run("InvalidReceiver", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.UnwrapToken(ctx, &types.MsgUnwrapToken{
			Creator:  creator,
			Receiver: "invalid_address",
			Amount:   sdk.NewCoin("asix", sdkmath.NewInt(attoToMicro)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidAddress)
	})

	t.Run("WrongDenom", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.UnwrapToken(ctx, &types.MsgUnwrapToken{
			Creator:  creator,
			Receiver: receiver,
			Amount:   sdk.NewCoin("usix", sdkmath.NewInt(attoToMicro)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	})

	t.Run("TokenNotFound", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.UnwrapToken(ctx, &types.MsgUnwrapToken{
			Creator:  creator,
			Receiver: receiver,
			Amount:   sdk.NewCoin("asix", sdkmath.NewInt(attoToMicro)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
	})

	t.Run("TokenBaseNotAsix", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetToken(ctx, types.Token{Creator: creator, Name: "asix", Base: "usix"})
		_, err := ms.UnwrapToken(ctx, &types.MsgUnwrapToken{
			Creator:  creator,
			Receiver: receiver,
			Amount:   sdk.NewCoin("asix", sdkmath.NewInt(attoToMicro)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	})
}
