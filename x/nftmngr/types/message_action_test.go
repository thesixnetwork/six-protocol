package types

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/six-protocol/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestMsgAddAction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddAction
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddAction{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddAction{
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

func TestMsgUpdateAction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateAction
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateAction{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateAction{
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

func TestMsgToggleAction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgToggleAction
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgToggleAction{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgToggleAction{
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
