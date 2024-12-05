package types

const (
	// ModuleName defines the module name
	ModuleName = "helloworld"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_helloworld"
)

var (
	ParamsKey = []byte("p_helloworld")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
