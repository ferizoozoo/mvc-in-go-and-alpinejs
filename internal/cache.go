package internal

import (
	"context"
	"sync"
	"time"
)

var once sync.Once
var cacheInstance *Cache

type CacheProvider interface {
	Get(context.Context, string) interface{}
	Set(context.Context, string, interface{}, time.Duration)
	Delete(context.Context, string)
}

type CacheOptions struct {
	Address  string
	Password string
}

// TODO: maybe this struct is not needed and CacheProvider
//
//	can be used directly (since Cache also implements CacheProvider interface)
type Cache struct {
	provider CacheProvider
}

func (c *Cache) SetProvider(p CacheProvider) *Cache {
	c.provider = p
	return c
}

// singleton pattern for getting the cache
func GetCache() *Cache {
	once.Do(func() {
		cacheInstance = &Cache{}
	})
	return cacheInstance
}

func (c *Cache) Get(ctx context.Context, key string) interface{} {
	return c.provider.Get(ctx, key)
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) {
	c.provider.Set(ctx, key, value, expiration)
}

func (c *Cache) Delete(ctx context.Context, key string) {
	c.provider.Delete(ctx, key)
}
