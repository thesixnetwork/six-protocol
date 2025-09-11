// keeper/mocks/tokenmngr_mock.go
package mocks

import (
	"github.com/stretchr/testify/mock"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MockTokenmngrKeeper struct {
	mock.Mock
}

func (m *MockTokenmngrKeeper) AttoCoinConverter(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, amount sdk.Int) error {
	args := m.Called(ctx, from, to, amount)
	return args.Error(0)
}
