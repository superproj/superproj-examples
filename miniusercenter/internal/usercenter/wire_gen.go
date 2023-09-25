// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package usercenter

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/biz"
	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/conf"
	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/server"
	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/service"
	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/store"
	"gorm.io/gorm"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *conf.Data, db *gorm.DB, logger log.Logger) (*kratos.App, func(), error) {
	datastore := store.NewStore(db)
	bizBiz := biz.NewBiz(datastore)
	userCenterService := service.NewUserCenterService(bizBiz)
	grpcServer := server.NewGRPCServer(confServer, userCenterService, logger)
	httpServer := server.NewHTTPServer(confServer, userCenterService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
	}, nil
}