package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

func TestMetadataCreatorQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMetadataCreator(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMetadataCreatorRequest
		response *types.QueryGetMetadataCreatorResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMetadataCreatorRequest{
				NftSchemaCode: msgs[0].NftSchemaCode,
			},
			response: &types.QueryGetMetadataCreatorResponse{MetadataCreator: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMetadataCreatorRequest{
				NftSchemaCode: msgs[1].NftSchemaCode,
			},
			response: &types.QueryGetMetadataCreatorResponse{MetadataCreator: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMetadataCreatorRequest{
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
			response, err := keeper.MetadataCreator(wctx, tc.request)
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

func TestMetadataCreatorQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMetadataCreator(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMetadataCreatorRequest {
		return &types.QueryAllMetadataCreatorRequest{
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
			resp, err := keeper.MetadataCreatorAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MetadataCreator), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MetadataCreator),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MetadataCreatorAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MetadataCreator), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MetadataCreator),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MetadataCreatorAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.MetadataCreator),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MetadataCreatorAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestNftCollectionQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNftCollection(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNftCollectionRequest
		response *types.QueryGetNftCollectionResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetNftCollectionRequest{
				NftSchemaCode: msgs[0].NftSchemaCode,
			},
			response: &types.QueryGetNftCollectionResponse{NftCollection: msgs[0].NftDatas},
		},
		{
			desc: "Second",
			request: &types.QueryGetNftCollectionRequest{
				NftSchemaCode: msgs[1].NftSchemaCode,
			},
			response: &types.QueryGetNftCollectionResponse{NftCollection: msgs[1].NftDatas},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetNftCollectionRequest{
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
			response, err := keeper.NftCollection(wctx, tc.request)
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

func TestNftCollectionQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNftCollection(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryGetNftCollectionRequest {
		return &types.QueryGetNftCollectionRequest{
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
			resp, err := keeper.NftCollection(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NftCollection), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NftCollection),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NftCollection(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NftCollection), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NftCollection),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.NftCollection(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.NftCollection),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.NftCollection(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestNftDataQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNftData(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNftDataRequest
		response *types.QueryGetNftDataResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetNftDataRequest{
				NftSchemaCode: msgs[0].NftSchemaCode,
				TokenId:       msgs[0].TokenId,
			},
			response: &types.QueryGetNftDataResponse{NftData: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetNftDataRequest{
				NftSchemaCode: msgs[1].NftSchemaCode,
				TokenId:       msgs[1].TokenId,
			},
			response: &types.QueryGetNftDataResponse{NftData: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetNftDataRequest{
				NftSchemaCode: strconv.Itoa(100000),
				TokenId:       strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NftData(wctx, tc.request)
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

func TestNftDataQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNftData(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllNftDataRequest {
		return &types.QueryAllNftDataRequest{
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
			resp, err := keeper.NftDataAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NftData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NftData),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NftDataAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NftData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NftData),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.NftDataAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.NftData),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.NftDataAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func createNNFTSchemaByContract(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NFTSchemaByContract {
	items := make([]types.NFTSchemaByContract, n)
	for i := range items {
		items[i].OriginContractAddress = strconv.Itoa(i)

		keeper.SetNFTSchemaByContract(ctx, items[i])
	}
	return items
}

func TestNFTSchemaByContractGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchemaByContract(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNFTSchemaByContract(ctx,
			item.OriginContractAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestNFTSchemaByContractRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchemaByContract(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNFTSchemaByContract(ctx,
			item.OriginContractAddress,
		)
		_, found := keeper.GetNFTSchemaByContract(ctx,
			item.OriginContractAddress,
		)
		require.False(t, found)
	}
}

func TestNFTSchemaByContractGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchemaByContract(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNFTSchemaByContract(ctx)),
	)
}
