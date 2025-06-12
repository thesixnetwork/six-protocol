package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/testutil/network"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/client/cli"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithSchemaAttributeObjects(t *testing.T, n int) (*network.Network, []types.SchemaAttribute) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		schemaAttribute := types.SchemaAttribute{
			NftSchemaCode: strconv.Itoa(i),
			Name:          strconv.Itoa(i),
		}
		nullify.Fill(&schemaAttribute)
		state.SchemaAttributeList = append(state.SchemaAttributeList, schemaAttribute)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	nw, err := network.New(t, cfg)
	require.NoError(t, err)

	return nw, state.SchemaAttributeList
}

func TestShowSchemaAttribute(t *testing.T) {
	net, objs := networkWithSchemaAttributeObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc            string
		idNftSchemaCode string
		idName          string

		args []string
		err  error
		obj  types.SchemaAttribute
	}{
		{
			desc:            "found",
			idNftSchemaCode: objs[0].NftSchemaCode,
			idName:          objs[0].Name,

			args: common,
			obj:  objs[0],
		},
		{
			desc:            "not found",
			idNftSchemaCode: strconv.Itoa(100000),
			idName:          strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idNftSchemaCode,
				tc.idName,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowSchemaAttribute(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetSchemaAttributeResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.SchemaAttribute)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.SchemaAttribute),
				)
			}
		})
	}
}

func TestListSchemaAttribute(t *testing.T) {
	net, objs := networkWithSchemaAttributeObjects(t, 5)

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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListSchemaAttribute(), args)
			require.NoError(t, err)
			var resp types.QueryAllSchemaAttributeResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.SchemaAttribute), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.SchemaAttribute),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListSchemaAttribute(), args)
			require.NoError(t, err)
			var resp types.QueryAllSchemaAttributeResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.SchemaAttribute), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.SchemaAttribute),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListSchemaAttribute(), args)
		require.NoError(t, err)
		var resp types.QueryAllSchemaAttributeResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.SchemaAttribute),
		)
	})
}
