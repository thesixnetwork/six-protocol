package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/six-protocol/testutil/sample"
)

func TestMsgCreateVirtualSchema_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateVirtualSchema
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateVirtualSchema{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateVirtualSchema{
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

func TestMsgDeleteVirtualSchema_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteVirtualSchema
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteVirtualSchema{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteVirtualSchema{
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
