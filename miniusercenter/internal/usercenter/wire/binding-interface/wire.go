//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

// initApp 声明 injector 的函数签名
func initApp(db *gorm.DB) *UserCenterService {
	wire.Build(StoreProviderSet, BizProviderSet, NewUserCenterService)

	return nil //返回值没有实际意义，只需符合函数签名即可
}
