package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActionSignerConfigKeyPrefix is the prefix to retrieve all ActionSignerConfig
	ActionSignerConfigKeyPrefix = "ActionSignerConfig/value/"
)

// ActionSignerConfigKey returns the store key to retrieve a ActionSignerConfig from the index fields
func ActionSignerConfigKey(
	chain string,
) []byte {
	var key []byte

	chainBytes := []byte(chain)
	key = append(key, chainBytes...)
	key = append(key, []byte("/")...)

	return key
}
