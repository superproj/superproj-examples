package main

import (
	"context"
	"fmt"
	"time"

	gocache "github.com/patrickmn/go-cache"

	"github.com/superproj/onex/pkg/cache"
	gocachestore "github.com/superproj/onex/pkg/cache/store/gocache"
)

type Book struct {
	ID   int
	Name string
}

func main() {
	gocacheClient := gocache.New(5*time.Minute, 10*time.Minute)
	gocacheStore := gocachestore.NewGoCache(gocacheClient)

	// Initialize a load function that loads your data from a custom source
	loadFunction := func(ctx context.Context, key any) (*Book, error) {
		// ... retrieve value from available source
		return &Book{ID: 1, Name: "My test amazing book"}, nil
	}

	cacheManager := cache.NewLoadable[*Book](loadFunction, cache.New[*Book](gocacheStore))
	ctx := context.Background()
	if err := cacheManager.Set(ctx, "my-key1", &Book{0, "go"}); err != nil {
		panic(err)
	}
	cacheManager.Wait(ctx)

	value, err := cacheManager.Get(ctx, "my-key2")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get value", value)
	value, err = cacheManager.Get(ctx, "my-key2")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get value", value)
	cacheManager.Del(ctx, "my-key")
	cacheManager.Del(ctx, "my-key2")
}
