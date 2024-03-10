package types

const (
	// ModuleName defines the module name
	ModuleName = "tttibc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tttibc"

	// Version defines the current version the IBC module supports
	Version = "tttibc-1"

	// PortID is the default port id that module binds to
	PortID = "tttibc"
)

var ParamsKey = []byte("p_tttibc")

// PortKey defines the key to store the port ID in store
var PortKey = KeyPrefix("tttibc-port-")

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	GameKey      = "Game/value/"
	GameCountKey = "Game/count/"
)
