package types

const (
	// ModuleName defines the module name
	ModuleName = "protocoladmin"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_protocoladmin"

	// Version defines the current version the IBC module supports
	Version = "protocoladmin-1"

	// PortID is the default port id that module binds to
	PortID = "protocoladmin"
)

// PortKey defines the key to store the port ID in store
var PortKey = KeyPrefix("protocoladmin-port-")

func KeyPrefix(p string) []byte {
	return []byte(p)
}
