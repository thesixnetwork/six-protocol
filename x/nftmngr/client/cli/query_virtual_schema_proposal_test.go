package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/testutil/network"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/client/cli"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithVirtualSchemaProposalObjects(t *testing.T, n int) (*network.Network, []types.VirtualSchemaProposal) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		virtualSchemaProposal := types.VirtualSchemaProposal{
			Id: strconv.Itoa(i),
		}
		nullify.Fill(&virtualSchemaProposal)
		state.VirtualSchemaProposalList = append(state.VirtualSchemaProposalList, virtualSchemaProposal)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.VirtualSchemaProposalList
}

func TestShowVirtualSchemaProposal(t *testing.T) {
	net, objs := networkWithVirtualSchemaProposalObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc    string
		idIndex string

		args []string
		err  error
		obj  types.VirtualSchemaProposal
	}{
		{
			desc:    "found",
			idIndex: objs[0].Id,

			args: common,
			obj:  objs[0],
		},
		{
			desc:    "not found",
			idIndex: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idIndex,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowVirtualSchemaProposal(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetVirtualSchemaProposalResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.VirtualSchemaProposal)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.VirtualSchemaProposal),
				)
			}
		})
	}
}

func TestListVirtualSchemaProposal(t *testing.T) {
	net, objs := networkWithVirtualSchemaProposalObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListVirtualSchemaProposal(), args)
			require.NoError(t, err)
			var resp types.QueryAllVirtualSchemaProposalResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.VirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.VirtualSchemaProposal),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListVirtualSchemaProposal(), args)
			require.NoError(t, err)
			var resp types.QueryAllVirtualSchemaProposalResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.VirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.VirtualSchemaProposal),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListVirtualSchemaProposal(), args)
		require.NoError(t, err)
		var resp types.QueryAllVirtualSchemaProposalResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.VirtualSchemaProposal),
		)
	})
}

func networkWithDisableVirtualSchemaObjects(t *testing.T, n int) (*network.Network, []types.DisableVirtualSchemaProposal) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		disableVirtualSchema := types.DisableVirtualSchemaProposal{
			Id: strconv.Itoa(i),
		}
		nullify.Fill(&disableVirtualSchema)
		state.DisableVirtualSchemaProposalList = append(state.DisableVirtualSchemaProposalList, disableVirtualSchema)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.DisableVirtualSchemaProposalList
}

func TestShowDisableVirtualSchema(t *testing.T) {
	net, objs := networkWithDisableVirtualSchemaObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc    string
		idIndex string

		args []string
		err  error
		obj  types.DisableVirtualSchemaProposal
	}{
		{
			desc:    "found",
			idIndex: objs[0].Id,

			args: common,
			obj:  objs[0],
		},
		{
			desc:    "not found",
			idIndex: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idIndex,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowDisableVirtualSchemaProposal(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetDisableVirtualSchemaProposalResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.DisableVirtualSchemaProposal)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.DisableVirtualSchemaProposal),
				)
			}
		})
	}
}

func TestListDisableVirtualSchema(t *testing.T) {
	net, objs := networkWithDisableVirtualSchemaObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDisableVirtualSchemaProposal(), args)
			require.NoError(t, err)
			var resp types.QueryAllDisableVirtualSchemaProposalResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.DisableVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.DisableVirtualSchemaProposal),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDisableVirtualSchemaProposal(), args)
			require.NoError(t, err)
			var resp types.QueryAllDisableVirtualSchemaProposalResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.DisableVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.DisableVirtualSchemaProposal),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDisableVirtualSchemaProposal(), args)
		require.NoError(t, err)
		var resp types.QueryAllDisableVirtualSchemaProposalResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.DisableVirtualSchemaProposal),
		)
	})
}

func networkWithEnableVirtualSchemaProposalObjects(t *testing.T, n int) (*network.Network, []types.EnableVirtualSchemaProposal) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		enableVirtualSchemaProposal := types.EnableVirtualSchemaProposal{
			Id: strconv.Itoa(i),
		}
		nullify.Fill(&enableVirtualSchemaProposal)
		state.EnableVirtualSchemaProposalList = append(state.EnableVirtualSchemaProposalList, enableVirtualSchemaProposal)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.EnableVirtualSchemaProposalList
}

func TestShowEnableVirtualSchemaProposal(t *testing.T) {
	net, objs := networkWithEnableVirtualSchemaProposalObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc    string
		idIndex string

		args []string
		err  error
		obj  types.EnableVirtualSchemaProposal
	}{
		{
			desc:    "found",
			idIndex: objs[0].Id,

			args: common,
			obj:  objs[0],
		},
		{
			desc:    "not found",
			idIndex: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idIndex,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowEnableVirtualSchemaProposal(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetEnableVirtualSchemaProposalResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.EnableVirtualSchemaProposal)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.EnableVirtualSchemaProposal),
				)
			}
		})
	}
}

func TestListEnableVirtualSchemaProposal(t *testing.T) {
	net, objs := networkWithEnableVirtualSchemaProposalObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListEnableVirtualSchemaProposal(), args)
			require.NoError(t, err)
			var resp types.QueryAllEnableVirtualSchemaProposalResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.EnableVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.EnableVirtualSchemaProposal),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListEnableVirtualSchemaProposal(), args)
			require.NoError(t, err)
			var resp types.QueryAllEnableVirtualSchemaProposalResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.EnableVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.EnableVirtualSchemaProposal),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListEnableVirtualSchemaProposal(), args)
		require.NoError(t, err)
		var resp types.QueryAllEnableVirtualSchemaProposalResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.EnableVirtualSchemaProposal),
		)
	})
}
