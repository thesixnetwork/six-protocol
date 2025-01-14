package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/testutil/sample"
)

func TestMsgCreateActionSignerConfig_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateActionSignerConfig
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateActionSignerConfig{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateActionSignerConfig{
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

func TestMsgUpdateActionSignerConfig_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateActionSignerConfig
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateActionSignerConfig{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateActionSignerConfig{
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

func TestMsgDeleteActionSignerConfig_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteActionSignerConfig
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteActionSignerConfig{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteActionSignerConfig{
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
