package cache

import (
	"sync"
	"time"
)

// Cache is a simple in-memory cache that stores key-value pairs with an expiration time.
type Cache struct {
	mu    sync.RWMutex
	items map[string]Item
}

func New() *Cache {
	return &Cache{
		items: make(map[string]Item),
	}
}

func (c *Cache) Set(key string, value []byte, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = Item{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
	}
}

// Get retrieves the value associated with the given key. 
// If the key does not exist or has expired, it returns nil.
func (c *Cache) Get(key string) []byte {
	c.mu.RLock()
	item, exists := c.items[key]
	c.mu.RUnlock()

	if !exists {
		return nil
	}

	if time.Now().After(item.ExpiresAt) {
		c.mu.Lock()
		delete(c.items, key)
		c.mu.Unlock()
		return nil
	}

	return item.Value
}
