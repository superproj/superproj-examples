package main

import (
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	cl := logger.With(zap.String("user.id", "colin"))

	cl.Debug("this is a debug message")
	cl.Info("this is a info message")
}

// 打印结果
// {"level":"debug","msg":"this is a debug message","user.id":"colin"}
// {"level":"info","msg":"this is a info message","user.id":"colin"}
