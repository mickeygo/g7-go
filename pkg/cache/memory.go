package cache

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

// DefaultEvery sets a timer for how often to recycle the expired cache items in memory (in seconds)
var DefaultEvery = 60 // 1 minute

var (
	ErrKeyExpired  = errors.New("the key is expired")
	ErrKeyNotExist = errors.New("the key isn't exist")
)

// MemoryItem stores memory cache item.
type MemoryItem struct {
	val         interface{}
	createdTime time.Time
	lifespan    time.Duration
}

func (mi *MemoryItem) isExpire() bool {
	// 0 means forever
	if mi.lifespan == 0 {
		return false
	}
	return time.Since(mi.createdTime) > mi.lifespan
}

// MemoryCache is a memory cache adapter.
// Contains a RW locker for safe map storage.
type MemoryCache struct {
	sync.RWMutex
	dur   time.Duration
	items map[string]*MemoryItem
	Every int // run an expiration check Every clock time
}

// NewMemoryCache returns a new MemoryCache.
func NewMemoryCache() Cache {
	cache := MemoryCache{items: make(map[string]*MemoryItem)}
	return &cache
}

// Get returns cache from memory.
// If non-existent or expired, return nil.
func (bc *MemoryCache) Get(key string) (interface{}, error) {
	bc.RLock()
	defer bc.RUnlock()
	if itm, ok := bc.items[key]; ok {
		if itm.isExpire() {
			return nil, ErrKeyExpired
		}
		return itm.val, nil
	}
	return nil, ErrKeyNotExist
}

// Put puts cache into memory.
// If lifespan is 0, it will never overwrite this value unless restarted
func (bc *MemoryCache) Put(key string, val interface{}, expire time.Duration) error {
	bc.Lock()
	defer bc.Unlock()
	bc.items[key] = &MemoryItem{
		val:         val,
		createdTime: time.Now(),
		lifespan:    expire,
	}
	return nil
}

// Delete cache in memory.
// If the key is not found, it will not return error
func (bc *MemoryCache) Delete(key string) error {
	bc.Lock()
	defer bc.Unlock()
	delete(bc.items, key)
	return nil
}

// ClearAll deletes all cache in memory.
func (bc *MemoryCache) ClearAll() error {
	bc.Lock()
	defer bc.Unlock()
	bc.items = make(map[string]*MemoryItem)
	return nil
}

// Start starts memory cache. Checks expiration in every clock time.
func (bc *MemoryCache) Start(config string) error {
	var cf map[string]int
	if err := json.Unmarshal([]byte(config), &cf); err != nil {
		// return berror.Wrapf(err, InvalidMemoryCacheCfg, "invalid config, please check your input: %s", config)
		return err
	}

	if _, ok := cf["interval"]; !ok {
		cf = make(map[string]int)
		cf["interval"] = DefaultEvery
	}

	dur := time.Duration(cf["interval"]) * time.Second
	bc.Every = cf["interval"]
	bc.dur = dur
	go bc.vacuum()

	return nil
}

// check expiration.
func (bc *MemoryCache) vacuum() {
	bc.RLock()
	every := bc.Every
	bc.RUnlock()

	if every < 1 {
		return
	}
	for {
		<-time.After(bc.dur)
		bc.RLock()
		if bc.items == nil {
			bc.RUnlock()
			return
		}
		bc.RUnlock()
		if keys := bc.expiredKeys(); len(keys) != 0 {
			bc.clearItems(keys)
		}
	}
}

// expiredKeys returns keys list which are expired.
func (bc *MemoryCache) expiredKeys() (keys []string) {
	bc.RLock()
	defer bc.RUnlock()
	for key, itm := range bc.items {
		if itm.isExpire() {
			keys = append(keys, key)
		}
	}
	return
}

// ClearItems removes all items who's key is in keys
func (bc *MemoryCache) clearItems(keys []string) {
	bc.Lock()
	defer bc.Unlock()
	for _, key := range keys {
		delete(bc.items, key)
	}
}
