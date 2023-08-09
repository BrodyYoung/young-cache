package src

import (
	"src/lru"
	"sync"
)

type cache struct {
	lru        lru.cache
	mu         sync.Mutex
	cacheBytes int64
}

func (c cache) Get(key string) (v ByteView) {

	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView)
	}
	return
}

func (c cache) Add(key string, val ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, val)
}
