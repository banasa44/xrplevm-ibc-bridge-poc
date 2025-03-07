package types

const (
	// ModuleName defines the module name
	ModuleName = "xrplevmbridgepoc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_xrplevmbridgepoc"
)

var (
	ParamsKey = []byte("p_xrplevmbridgepoc")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
