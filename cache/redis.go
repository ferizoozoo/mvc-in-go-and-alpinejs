package cache

import (
	"context"
	"time"

	"github.com/ferizoozoo/sake/internal"
	"github.com/redis/go-redis/v9"
)

type RedisCacheProvider struct {
	provider *redis.Client
}

func NewRedisCacheProvider(co internal.CacheOptions) *RedisCacheProvider {
	provider := redis.NewClient(&redis.Options{
		Addr:     co.Address,
		Password: co.Password,
	})
	return &RedisCacheProvider{
		provider,
	}
}

func (rc *RedisCacheProvider) Get(ctx context.Context, key string) interface{} {
	return rc.provider.Get(ctx, key)
}

func (rc *RedisCacheProvider) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) {
	rc.provider.Set(ctx, key, value, expiration)
}

func (rc *RedisCacheProvider) Delete(ctx context.Context, key string) {
	rc.provider.Del(ctx, key)
}
