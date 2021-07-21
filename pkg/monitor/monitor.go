package monitor

import (
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

// scanner 扫描器
type scanner struct {
	frequency time.Duration // 扫描频率

	mu sync.Mutex
}

func (*scanner) Scan() {
	cron.New(func(c *cron.Cron) {

	})
}
