package lru

import (
	"container/list"
	"sync"
	"time"
)

// data in LRU
type data struct {
	key   interface{}
	value interface{}
	ttl   time.Time
}

// LRU list + Map
type LRU struct {
	mu      sync.RWMutex
	size    int
	list    *list.List
	items   map[interface{}]*list.Element
	expiry  time.Duration
	onEvict func(key interface{}, value interface{})
}

// Get data in LRU
func (c *LRU) Get(k interface{}) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Lock()

	if ele, ok := c.items[k]; ok {
		c.list.MoveToFront(ele)
		return ele.Value.(*data).value, true
	}

	return nil, false
}
