package linkycore

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var (
	CacheDB *redis.Client
)

func InitCache(opts *LinkyCoreOptions) {
	CacheDB = redis.NewClient(&redis.Options{
		Addr:     opts.CacheDBAddr,
		Password: opts.CacheDBPassword,
		DB:       opts.CacheDBIndex,
	})
}

func CacheGetItem(key string) (string, error) {
	return CacheDB.Get(ctx, key).Result()
}

func CacheSetItem(key string, value string) error {
	return CacheDB.Set(ctx, key, value, 0).Err()
}
