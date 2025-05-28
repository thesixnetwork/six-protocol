package types

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/six-protocol/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestMsgUnwrapToken_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUnwrapToken
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUnwrapToken{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUnwrapToken{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
