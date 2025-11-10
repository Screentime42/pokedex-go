package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu				sync.Mutex
	store			map[string]cacheEntry
	interval		time.Duration
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}


func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	entry := cacheEntry{
		createdAt:	time.Now(),
		val:			val,
	}
	c.store[key] = entry
	c.mu.Unlock()
}


func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.store[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}


func NewCache(interval time.Duration) *Cache {
	c := &Cache {
		store:		make(map[string]cacheEntry),
		interval: 	interval,
	}
	go c.reapLoop()
	return c
} 


func (c *Cache) reapLoop() {
	for {
		time.Sleep(c.interval)

		c.mu.Lock()
		
		for key, entry := range c.store {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.store, key)
			}
		}
		c.mu.Unlock()
	}
}