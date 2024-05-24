package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (ce *cacheEntry) isExpired(interval time.Duration) bool {
	return time.Since(ce.createdAt) > interval
}

type Cache struct {
	interval     time.Duration
	cacheEntries map[string]cacheEntry
	mu           sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{interval: interval, cacheEntries: make(map[string]cacheEntry)}
	c.reapLoop()
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ce := cacheEntry{createdAt: time.Now(), val: val}
	c.cacheEntries[key] = ce
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.cacheEntries[key]
	if !ok {
		return nil, ok
	} else {
		return val.val, ok
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				c.mu.Lock()
				for k := range c.cacheEntries {
					ce := c.cacheEntries[k]
					if ce.isExpired(c.interval) {
						delete(c.cacheEntries, k)
					}
				}
				c.mu.Unlock()
			}
		}
	}()
}
