//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

//go:generate go run github.com/google/wire/cmd/wire

import (
	"io"
	"os"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// initApp 声明 injector 的函数签名
func initApp(db *gorm.DB) *UserCenterService {
	// wire.Build 声明要获取一个 *UserCenterService 需要调用到哪些 Provider
	wire.Build(
		NewStore,
		NewBiz,
		wire.Struct(new(UserCenterService), "*"), // struct 属性注入
		wire.Value(10),                           // 值绑定
		wire.InterfaceValue(new(io.Reader), os.Stdin), // 接口值绑定
	)

	return nil //返回值没有实际意义，只需符合函数签名即可
}
