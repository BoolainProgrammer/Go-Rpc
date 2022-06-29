package cache

import (
	"context"
	"fmt"
	"github.com/coocood/freecache"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"time"
)

type FreeCache struct {
	cacheManger *cache.Cache
}
func NewFreeCache(cfg *Config) *FreeCache {
	freecacheStore := store.NewFreecache(freecache.NewCache(1000), &store.Options{
		Expiration: 10 * time.Second,
	})

	free :=&FreeCache{
		cacheManger:  cache.New(freecacheStore),
	}
	ctx := context.Background()

	arr := []byte{1,2,3}
	free.Set(ctx, "my-key", arr, &store.Options{
		Expiration: 15*time.Second, // Override default value of 10 seconds defined in the store
	})

	fmt.Println(free.Get(ctx, "my-key"))

	return free
}

func (cache *FreeCache) Get(ctx context.Context, key interface{}) (interface{}, error) {
	return cache.cacheManger.Get(ctx, key)
}

func (cache *FreeCache) Set(ctx context.Context, key, object interface{}, options *store.Options) error {
	return cache.cacheManger.Set(ctx, key, object, options)
}

func (cache *FreeCache) Delete(ctx context.Context, key interface{}) error {
	return cache.cacheManger.Delete(ctx, key)
}

func (cache *FreeCache) Clear(ctx context.Context) error {
	return cache.cacheManger.Clear(ctx)
}