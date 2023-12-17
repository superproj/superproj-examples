package main

import (
	"context"
	"fmt"
	"time"

	gocache "github.com/patrickmn/go-cache"

	"github.com/superproj/onex/pkg/cache"
	gocachestore "github.com/superproj/onex/pkg/cache/store/gocache"
)

func main() {
	gocacheClient := gocache.New(5*time.Minute, 10*time.Minute)
	gocacheStore := gocachestore.NewGoCache(gocacheClient)

	cacheManager := cache.New[string](gocacheStore)
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
