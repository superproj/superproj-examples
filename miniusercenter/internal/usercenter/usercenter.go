package usercenter

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/conf"
)

// go build -ldflags "-X github.com/superproj/superproj-examples/miniusercenter/internal/usercenter.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string

	id, _ = os.Hostname()
)

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(gs, hs),
	)
}

type Server struct {
	app     *kratos.App
	cleanup func()
}

func NewServer(bc conf.Bootstrap) (*Server, error) {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		bc.Data.Db.Username,
		bc.Data.Db.Password,
		bc.Data.Db.Addr,
		bc.Data.Db.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, db, logger)
	if err != nil {
		return nil, err
	}

	return &Server{app: app, cleanup: cleanup}, nil
}

func (s *Server) Run() {
	defer s.cleanup()

	if err := s.app.Run(); err != nil {
		panic(err)
	}
}
