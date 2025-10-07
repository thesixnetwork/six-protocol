package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/testutil/network"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/client/cli"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithBindedSignerObjects(t *testing.T, n int) (*network.Network, []types.BindedSigner) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		bindedSigner := types.BindedSigner{
			OwnerAddress: strconv.Itoa(i),
		}
		nullify.Fill(&bindedSigner)
		state.BindedSignerList = append(state.BindedSignerList, bindedSigner)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	nw, err := network.New(t, cfg)
	require.NoError(t, err)

	return nw, state.BindedSignerList
}

func TestShowBindedSigner(t *testing.T) {
	net, objs := networkWithBindedSignerObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc           string
		idOwnerAddress string

		args []string
		err  error
		obj  types.BindedSigner
	}{
		{
			desc:           "found",
			idOwnerAddress: objs[0].OwnerAddress,

			args: common,
			obj:  objs[0],
		},
		{
			desc:           "not found",
			idOwnerAddress: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idOwnerAddress,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowBindedSigner(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetBindedSignerResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.BindedSigner)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.BindedSigner),
				)
			}
		})
	}
}
