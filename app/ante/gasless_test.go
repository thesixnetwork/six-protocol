package ante_test

import (
	"math"
	"testing"
	"time"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"

	"github.com/thesixnetwork/six-protocol/app/ante"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	nftadmintypes "github.com/thesixnetwork/six-protocol/x/nftadmin/types"
	nftoracletypes "github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

// ─────────────────────────────────────────────────────────────────────────────
// Test suite setup
// ─────────────────────────────────────────────────────────────────────────────

type GaslessTestSuite struct {
	suite.Suite
	keepers            keepertest.CombinedKeepers
	ctx                sdk.Context
	gaslessDecorator   ante.GaslessDecorator
	voteAloneDecorator ante.VoteAloneDecorator
	oracleAddr         sdk.AccAddress
	nonOracleAddr      sdk.AccAddress
}

func TestGaslessTestSuite(t *testing.T) {
	suite.Run(t, new(GaslessTestSuite))
}

func (suite *GaslessTestSuite) SetupTest() {
	suite.keepers = keepertest.NewCombinedKeepers(suite.T())
	suite.ctx = suite.keepers.Ctx

	suite.oracleAddr = sdk.AccAddress("oracle_address_test_")
	suite.nonOracleAddr = sdk.AccAddress("non_oracle_addr_____")

	suite.grantOraclePermission(suite.oracleAddr)

	wrappedDecorators := []sdk.AnteDecorator{MockFeeDecorator{}}
	suite.gaslessDecorator = ante.NewGaslessDecorator(
		wrappedDecorators,
		suite.keepers.OracleKeeper,
		suite.keepers.AdminKeeper,
	)
	suite.voteAloneDecorator = ante.NewVoteAloneDecorator()
}

// grantOraclePermission sets the oracle permission for the given address.
func (suite *GaslessTestSuite) grantOraclePermission(addr sdk.AccAddress) {
	auth := nftadmintypes.Authorization{
		RootAdmin: "cosmos1rootadmin",
		Permissions: []*nftadmintypes.Permission{
			{
				Name:      nftoracletypes.KeyPermissionOracle,
				Addresses: []string{addr.String()},
			},
		},
	}
	suite.keepers.AdminKeeper.SetAuthorization(suite.ctx, auth)
}

// ─────────────────────────────────────────────────────────────────────────────
// Mock helpers
// ─────────────────────────────────────────────────────────────────────────────

// MockFeeDecorator is a no-op fee decorator.
type MockFeeDecorator struct{}

func (mfd MockFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	return next(ctx, tx, simulate)
}

// FailingDecorator always returns an error when called.
type FailingDecorator struct{}

func (fd *FailingDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	return ctx, errortypes.ErrInvalidRequest.Wrap("failing decorator executed")
}

// mockTx implements sdk.Tx for testing.
type mockTx struct{ msgs []sdk.Msg }

func (tx *mockTx) GetMsgs() []sdk.Msg                  { return tx.msgs }
func (tx *mockTx) GetMsgsV2() ([]proto.Message, error) { return []proto.Message{}, nil }
func (tx *mockTx) ValidateBasic() error                { return nil }

func (suite *GaslessTestSuite) createMockTx(msgs ...sdk.Msg) sdk.Tx {
	return &mockTx{msgs: msgs}
}

// passthroughHandler is a terminal AnteHandler that simply returns its context.
func passthroughHandler(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) { return ctx, nil }

// ─────────────────────────────────────────────────────────────────────────────
// Helpers that seed the oracle KV store
// ─────────────────────────────────────────────────────────────────────────────

func (suite *GaslessTestSuite) newPendingMintRequest(id uint64, confirmers ...string) {
	suite.keepers.OracleKeeper.SetMintRequest(suite.ctx, nftoracletypes.MintRequest{
		Id:              id,
		NftSchemaCode:   "schema.test",
		TokenId:         "token1",
		RequiredConfirm: 1,
		Status:          nftoracletypes.RequestStatus_PENDING,
		Confirmers:      confirmers,
		CreatedAt:       time.Now(),
		ValidUntil:      time.Now().Add(time.Hour),
	})
}

func (suite *GaslessTestSuite) newPendingActionRequest(id uint64, confirmers ...string) {
	suite.keepers.OracleKeeper.SetActionRequest(suite.ctx, nftoracletypes.ActionOracleRequest{
		Id:              id,
		NftSchemaCode:   "schema.test",
		TokenId:         "token1",
		Action:          "testAction",
		RequiredConfirm: 1,
		Status:          nftoracletypes.RequestStatus_PENDING,
		Confirmers:      confirmers,
		CreatedAt:       time.Now(),
		ValidUntil:      time.Now().Add(time.Hour),
	})
}

func (suite *GaslessTestSuite) newPendingCollectionRequest(id uint64, confirmers ...string) {
	suite.keepers.OracleKeeper.SetCollectionOwnerRequest(suite.ctx, nftoracletypes.CollectionOwnerRequest{
		Id:              id,
		NftSchemaCode:   "schema.test",
		RequiredConfirm: 1,
		Status:          nftoracletypes.RequestStatus_PENDING,
		Confirmers:      confirmers,
		CreatedAt:       time.Now(),
		ValidUntil:      time.Now().Add(time.Hour),
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// IsTxGasless – basic edge cases
// ─────────────────────────────────────────────────────────────────────────────

func (suite *GaslessTestSuite) TestIsTxGasless_EmptyTx() {
	tx := suite.createMockTx()
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().NoError(err)
	suite.Require().False(isGasless)
}

func (suite *GaslessTestSuite) TestIsTxGasless_NonOracleMsg() {
	tx := suite.createMockTx(&banktypes.MsgSend{})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().NoError(err)
	suite.Require().False(isGasless)
}

func (suite *GaslessTestSuite) TestIsTxGasless_BundledOracleMsgs_NotGasless() {
	tx := suite.createMockTx(
		&nftoracletypes.MsgSubmitMintResponse{Creator: suite.oracleAddr.String(), MintRequestID: 1},
		&banktypes.MsgSend{},
	)
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().NoError(err)
	suite.Require().False(isGasless, "bundled oracle tx must not be gasless")
}

// ─────────────────────────────────────────────────────────────────────────────
// IsTxGasless – MintResponse
// ─────────────────────────────────────────────────────────────────────────────

func (suite *GaslessTestSuite) TestIsTxGasless_MintResponse_HappyPath() {
	suite.newPendingMintRequest(1)
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitMintResponse{
		Creator:       suite.oracleAddr.String(),
		MintRequestID: 1,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().NoError(err)
	suite.Require().True(isGasless, "valid oracle vote must be gasless")
}

func (suite *GaslessTestSuite) TestIsTxGasless_MintResponse_NoPermission() {
	suite.newPendingMintRequest(1)
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitMintResponse{
		Creator:       suite.nonOracleAddr.String(),
		MintRequestID: 1,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().Error(err)
	suite.Require().False(isGasless)
	suite.Require().ErrorContains(err, nftoracletypes.ErrNoOraclePermission.Error())
}

func (suite *GaslessTestSuite) TestIsTxGasless_MintResponse_RequestNotFound() {
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitMintResponse{
		Creator:       suite.oracleAddr.String(),
		MintRequestID: 999,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().Error(err)
	suite.Require().False(isGasless)
	suite.Require().ErrorContains(err, nftoracletypes.ErrMintRequestNotFound.Error())
}

func (suite *GaslessTestSuite) TestIsTxGasless_MintResponse_NotPending() {
	suite.keepers.OracleKeeper.SetMintRequest(suite.ctx, nftoracletypes.MintRequest{
		Id:     1,
		Status: nftoracletypes.RequestStatus_SUCCESS_WITH_CONSENSUS,
	})
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitMintResponse{
		Creator:       suite.oracleAddr.String(),
		MintRequestID: 1,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().Error(err)
	suite.Require().False(isGasless)
	suite.Require().ErrorContains(err, nftoracletypes.ErrMintRequestNotPending.Error())
}

func (suite *GaslessTestSuite) TestIsTxGasless_MintResponse_DuplicateVote() {
	suite.newPendingMintRequest(1, suite.oracleAddr.String())
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitMintResponse{
		Creator:       suite.oracleAddr.String(),
		MintRequestID: 1,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().Error(err)
	suite.Require().False(isGasless)
	suite.Require().ErrorContains(err, nftoracletypes.ErrOracleAlreadyVoted.Error())
}

// A different oracle already voted – the current oracle should still be allowed.
func (suite *GaslessTestSuite) TestIsTxGasless_MintResponse_OtherOracleAlreadyVoted_Allowed() {
	otherOracle := sdk.AccAddress("other_oracle________")
	suite.newPendingMintRequest(1, otherOracle.String())
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitMintResponse{
		Creator:       suite.oracleAddr.String(),
		MintRequestID: 1,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().NoError(err)
	suite.Require().True(isGasless, "oracle may still vote even if other oracles have voted")
}

// ─────────────────────────────────────────────────────────────────────────────
// IsTxGasless – ActionResponse
// ─────────────────────────────────────────────────────────────────────────────

func (suite *GaslessTestSuite) TestIsTxGasless_ActionResponse_HappyPath() {
	suite.newPendingActionRequest(1)
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitActionResponse{
		Creator:         suite.oracleAddr.String(),
		ActionRequestID: 1,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().NoError(err)
	suite.Require().True(isGasless)
}

func (suite *GaslessTestSuite) TestIsTxGasless_ActionResponse_NoPermission() {
	suite.newPendingActionRequest(1)
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitActionResponse{
		Creator:         suite.nonOracleAddr.String(),
		ActionRequestID: 1,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().Error(err)
	suite.Require().False(isGasless)
}

func (suite *GaslessTestSuite) TestIsTxGasless_ActionResponse_NotFound() {
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitActionResponse{
		Creator:         suite.oracleAddr.String(),
		ActionRequestID: 999,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().Error(err)
	suite.Require().False(isGasless)
	suite.Require().ErrorContains(err, nftoracletypes.ErrActionRequestNotFound.Error())
}

func (suite *GaslessTestSuite) TestIsTxGasless_ActionResponse_DuplicateVote() {
	suite.newPendingActionRequest(1, suite.oracleAddr.String())
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitActionResponse{
		Creator:         suite.oracleAddr.String(),
		ActionRequestID: 1,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().Error(err)
	suite.Require().False(isGasless)
	suite.Require().ErrorContains(err, nftoracletypes.ErrOracleAlreadyVoted.Error())
}

// ─────────────────────────────────────────────────────────────────────────────
// IsTxGasless – CollectionVerify
// ─────────────────────────────────────────────────────────────────────────────

func (suite *GaslessTestSuite) TestIsTxGasless_CollectionVerify_HappyPath() {
	suite.newPendingCollectionRequest(1)
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitVerifyCollectionOwner{
		Creator:         suite.oracleAddr.String(),
		VerifyRequestID: 1,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().NoError(err)
	suite.Require().True(isGasless)
}

func (suite *GaslessTestSuite) TestIsTxGasless_CollectionVerify_NoPermission() {
	suite.newPendingCollectionRequest(1)
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitVerifyCollectionOwner{
		Creator:         suite.nonOracleAddr.String(),
		VerifyRequestID: 1,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().Error(err)
	suite.Require().False(isGasless)
}

func (suite *GaslessTestSuite) TestIsTxGasless_CollectionVerify_NotFound() {
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitVerifyCollectionOwner{
		Creator:         suite.oracleAddr.String(),
		VerifyRequestID: 999,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().Error(err)
	suite.Require().False(isGasless)
	suite.Require().ErrorContains(err, nftoracletypes.ErrVerifyRequestNotFound.Error())
}

func (suite *GaslessTestSuite) TestIsTxGasless_CollectionVerify_DuplicateVote() {
	suite.newPendingCollectionRequest(1, suite.oracleAddr.String())
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitVerifyCollectionOwner{
		Creator:         suite.oracleAddr.String(),
		VerifyRequestID: 1,
	})
	isGasless, err := ante.IsTxGasless(tx, suite.ctx, suite.keepers.OracleKeeper, suite.keepers.AdminKeeper)
	suite.Require().Error(err)
	suite.Require().False(isGasless)
	suite.Require().ErrorContains(err, nftoracletypes.ErrOracleAlreadyVoted.Error())
}

// ─────────────────────────────────────────────────────────────────────────────
// GaslessDecorator.AnteHandle – gas meter and priority
// ─────────────────────────────────────────────────────────────────────────────

// Gasless tx in CheckTx: infinite gas meter + oracle priority must be set.
func (suite *GaslessTestSuite) TestGaslessDecorator_CheckTx_InfiniteGasAndPriority() {
	suite.newPendingMintRequest(1)
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitMintResponse{
		Creator:       suite.oracleAddr.String(),
		MintRequestID: 1,
	})

	ctx := suite.ctx.WithIsCheckTx(true).WithGasMeter(storetypes.NewGasMeter(100_000))
	resultCtx, err := suite.gaslessDecorator.AnteHandle(ctx, tx, false, passthroughHandler)
	suite.Require().NoError(err)

	// An infinite gas meter has Limit() == math.MaxUint64.
	suite.Require().Equal(uint64(math.MaxUint64), resultCtx.GasMeter().Limit(),
		"gas meter must be infinite for gasless CheckTx")
	suite.Require().Equal(int64(ante.OraclePriority), resultCtx.Priority())
}

// Gasless tx in DeliverTx: infinite gas meter must also be set.
func (suite *GaslessTestSuite) TestGaslessDecorator_DeliverTx_InfiniteGas() {
	suite.newPendingMintRequest(1)
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitMintResponse{
		Creator:       suite.oracleAddr.String(),
		MintRequestID: 1,
	})

	ctx := suite.ctx.WithIsCheckTx(false).WithGasMeter(storetypes.NewGasMeter(100_000))
	resultCtx, err := suite.gaslessDecorator.AnteHandle(ctx, tx, false, passthroughHandler)
	suite.Require().NoError(err)

	suite.Require().Equal(uint64(math.MaxUint64), resultCtx.GasMeter().Limit(),
		"gas meter must be infinite for gasless DeliverTx")
}

// Non-gasless tx: wrapped (fee) decorators must be executed.
func (suite *GaslessTestSuite) TestGaslessDecorator_NonGasless_WrappedDecoratorsExecuted() {
	gd := ante.NewGaslessDecorator(
		[]sdk.AnteDecorator{&FailingDecorator{}},
		suite.keepers.OracleKeeper,
		suite.keepers.AdminKeeper,
	)
	tx := suite.createMockTx(&banktypes.MsgSend{})
	_, err := gd.AnteHandle(suite.ctx, tx, false, passthroughHandler)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "failing decorator executed")
}

// Gasless tx: wrapped (fee) decorators must be skipped.
func (suite *GaslessTestSuite) TestGaslessDecorator_Gasless_WrappedDecoratorsSkipped() {
	suite.newPendingMintRequest(1)
	gd := ante.NewGaslessDecorator(
		[]sdk.AnteDecorator{&FailingDecorator{}},
		suite.keepers.OracleKeeper,
		suite.keepers.AdminKeeper,
	)
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitMintResponse{
		Creator:       suite.oracleAddr.String(),
		MintRequestID: 1,
	})
	_, err := gd.AnteHandle(suite.ctx, tx, false, passthroughHandler)
	suite.Require().NoError(err, "wrapped decorators must be skipped for gasless txs")
}

// ─────────────────────────────────────────────────────────────────────────────
// VoteAloneDecorator
// ─────────────────────────────────────────────────────────────────────────────

func (suite *GaslessTestSuite) TestVoteAlone_SingleOracleMsg_Allowed() {
	tx := suite.createMockTx(&nftoracletypes.MsgSubmitMintResponse{Creator: suite.oracleAddr.String()})
	_, err := suite.voteAloneDecorator.AnteHandle(suite.ctx, tx, false, passthroughHandler)
	suite.Require().NoError(err)
}

func (suite *GaslessTestSuite) TestVoteAlone_SingleNonOracleMsg_Allowed() {
	tx := suite.createMockTx(&banktypes.MsgSend{})
	_, err := suite.voteAloneDecorator.AnteHandle(suite.ctx, tx, false, passthroughHandler)
	suite.Require().NoError(err)
}

func (suite *GaslessTestSuite) TestVoteAlone_MultipleNonOracleMsgs_Allowed() {
	tx := suite.createMockTx(&banktypes.MsgSend{}, &banktypes.MsgSend{})
	_, err := suite.voteAloneDecorator.AnteHandle(suite.ctx, tx, false, passthroughHandler)
	suite.Require().NoError(err)
}

func (suite *GaslessTestSuite) TestVoteAlone_OracleBundledWithOther_Rejected() {
	tx := suite.createMockTx(
		&nftoracletypes.MsgSubmitMintResponse{Creator: suite.oracleAddr.String()},
		&banktypes.MsgSend{},
	)
	_, err := suite.voteAloneDecorator.AnteHandle(suite.ctx, tx, false, passthroughHandler)
	suite.Require().Error(err)
	suite.Require().ErrorContains(err, "oracle votes cannot be bundled with other messages")
}

func (suite *GaslessTestSuite) TestVoteAlone_TwoOracleMsgs_Rejected() {
	tx := suite.createMockTx(
		&nftoracletypes.MsgSubmitMintResponse{Creator: suite.oracleAddr.String()},
		&nftoracletypes.MsgSubmitActionResponse{Creator: suite.oracleAddr.String()},
	)
	_, err := suite.voteAloneDecorator.AnteHandle(suite.ctx, tx, false, passthroughHandler)
	suite.Require().Error(err)
}

func (suite *GaslessTestSuite) TestVoteAlone_AllOracleMessageTypes_BundleRejected() {
	oracleMsgTypes := []sdk.Msg{
		&nftoracletypes.MsgSubmitMintResponse{Creator: suite.oracleAddr.String()},
		&nftoracletypes.MsgSubmitActionResponse{Creator: suite.oracleAddr.String()},
		&nftoracletypes.MsgSubmitVerifyCollectionOwner{Creator: suite.oracleAddr.String()},
	}
	for _, oracleMsg := range oracleMsgTypes {
		tx := suite.createMockTx(oracleMsg, &banktypes.MsgSend{})
		_, err := suite.voteAloneDecorator.AnteHandle(suite.ctx, tx, false, passthroughHandler)
		suite.Require().Error(err, "oracle msg type %T must be rejected when bundled", oracleMsg)
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Oracle vote-height KV store
// ─────────────────────────────────────────────────────────────────────────────

func (suite *GaslessTestSuite) TestOracleLastVoteHeight_SetAndGet() {
	ctx100 := suite.ctx.WithBlockHeight(100)
	suite.keepers.OracleKeeper.SetOracleLastVoteHeight(ctx100, suite.oracleAddr, 100)
	got := suite.keepers.OracleKeeper.GetOracleLastVoteHeight(ctx100, suite.oracleAddr)
	suite.Require().Equal(int64(100), got)
}

func (suite *GaslessTestSuite) TestOracleLastVoteHeight_NeverVoted_ReturnsZero() {
	got := suite.keepers.OracleKeeper.GetOracleLastVoteHeight(suite.ctx, suite.nonOracleAddr)
	suite.Require().Equal(int64(0), got)
}

func (suite *GaslessTestSuite) TestOracleLastVoteHeight_Delete() {
	ctx100 := suite.ctx.WithBlockHeight(100)
	suite.keepers.OracleKeeper.SetOracleLastVoteHeight(ctx100, suite.oracleAddr, 100)
	suite.keepers.OracleKeeper.DeleteOracleLastVoteHeight(ctx100, suite.oracleAddr)
	got := suite.keepers.OracleKeeper.GetOracleLastVoteHeight(ctx100, suite.oracleAddr)
	suite.Require().Equal(int64(0), got)
}

func (suite *GaslessTestSuite) TestCleanupOldVoteHeights_RemovesOldEntries() {
	oracle1 := sdk.AccAddress("oracle1_____________")
	oracle2 := sdk.AccAddress("oracle2_____________")

	ctxOld := suite.ctx.WithBlockHeight(1)
	ctxNew := suite.ctx.WithBlockHeight(200)

	// oracle1 voted at block 1, oracle2 voted at block 200.
	suite.keepers.OracleKeeper.SetOracleLastVoteHeight(ctxOld, oracle1, 1)
	suite.keepers.OracleKeeper.SetOracleLastVoteHeight(ctxNew, oracle2, 200)

	// Cleanup at height 200, maxAge=100 → cutoff=100, entries < 100 are removed.
	suite.keepers.OracleKeeper.CleanupOldVoteHeights(ctxNew, 100)

	// oracle1 (height 1) should be removed.
	suite.Require().Equal(int64(0), suite.keepers.OracleKeeper.GetOracleLastVoteHeight(ctxNew, oracle1))
	// oracle2 (height 200) should remain.
	suite.Require().Equal(int64(200), suite.keepers.OracleKeeper.GetOracleLastVoteHeight(ctxNew, oracle2))
}

// ─────────────────────────────────────────────────────────────────────────────
// Constructors and constants
// ─────────────────────────────────────────────────────────────────────────────

func (suite *GaslessTestSuite) TestOraclePriorityConstant() {
	suite.Require().Greater(int64(ante.OraclePriority), int64(1_000_000))
	suite.Require().Less(int64(ante.OraclePriority), int64(9_223_372_036_854_775_807))
}

func (suite *GaslessTestSuite) TestNewGaslessDecorator() {
	d := ante.NewGaslessDecorator(
		[]sdk.AnteDecorator{MockFeeDecorator{}},
		suite.keepers.OracleKeeper,
		suite.keepers.AdminKeeper,
	)
	suite.Require().IsType(ante.GaslessDecorator{}, d)
}

func (suite *GaslessTestSuite) TestNewVoteAloneDecorator() {
	d := ante.NewVoteAloneDecorator()
	suite.Require().IsType(ante.VoteAloneDecorator{}, d)
}

// ─────────────────────────────────────────────────────────────────────────────
// Benchmark
// ─────────────────────────────────────────────────────────────────────────────

func BenchmarkGaslessDecorator(b *testing.B) {
	keepers := keepertest.NewCombinedKeepers(b)
	gd := ante.NewGaslessDecorator(
		[]sdk.AnteDecorator{MockFeeDecorator{}},
		keepers.OracleKeeper,
		keepers.AdminKeeper,
	)
	tx := &mockTx{msgs: []sdk.Msg{&banktypes.MsgSend{}}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gd.AnteHandle(keepers.Ctx, tx, false, passthroughHandler)
	}
}
