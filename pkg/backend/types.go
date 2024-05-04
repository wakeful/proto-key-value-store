package backend

// Store a simple key-value store.
type Store interface {
	Get(key string) (string, bool)
	Set(key, value string) (string, string)
}
