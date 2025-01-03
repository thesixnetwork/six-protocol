package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/testutil/sample"
)

func TestMsgVoteCreateVirtualSchema_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgVoteCreateVirtualSchema
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgVoteCreateVirtualSchema{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgVoteCreateVirtualSchema{
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

func TestMsgVoteDisableVirtualSchema_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgVoteDisableVirtualSchema
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgVoteDisableVirtualSchema{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgVoteDisableVirtualSchema{
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

func TestMsgVoteEnableVirtualSchema_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgVoteEnableVirtualSchema
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgVoteEnableVirtualSchema{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgVoteEnableVirtualSchema{
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
