package cache

import "time"

// Item represents a cache item with its value and expiration time.
type Item struct {
	Value     []byte
	ExpiresAt time.Time
}
