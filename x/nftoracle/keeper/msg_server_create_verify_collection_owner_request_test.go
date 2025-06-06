package keeper_test

import (
	"testing"

	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

func TestCollectionOwnerRequestMsgServerCreate(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateVerifyCollectionOwnerRequest(ctx, &types.MsgCreateVerifyCollectionOwnerRequest{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestCollectionOwnerRequestMsgServerSubmit(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.SubmitVerifyCollectionOwner(ctx, &types.MsgSubmitVerifyCollectionOwner{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.VerifyRequestID))
	}
}
