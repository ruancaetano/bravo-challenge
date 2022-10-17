package cache

import "time"

type CacheInterface interface {
	Set(key string, value interface{}, ttl time.Duration)
	Get(key string) interface{}
}
