package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/conf"
	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/service"
	v1 "github.com/superproj/superproj-examples/miniusercenter/pkg/api/miniusercenter/v1"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, uc *service.UserCenterService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserCenterServer(srv, uc)
	return srv
}
