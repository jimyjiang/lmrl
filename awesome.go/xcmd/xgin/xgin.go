package xgin

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"

	"lmrl/awesome.go/xcmd/xserver"
	"lmrl/awesome.go/xcmd/xtracing"

	"github.com/gin-gonic/gin"
)

type XGinServer struct {
	xserver.XServer
	Gin *gin.Engine
}

func Run(options ...Option) (err error) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer func() {
		stop()
	}()

	// 初始化 SDK
	otelShutdown, err := xtracing.SetupOTelSDK(ctx)
	if err != nil {
		return
	}
	// 优雅关闭
	defer func() {
		err = errors.Join(err, otelShutdown.Shutdown(ctx))
	}()

	// 启动HTTP服务
	server := NewXGinServer(options...)

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- server.Gin.Run(fmt.Sprintf("%s:%d", server.Host, server.Port))
	}()

	select {
	case err = <-srvErr:
		return
	case <-ctx.Done():
		stop()
	}
	return
}

func NewXGinServer(options ...Option) *XGinServer {
	server := &XGinServer{
		XServer: xserver.XServer{
			Host: "",
			Port: 8080,
		},
		Gin: gin.Default(),
	}
	for _, option := range options {
		option(server)
	}
	return server
}
