//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import "github.com/google/wire"

func InitializeFooBar() FooBar {
	wire.Build(Set)
	return FooBar{}
}
