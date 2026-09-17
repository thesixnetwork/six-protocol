package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
)

func TestMsgUpdateParams_ValidateBasic(t *testing.T) {
	tests := []struct {
		name    string
		msg     MsgUpdateParams
		wantErr bool
	}{
		{
			name: "invalid authority address",
			msg: MsgUpdateParams{
				Authority: "invalid_address",
				Params:    DefaultParams(),
			},
			wantErr: true,
		}, {
			name: "invalid params",
			msg: MsgUpdateParams{
				Authority: sample.AccAddress(),
				Params:    Params{},
			},
			wantErr: true,
		}, {
			name: "valid message",
			msg: MsgUpdateParams{
				Authority: sample.AccAddress(),
				Params:    DefaultParams(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
