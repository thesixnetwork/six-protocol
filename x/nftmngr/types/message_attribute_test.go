package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/testutil/sample"
)

func TestMsgAddAttribute_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddAttribute
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddAttribute{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddAttribute{
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

func TestMsgUpdateSchemaAttribute_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateSchemaAttribute
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateSchemaAttribute{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateSchemaAttribute{
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

func TestMsgShowAttributes_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgShowAttributes
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgShowAttributes{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgShowAttributes{
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

func TestMsgSetAttributeOveriding_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetAttributeOveriding
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetAttributeOveriding{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetAttributeOveriding{
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

func TestMsgResyncAttributes_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgResyncAttributes
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgResyncAttributes{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgResyncAttributes{
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
