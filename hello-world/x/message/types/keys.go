package types

const (
	// ModuleName defines the module name
	ModuleName = "message"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_message"
)

var (
	ParamsKey = []byte("p_message")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
