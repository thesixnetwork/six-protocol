package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActionExecutorKeyPrefix is the prefix to retrieve all ActionExecutor
	ActionExecutorKeyPrefix = "ActionExecutor/value/"
)

// ActionExecutorKey returns the store key to retrieve a ActionExecutor from the index fields
func ActionExecutorKey(
	nftSchemaCode string,
	executorAddress string,
) []byte {
	var key []byte

	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)

	executorAddressBytes := []byte(executorAddress)
	key = append(key, executorAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}
