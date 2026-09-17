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

// Only the pre-bank error paths are exercised here: the burn path itself
// requires a bank keeper, which is not wired in the test keeper.
func TestBurnMsgServer(t *testing.T) {
	t.Run("InvalidCreator", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.Burn(ctx, &types.MsgBurn{
			Creator: "invalid_address",
			Amount:  sdk.NewCoin("usix", sdkmath.NewInt(100)),
		})
		require.Error(t, err)
	})

	t.Run("ZeroAmount", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.Burn(ctx, &types.MsgBurn{
			Creator: sample.AccAddress(),
			Amount:  sdk.NewCoin("usix", sdkmath.NewInt(0)),
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	})
}
