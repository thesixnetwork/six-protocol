// Copyright (C) 2024 SIX Network
// This file is part of the modified EVM module from Ethermint (https://github.com/evmos/ethermint),
// and is licensed under the terms of the GNU Lesser General Public License v3
package types

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/eth/tracers/logger"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func TestFormatLogs(t *testing.T) {
	zeroUint256 := []uint256.Int{*uint256.NewInt(0)}
	zeroByte := []byte{5}
	zeroStorage := make(map[string]string)

	testCases := []struct {
		name string
		logs []logger.StructLog
		exp  []StructLogRes
	}{
		{
			"empty logs",
			[]logger.StructLog{},
			[]StructLogRes{},
		},
		{
			"non-empty stack",
			[]logger.StructLog{
				{
					Stack: zeroUint256,
				},
			},
			[]StructLogRes{
				{
					Pc:    uint64(0),
					Op:    "STOP",
					Stack: &[]string{fmt.Sprintf("%x", zeroUint256[0])},
				},
			},
		},
		{
			"non-empty memory",
			[]logger.StructLog{
				{
					Memory: zeroByte,
				},
			},
			[]StructLogRes{
				{
					Pc:     uint64(0),
					Op:     "STOP",
					Memory: &[]string{"05"},
				},
			},
		},
		{
			"non-empty storage",
			[]logger.StructLog{
				{
					Storage: make(map[common.Hash]common.Hash),
				},
			},
			[]StructLogRes{
				{
					Pc:      uint64(0),
					Op:      "STOP",
					Storage: &zeroStorage,
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := FormatLogs(tc.logs)

			require.Equal(t, tc.exp, actual)
		})
	}
}

func TestNewNoOpTracer(t *testing.T) {
	require.Equal(t, &NoOpTracer{}, NewNoOpTracer())
}
