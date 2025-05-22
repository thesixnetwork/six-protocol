package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestActionExecutorQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActionExecutor(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetActionExecutorRequest
		response *types.QueryGetActionExecutorResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetActionExecutorRequest{
				NftSchemaCode:   msgs[0].NftSchemaCode,
				ExecutorAddress: msgs[0].ExecutorAddress,
			},
			response: &types.QueryGetActionExecutorResponse{ActionExecutor: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetActionExecutorRequest{
				NftSchemaCode:   msgs[1].NftSchemaCode,
				ExecutorAddress: msgs[1].ExecutorAddress,
			},
			response: &types.QueryGetActionExecutorResponse{ActionExecutor: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetActionExecutorRequest{
				NftSchemaCode:   strconv.Itoa(100000),
				ExecutorAddress: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ActionExecutor(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestActionExecutorQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActionExecutor(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActionExecutorRequest {
		return &types.QueryAllActionExecutorRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActionExecutorAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionExecutor), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionExecutor),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActionExecutorAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionExecutor), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionExecutor),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ActionExecutorAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ActionExecutor),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ActionExecutorAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestExecutorOfSchemaQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNExecutorOfSchema(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetExecutorOfSchemaRequest
		response *types.QueryGetExecutorOfSchemaResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetExecutorOfSchemaRequest{
				NftSchemaCode: msgs[0].NftSchemaCode,
			},
			response: &types.QueryGetExecutorOfSchemaResponse{ExecutorOfSchema: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetExecutorOfSchemaRequest{
				NftSchemaCode: msgs[1].NftSchemaCode,
			},
			response: &types.QueryGetExecutorOfSchemaResponse{ExecutorOfSchema: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetExecutorOfSchemaRequest{
				NftSchemaCode: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ExecutorOfSchema(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestExecutorOfSchemaQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNExecutorOfSchema(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllExecutorOfSchemaRequest {
		return &types.QueryAllExecutorOfSchemaRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ExecutorOfSchemaAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ExecutorOfSchema), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ExecutorOfSchema),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ExecutorOfSchemaAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ExecutorOfSchema), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ExecutorOfSchema),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ExecutorOfSchemaAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ExecutorOfSchema),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ExecutorOfSchemaAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestActionExecutorMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateActionExecutor{
			Creator:         creator,
			NftSchemaCode:   strconv.Itoa(i),
			ExecutorAddress: strconv.Itoa(i),
		}
		_, err := srv.CreateActionExecutor(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetActionExecutor(ctx,
			expected.NftSchemaCode,
			expected.ExecutorAddress,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestActionExecutorMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateActionExecutor
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateActionExecutor{
				Creator:         creator,
				NftSchemaCode:   strconv.Itoa(0),
				ExecutorAddress: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateActionExecutor{
				Creator:         "B",
				NftSchemaCode:   strconv.Itoa(0),
				ExecutorAddress: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateActionExecutor{
				Creator:         creator,
				NftSchemaCode:   strconv.Itoa(100000),
				ExecutorAddress: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NftmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateActionExecutor{
				Creator:         creator,
				NftSchemaCode:   strconv.Itoa(0),
				ExecutorAddress: strconv.Itoa(0),
			}
			_, err := srv.CreateActionExecutor(wctx, expected)
			require.NoError(t, err)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetActionExecutor(ctx,
					expected.NftSchemaCode,
					expected.ExecutorAddress,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestActionExecutorMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteActionExecutor
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteActionExecutor{
				Creator:         creator,
				NftSchemaCode:   strconv.Itoa(0),
				ExecutorAddress: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteActionExecutor{
				Creator:         "B",
				NftSchemaCode:   strconv.Itoa(0),
				ExecutorAddress: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteActionExecutor{
				Creator:         creator,
				NftSchemaCode:   strconv.Itoa(100000),
				ExecutorAddress: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NftmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateActionExecutor(wctx, &types.MsgCreateActionExecutor{
				Creator:         creator,
				NftSchemaCode:   strconv.Itoa(0),
				ExecutorAddress: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteActionExecutor(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetActionExecutor(ctx,
					tc.request.NftSchemaCode,
					tc.request.ExecutorAddress,
				)
				require.False(t, found)
			}
		})
	}
}
