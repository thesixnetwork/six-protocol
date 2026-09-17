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

const testEthAddress = "0x966fbC80c384355F08a168a15A4E7A8535586eC5"

// Only the pre-bank error paths are exercised here: the send path itself
// requires a bank keeper, which is not wired in the test keeper.
func TestSendWrapTokenMsgServer(t *testing.T) {
	creator := sample.AccAddress()

	t.Run("InvalidEthAddress", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.SendWrapToken(ctx, &types.MsgSendWrapToken{
			Creator:    creator,
			EthAddress: "invalid_address",
			Amount:     sdk.NewCoin("asix", sdkmath.NewInt(100)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidAddress)
	})

	t.Run("TokenNotFound", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.SendWrapToken(ctx, &types.MsgSendWrapToken{
			Creator:    creator,
			EthAddress: testEthAddress,
			Amount:     sdk.NewCoin("asix", sdkmath.NewInt(100)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
	})

	t.Run("TokenBaseNotAsix", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetToken(ctx, types.Token{Creator: creator, Name: "wsix", Base: "usix"})
		_, err := ms.SendWrapToken(ctx, &types.MsgSendWrapToken{
			Creator:    creator,
			EthAddress: testEthAddress,
			Amount:     sdk.NewCoin("wsix", sdkmath.NewInt(100)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	})

	t.Run("ZeroAmount", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetToken(ctx, types.Token{Creator: creator, Name: "wsix", Base: "asix"})
		_, err := ms.SendWrapToken(ctx, &types.MsgSendWrapToken{
			Creator:    creator,
			EthAddress: testEthAddress,
			Amount:     sdk.NewCoin("wsix", sdkmath.NewInt(0)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	})

	t.Run("InvalidCreator", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetToken(ctx, types.Token{Creator: creator, Name: "wsix", Base: "asix"})
		_, err := ms.SendWrapToken(ctx, &types.MsgSendWrapToken{
			Creator:    "invalid_address",
			EthAddress: testEthAddress,
			Amount:     sdk.NewCoin("wsix", sdkmath.NewInt(100)),
		})
		require.Error(t, err)
	})
}
