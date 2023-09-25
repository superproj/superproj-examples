package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/conf"
	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/service"
	v1 "github.com/superproj/superproj-examples/miniusercenter/pkg/api/miniusercenter/v1"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, uc *service.UserCenterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterUserCenterHTTPServer(srv, uc)
	return srv
}
