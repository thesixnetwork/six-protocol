package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/types"
)

func TestListAdminOfGroupQuery(t *testing.T) {
	k, ctx := keepertest.ProtocoladminKeeper(t)

	group1Admins := []string{sample.AccAddress(), sample.AccAddress()}
	for _, admin := range group1Admins {
		k.SetAdmin(ctx, types.Admin{Group: "group1", Admin: admin})
	}
	otherAdmin := sample.AccAddress()
	k.SetAdmin(ctx, types.Admin{Group: "group2", Admin: otherAdmin})

	t.Run("OnlyRequestedGroup", func(t *testing.T) {
		resp, err := k.ListAdminOfGroup(ctx, &types.QueryListAdminOfGroupRequest{Group: "group1"})
		require.NoError(t, err)
		require.ElementsMatch(t, group1Admins, resp.Admin)
	})
	t.Run("EmptyGroup", func(t *testing.T) {
		resp, err := k.ListAdminOfGroup(ctx, &types.QueryListAdminOfGroupRequest{Group: "unknown"})
		require.NoError(t, err)
		require.Empty(t, resp.Admin)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.ListAdminOfGroup(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
