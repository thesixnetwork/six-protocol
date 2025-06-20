package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/testutil/network"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftadmin/client/cli"
	"github.com/thesixnetwork/six-protocol/x/nftadmin/types"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
)

func networkWithAuthorizationObjects(t *testing.T) (*network.Network, types.Authorization) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	authorization := &types.Authorization{}
	nullify.Fill(&authorization)
	state.Authorization = authorization
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	nw, err := network.New(t, cfg)
	require.NoError(t, err)

	return nw, *state.Authorization
}

func TestShowAuthorization(t *testing.T) {
	net, obj := networkWithAuthorizationObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.Authorization
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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowAuthorization(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetAuthorizationResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Authorization)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Authorization),
				)
			}
		})
	}
}
