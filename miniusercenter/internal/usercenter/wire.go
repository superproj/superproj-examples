//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package usercenter

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/biz"
	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/conf"
	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/server"
	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/service"
	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/store"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *gorm.DB, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, store.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
