package keeper_test

import (
	"encoding/base64"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

func TestMsgServerCreateNFTSchema(t *testing.T) {
	creator := sample.AccAddress()

	t.Run("invalid creator address", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.CreateNFTSchema(ctx, &types.MsgCreateNFTSchema{
			Creator:         "invalid_address",
			NftSchemaBase64: "",
		})
		require.ErrorIs(t, err, sdkerrors.ErrInvalidAddress)
	})

	t.Run("invalid base64", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.CreateNFTSchema(ctx, &types.MsgCreateNFTSchema{
			Creator:         creator,
			NftSchemaBase64: "not-valid-base64!!!",
		})
		require.ErrorIs(t, err, types.ErrParsingBase64)
	})

	t.Run("invalid schema json", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.CreateNFTSchema(ctx, &types.MsgCreateNFTSchema{
			Creator:         creator,
			NftSchemaBase64: base64.StdEncoding.EncodeToString([]byte("not-json")),
		})
		require.ErrorIs(t, err, types.ErrParsingSchemaMessage)
	})
}
