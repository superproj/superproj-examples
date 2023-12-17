package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto"
	gocache "github.com/patrickmn/go-cache"

	"github.com/superproj/onex/pkg/cache"
	gocachestore "github.com/superproj/onex/pkg/cache/store/gocache"
	ristrettostore "github.com/superproj/onex/pkg/cache/store/ristretto"
)

func main() {
	gocacheClient := gocache.New(5*time.Minute, 10*time.Minute)
	gocacheStore := gocachestore.NewGoCache(gocacheClient)

	// Initialize Ristretto cache and Redis client
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1000,
		MaxCost:     100,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}
	ristrettoStore := ristrettostore.NewRistretto(ristrettoCache)

	cacheManager := cache.NewChain[string](
		cache.New[string](gocacheStore),
		cache.New[string](ristrettoStore),
	)

	ctx := context.Background()
	if err := cacheManager.Set(ctx, "my-key", "my-value"); err != nil {
		panic(err)
	}
	cacheManager.Wait(ctx)

	value, err := cacheManager.Get(ctx, "my-key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get value", value)
	cacheManager.Del(ctx, "my-key")
}
