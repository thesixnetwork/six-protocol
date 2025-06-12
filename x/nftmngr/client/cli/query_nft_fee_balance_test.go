package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/testutil/network"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/client/cli"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
)

func networkWithNFTFeeBalanceObjects(t *testing.T) (*network.Network, types.NFTFeeBalance) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	nFTFeeBalance := &types.NFTFeeBalance{}
	nullify.Fill(&nFTFeeBalance)
	state.NFTFeeBalance = nFTFeeBalance
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	nw, err := network.New(t, cfg)
	require.NoError(t, err)

	return nw,  *state.NFTFeeBalance
}

func TestShowNFTFeeBalance(t *testing.T) {
	net, obj := networkWithNFTFeeBalanceObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.NFTFeeBalance
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowNFTFeeBalance(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetNFTFeeBalanceResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.NFTFeeBalance)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.NFTFeeBalance),
				)
			}
		})
	}
}
