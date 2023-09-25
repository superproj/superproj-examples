//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/google/wire"
)

// initApp 声明 injector 的函数签名
func initApp() (*UserCenterService, error) {
	// wire.Build 声明要获取一个 *UserCenterService 需要调用到哪些 Provider
	wire.Build(NewDB, NewStore, NewBiz, NewUserCenterService)

	return new(UserCenterService), nil //返回值没有实际意义，只需符合函数签名即可
}