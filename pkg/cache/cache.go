package cache

import (
	"time"
)

// Cache is the interface
type Cache interface {
	// Get the cache by specified key.
	Get(key string) (interface{}, error)

	// Set a cached value with key and expire time.
	Put(key string, val interface{}, expire time.Duration) error

	// Delete cached value by key.
	// Should not return error if key not found
	Delete(key string) error

	// Clear all cache.
	ClearAll() error
}
