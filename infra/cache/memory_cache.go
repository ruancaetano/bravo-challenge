package cache

import (
	"time"
)

type MemoryCache struct {
	cache map[string]CacheObject
}

type CacheObject struct {
	ttl  time.Time
	data interface{}
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		cache: make(map[string]CacheObject),
	}
}

func (m *MemoryCache) Set(key string, value interface{}, ttl time.Duration) {
	m.cache[key] = CacheObject{
		ttl:  time.Now().Add(ttl),
		data: value,
	}
}

func (m *MemoryCache) Get(key string) interface{} {
	cacheOnKey, ok := m.cache[key]
	if !ok {
		return nil
	}

	if time.Now().Before(cacheOnKey.ttl) {
		return cacheOnKey.data
	}

	delete(m.cache, key)

	return nil
}
