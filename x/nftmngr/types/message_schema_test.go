package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/testutil/sample"
)

func TestMsgSetUriRetrievalMethod_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetUriRetrievalMethod
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetUriRetrievalMethod{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetUriRetrievalMethod{
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

func TestMsgCreateNFTSchema_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateNFTSchema
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateNFTSchema{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateNFTSchema{
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

func TestMsgSetBaseUri_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetBaseUri
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetBaseUri{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetBaseUri{
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

func TestMsgSetMetadataFormat_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetMetadataFormat
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetMetadataFormat{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetMetadataFormat{
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

func TestMsgSetMintauth_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetMintauth
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetMintauth{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetMintauth{
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

func TestMsgSetOriginChain_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetOriginChain
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetOriginChain{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetOriginChain{
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

func TestMsgSetOriginContract_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetOriginContract
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetOriginContract{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetOriginContract{
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
