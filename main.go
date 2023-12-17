package main

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"

	"github.com/superproj/onex/pkg/cache"
	gocachestore "github.com/superproj/onex/pkg/cache/gocache"
	"github.com/superproj/onex/pkg/cache/store"
)

func main() {
	gocacheClient := gocache.New(5*time.Minute, 10*time.Minute)
	gocacheStore := gocachestore.NewGoCache(gocacheClient)

	cacheManager := cache.New[string](gocacheStore)
	ctx := context.Background()
	err = cacheManager.Set(ctx, "my-key", "my-value", store.WithCost(100))
	if err != nil {
		panic(err)
	}
	ristrettoCache.Wait()

	value, err := cacheManager.Get(ctx, "my-key")
	fmt.Println("Get value", value)
	cacheManager.Delete(ctx, "my-key")
}
