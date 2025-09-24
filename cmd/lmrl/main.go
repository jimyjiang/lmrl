package main

import (
	"log"
	"os"

	"lmrl/logic/jobs"
	"lmrl/router"

	"awesome.go/xcmd/xgin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
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
