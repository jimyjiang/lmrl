package jobs

import (
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

var (
	scheduler *cron.Cron
	mux       sync.Mutex
)

func Init() {
	if scheduler == nil {
		mux.Lock()
		defer mux.Unlock()
		scheduler = cron.New(cron.WithLocation(time.Local))
	}
}
func Start() {
	if scheduler != nil {
		scheduler.Start()
		fmt.Println("定时任务已启动")
	}
}
