package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"
)

// MigrateDelegation is currently a no-op handler that always succeeds.
func TestMigrateDelegationMsgServer(t *testing.T) {
	_, ms, ctx := setupMsgServer(t)

	resp, err := ms.MigrateDelegation(ctx, &types.MsgMigrateDelegation{
		Creator: sample.AccAddress(),
	})
	require.NoError(t, err)
	require.Equal(t, &types.MsgMigrateDelegationResponse{}, resp)
}
