package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestVirtualSchemaMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgProposalVirtualSchema{
			Creator:      creator,
			ProposalType: types.ProposalType_CREATE,
			VirtualSchema: &types.VirtualSchema{
				VirtualNftSchemaCode: "virtualNftSchemaCode",
				Registry: []*types.VirtualSchemaRegistry{
					{
						NftSchemaCode:    "nftSchemaCode",
						SharedAttributes: []string{"SharedAttributes"},
						Decision:         types.RegistryStatus_PENDING,
					},
				},
			},
		}
		_, err := srv.ProposalVirtualSchema(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetVirtualSchema(ctx,
			expected.VirtualSchema.VirtualNftSchemaCode,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.VirtualNftSchemaCode)
	}
}
