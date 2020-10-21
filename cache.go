package linkycore

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var (
	CacheDB CustomHTTPClient
)

func InitCache(opts *LinkyCoreOptions) {
	CacheDB := redis.NewClient(&redis.Options{
		Addr:     opts.CacheDBAddr,
		Password: opts.CacheDBPassword,
		DB:       opts.CacheDBIndex,
	})
}

func CacheSetItem(key string, value string) {
	return CacheDB.Set(ctx, key, value, 0).Err()
}

func CacheGetItem(key string) {
	return CacheDB.Get(ctx, key).Result()
}
