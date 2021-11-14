package iface

type Cache interface {
	// Set inserts or updates the specified key-value pair.
	Set(key, value interface{}) error
	// SetWithExpire inserts or updates the specified key-value pair with an expiration time.
	SetWithExpire(key, value interface{}, expiration time.Duration) error
	// Get returns the value for the specified key if it is present in the cache.
	// If the key is not present in the cache and the cache has LoaderFunc,
	// invoke the `LoaderFunc` function and inserts the key-value pair in the cache.
	// If the key is not present in the cache and the cache does not have a LoaderFunc,
	// return KeyNotFoundError.
	Get(key interface{}) (interface{}, error)
	// GetIFPresent returns the value for the specified key if it is present in the cache.
	// Return KeyNotFoundError if the key is not present.
	GetIFPresent(key interface{}) (interface{}, error)
	// GetAll returns a map containing all key-value pairs in the cache.
	GetALL(checkExpired bool) map[interface{}]interface{}
	get(key interface{}, onLoad bool) (interface{}, error)
	// Remove removes the specified key from the cache if the key is present.
	// Returns true if the key was present and the key has been deleted.
	Remove(key interface{}) bool
	// Purge removes all key-value pairs from the cache.
	Purge()
	// Keys returns a slice containing all keys in the cache.
	Keys(checkExpired bool) []interface{}
	// Len returns the number of items in the cache.
	Len(checkExpired bool) int
	// Has returns true if the key exists in the cache.
	Has(key interface{}) bool

	statsAccessor
}