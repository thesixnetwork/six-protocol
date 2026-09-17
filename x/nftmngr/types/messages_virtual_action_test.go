package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
)

func TestMsgCreateVirtualAction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateVirtualAction
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateVirtualAction{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateVirtualAction{
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

func TestMsgUpdateVirtualAction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateVirtualAction
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateVirtualAction{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateVirtualAction{
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

func TestMsgDeleteVirtualAction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteVirtualAction
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteVirtualAction{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteVirtualAction{
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
