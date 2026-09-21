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

// Only the permission and token existence error paths are exercised here: the
// mint path itself requires a bank keeper, which is not wired in the test keeper.
func TestMintMsgServer(t *testing.T) {
	creator := sample.AccAddress()

	t.Run("NoMintPermission", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.Mint(ctx, &types.MsgMint{
			Creator: creator,
			Amount:  sdk.NewCoin("faketoken", sdkmath.NewInt(100)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	})

	t.Run("TokenNotFound", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetMintperm(ctx, types.Mintperm{
			Creator: creator,
			Token:   "faketoken",
			Address: creator,
		})
		_, err := ms.Mint(ctx, &types.MsgMint{
			Creator: creator,
			Amount:  sdk.NewCoin("faketoken", sdkmath.NewInt(100)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
	})
}
