package memcache

import "sync"

// Cache holds the cached values.
type Cache struct {
	cache map[string]cacheEntry
	mu    sync.RWMutex
}

type cacheEntry struct {
	value interface{}
	err   error
}

// Clear clears the cache state.
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(map[string]cacheEntry)
}

// GetOrCreate tries to get the value with the given cache key, if not found
// create will be called and cached.
// This method is thread safe. It also guarantees that the create func for a given
// key is invoked only once for this cache.
func (c *Cache) GetOrCreate(key string, create func() (interface{}, error)) (interface{}, error) {
	c.mu.RLock()
	entry, found := c.cache[key]
	c.mu.RUnlock()

	if found {
		return entry.value, entry.err
	}

	// Create it.
	value, err := create()

	c.mu.Lock()
	c.cache[key] = cacheEntry{value: value, err: err}
	c.mu.Unlock()

	return value, err
}
