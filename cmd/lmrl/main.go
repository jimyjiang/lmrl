package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"lmrl/awesome.go/xcmd/xgin"
	"lmrl/logic/jobs"
	"lmrl/logic/mp3file"
	"lmrl/router"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer func() {
		stop()
	}()
	mp3file.StartWorker(ctx)
	jobs.RegisterDownloadMp3Job()
	jobs.Start()
	ro := []xgin.Option{
		xgin.WithHost(""),
		xgin.WithPort(3001),
		xgin.Use(otelgin.Middleware(hostName)),
		xgin.Gin(router.Init),
	}

	if err := xgin.Run(ro...); err != nil {
		log.Fatalln(err)
	}
}
