package testutil

import (
	"context"
	"time"

	"cosmossdk.io/core/address"
	"cosmossdk.io/math"
	"github.com/stretchr/testify/mock"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// MockStakingKeeper is a mock implementation of the staking keeper interface
type MockStakingKeeper struct {
	mock.Mock
}

func NewMockStakingKeeper() *MockStakingKeeper {
	return &MockStakingKeeper{}
}

func (m *MockStakingKeeper) GetDelegatorDelegations(ctx context.Context, delegator sdk.AccAddress, maxRetrieve uint16) ([]stakingtypes.Delegation, error) {
	args := m.Called(ctx, delegator, maxRetrieve)
	return args.Get(0).([]stakingtypes.Delegation), args.Error(1)
}

func (m *MockStakingKeeper) RemoveDelegation(ctx context.Context, delegation stakingtypes.Delegation) error {
	args := m.Called(ctx, delegation)
	return args.Error(0)
}

func (m *MockStakingKeeper) SetDelegation(ctx context.Context, delegation stakingtypes.Delegation) error {
	args := m.Called(ctx, delegation)
	return args.Error(0)
}

func (m *MockStakingKeeper) GetUnbondingDelegations(ctx context.Context, delegator sdk.AccAddress, maxRetrieve uint16) ([]stakingtypes.UnbondingDelegation, error) {
	args := m.Called(ctx, delegator, maxRetrieve)
	return args.Get(0).([]stakingtypes.UnbondingDelegation), args.Error(1)
}

func (m *MockStakingKeeper) RemoveUnbondingDelegation(ctx context.Context, ubd stakingtypes.UnbondingDelegation) error {
	args := m.Called(ctx, ubd)
	return args.Error(0)
}

func (m *MockStakingKeeper) SetUnbondingDelegation(ctx context.Context, ubd stakingtypes.UnbondingDelegation) error {
	args := m.Called(ctx, ubd)
	return args.Error(0)
}

func (m *MockStakingKeeper) InsertUBDQueue(ctx context.Context, ubd stakingtypes.UnbondingDelegation, completionTime time.Time) error {
	args := m.Called(ctx, ubd, completionTime)
	return args.Error(0)
}

func (m *MockStakingKeeper) GetRedelegations(ctx context.Context, delegator sdk.AccAddress, maxRetrieve uint16) ([]stakingtypes.Redelegation, error) {
	args := m.Called(ctx, delegator, maxRetrieve)
	return args.Get(0).([]stakingtypes.Redelegation), args.Error(1)
}

func (m *MockStakingKeeper) RemoveRedelegation(ctx context.Context, red stakingtypes.Redelegation) error {
	args := m.Called(ctx, red)
	return args.Error(0)
}

func (m *MockStakingKeeper) SetRedelegation(ctx context.Context, red stakingtypes.Redelegation) error {
	args := m.Called(ctx, red)
	return args.Error(0)
}

func (m *MockStakingKeeper) InsertRedelegationQueue(ctx context.Context, red stakingtypes.Redelegation, completionTime time.Time) error {
	args := m.Called(ctx, red, completionTime)
	return args.Error(0)
}

func (m *MockStakingKeeper) GetDelegatorBonded(ctx context.Context, delegator sdk.AccAddress) (math.Int, error) {
	args := m.Called(ctx, delegator)
	return args.Get(0).(math.Int), args.Error(1)
}

func (m *MockStakingKeeper) GetDelegatorUnbonding(ctx context.Context, delegator sdk.AccAddress) (math.Int, error) {
	args := m.Called(ctx, delegator)
	return args.Get(0).(math.Int), args.Error(1)
}

// MockAccountKeeper is a mock implementation of the account keeper interface
type MockAccountKeeper struct {
	mock.Mock
}

func NewMockAccountKeeper() *MockAccountKeeper {
	return &MockAccountKeeper{}
}

func (m *MockAccountKeeper) AddressCodec() address.Codec {
	args := m.Called()
	return args.Get(0).(address.Codec)
}

// MockAddressCodec is a mock implementation of the address codec interface
type MockAddressCodec struct {
	BytesToStringFunc func([]byte) (string, error)
	StringToBytesFunc func(string) ([]byte, error)
}

func (m *MockAddressCodec) BytesToString(addr []byte) (string, error) {
	if m.BytesToStringFunc != nil {
		return m.BytesToStringFunc(addr)
	}
	return sdk.AccAddress(addr).String(), nil
}

func (m *MockAddressCodec) StringToBytes(addr string) ([]byte, error) {
	if m.StringToBytesFunc != nil {
		return m.StringToBytesFunc(addr)
	}
	accAddr, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, err
	}
	return accAddr.Bytes(), nil
}
