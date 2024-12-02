package types

const (
	// ModuleName defines the module name
	ModuleName = "helloapp"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_helloapp"
)

var (
	ParamsKey = []byte("p_helloapp")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
