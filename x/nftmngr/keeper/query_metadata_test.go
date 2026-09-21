package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/types/query"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

func TestMetadataCreatorQuerySingle(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNMetadataCreator(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetMetadataCreatorRequest
		response *types.QueryGetMetadataCreatorResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetMetadataCreatorRequest{NftSchemaCode: msgs[0].NftSchemaCode},
			response: &types.QueryGetMetadataCreatorResponse{MetadataCreator: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetMetadataCreatorRequest{NftSchemaCode: msgs[1].NftSchemaCode},
			response: &types.QueryGetMetadataCreatorResponse{MetadataCreator: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetMetadataCreatorRequest{NftSchemaCode: strconv.Itoa(100000)},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.MetadataCreator(ctx, tc.request)
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
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNMetadataCreator(k, ctx, 5)

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
			resp, err := k.MetadataCreatorAll(ctx, request(nil, uint64(i), uint64(step), false))
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
			resp, err := k.MetadataCreatorAll(ctx, request(next, 0, uint64(step), false))
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
		resp, err := k.MetadataCreatorAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.MetadataCreator),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.MetadataCreatorAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestNftCollectionQuery(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)

	_, err := k.NftCollection(ctx, nil)
	require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))

	nftDatas := createNNftData(k, ctx, 3)
	for i := range nftDatas {
		nftDatas[i].NftSchemaCode = "collection"
		k.AddMetadataToCollection(ctx, &nftDatas[i])
	}

	resp, err := k.NftCollection(ctx, &types.QueryGetNftCollectionRequest{
		NftSchemaCode: "collection",
	})
	require.NoError(t, err)
	require.Len(t, resp.NftCollection, len(nftDatas))

	actual := make([]types.NftData, len(resp.NftCollection))
	for i, data := range resp.NftCollection {
		actual[i] = *data
	}
	require.ElementsMatch(t,
		nullify.Fill(nftDatas),
		nullify.Fill(actual),
	)
}
