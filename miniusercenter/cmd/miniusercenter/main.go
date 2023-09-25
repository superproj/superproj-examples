package main

import (
	"flag"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	_ "go.uber.org/automaxprocs"

	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter"
	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/conf"
)

// flagconf is the config flag.
var flagconf string

func init() {
	flag.StringVar(&flagconf, "config", "config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()

	c := config.New(config.WithSource(file.NewSource(flagconf)))
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	server, err := usercenter.NewServer(bc)
	if err != nil {
		panic(err)
	}

	server.Run()
}
