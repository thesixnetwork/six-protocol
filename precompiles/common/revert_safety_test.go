package common_test

import (
	"errors"
	"math/big"
	"strings"
	"testing"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/evmos/evmos/v20/x/evm/core/vm"
	"github.com/evmos/evmos/v20/x/evm/statedb"

	pcommon "github.com/thesixnetwork/six-protocol/v4/precompiles/common"
)

var (
	testStoreKey  = storetypes.NewKVStoreKey("revert_safety_test")
	testTransient = storetypes.NewTransientStoreKey("revert_safety_test_t")
	markerKey     = []byte("cosmos-write-marker")
	markerValue   = []byte("written-by-precompile")
	precompAddr   = ethcommon.HexToAddress("0x0000000000000000000000000000000000009999")
)

const testABIJSON = `[{"type":"function","name":"foo","inputs":[],"outputs":[],"stateMutability":"nonpayable"}]`

// nopKeeper satisfies statedb.Keeper; the tests only exercise cosmos-side
// writes performed by the executor through the cache context, so no EVM
// account state is needed.
type nopKeeper struct{}

func (nopKeeper) GetAccount(_ sdk.Context, _ ethcommon.Address) *statedb.Account   { return nil }
func (nopKeeper) GetState(_ sdk.Context, _ ethcommon.Address, _ ethcommon.Hash) ethcommon.Hash {
	return ethcommon.Hash{}
}
func (nopKeeper) GetCode(_ sdk.Context, _ ethcommon.Hash) []byte { return nil }
func (nopKeeper) ForEachStorage(_ sdk.Context, _ ethcommon.Address, _ func(key, value ethcommon.Hash) bool) {
}
func (nopKeeper) SetAccount(_ sdk.Context, _ ethcommon.Address, _ statedb.Account) error {
	return nil
}
func (nopKeeper) SetState(_ sdk.Context, _ ethcommon.Address, _ ethcommon.Hash, _ []byte) {}
func (nopKeeper) SetCode(_ sdk.Context, _ []byte, _ []byte)                               {}
func (nopKeeper) DeleteAccount(_ sdk.Context, _ ethcommon.Address) error                  { return nil }

// markerExecutor writes a marker into the transaction KV store (through the
// precompile cache context) and then exits the way the test dictates:
// returning an error, panicking out-of-gas, or succeeding.
type markerExecutor struct {
	exit func() error // called after the write; its return is Execute's error
}

func (e markerExecutor) Address() ethcommon.Address { return precompAddr }

func (e markerExecutor) RequiredGas(_ []byte, _ *abi.Method) uint64 { return 0 }

func (e markerExecutor) Execute(ctx sdk.Context, _ *abi.Method, _ ethcommon.Address, _ ethcommon.Address, _ []interface{}, _ *big.Int, _ bool, _ *vm.EVM) ([]byte, error) {
	ctx.KVStore(testStoreKey).Set(markerKey, markerValue)
	if err := e.exit(); err != nil {
		return nil, err
	}
	return nil, nil
}

func setupRevertTest(t *testing.T, exit func() error) (sdk.Context, *statedb.StateDB, *pcommon.Precompile, []byte) {
	t.Helper()
	parsedABI, err := abi.JSON(strings.NewReader(testABIJSON))
	require.NoError(t, err)

	ctx := testutil.DefaultContext(testStoreKey, testTransient)
	db := statedb.New(ctx, nopKeeper{}, statedb.NewEmptyTxConfig(ethcommon.Hash{}))

	p := pcommon.NewPrecompile(parsedABI, markerExecutor{exit: exit}, precompAddr, "revert-safety-test")
	input := parsedABI.Methods["foo"].ID
	return ctx, db, p, input
}

// runPrecompileFrame mirrors what the EVM does around a precompile call: take
// a frame snapshot, run the contract, and revert the frame when it errors
// (x/evm/core/vm/evm.go Call/CallCode/DelegateCall all follow this pattern).
func runPrecompileFrame(db *statedb.StateDB, p *pcommon.Precompile, input []byte) error {
	evm := &vm.EVM{StateDB: db}
	frame := db.Snapshot()
	_, err := p.Run(evm, ethcommon.Address{1}, ethcommon.Address{1}, input, nil, false)
	if err != nil {
		db.RevertToSnapshot(frame)
	}
	return err
}

// TestExecutorErrorRevertsCosmosWrites is the H-1 regression: an executor that
// partially writes and then errors must have those writes rolled back by the
// frame revert. Before the fix, the revert-protection journal entry was only
// registered after a successful Execute, so the error path left the writes
// unjournaled and the unconditional cacheCtx flush at Commit persisted them
// (state-through-revert, ASA-2026-002 class).
func TestExecutorErrorRevertsCosmosWrites(t *testing.T) {
	ctx, db, p, input := setupRevertTest(t, func() error { return errors.New("boom after write") })

	err := runPrecompileFrame(db, p, input)
	require.Error(t, err)

	require.NoError(t, db.Commit())
	require.Nil(t, ctx.KVStore(testStoreKey).Get(markerKey),
		"cosmos write from a failed precompile call must not survive the EVM revert")
}

// TestExecutorOutOfGasRevertsCosmosWrites covers the same hole on the
// out-of-gas path: the OOG panic recovered by HandleGasError skips everything
// after Execute, so only a pre-registered journal entry protects the writes.
func TestExecutorOutOfGasRevertsCosmosWrites(t *testing.T) {
	ctx, db, p, input := setupRevertTest(t, func() error {
		panic(storetypes.ErrorOutOfGas{Descriptor: "test"})
	})

	err := runPrecompileFrame(db, p, input)
	require.ErrorIs(t, err, vm.ErrOutOfGas)

	require.NoError(t, db.Commit())
	require.Nil(t, ctx.KVStore(testStoreKey).Get(markerKey),
		"cosmos write from an out-of-gas precompile call must not survive the EVM revert")
}

// TestExecutorSuccessPersistsCosmosWrites guards the normal path: a
// successful call's writes must reach the transaction store at Commit.
func TestExecutorSuccessPersistsCosmosWrites(t *testing.T) {
	ctx, db, p, input := setupRevertTest(t, func() error { return nil })

	err := runPrecompileFrame(db, p, input)
	require.NoError(t, err)

	require.NoError(t, db.Commit())
	require.Equal(t, markerValue, ctx.KVStore(testStoreKey).Get(markerKey))
}
