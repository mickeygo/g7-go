package monitor

import (
	"sync"
	"time"

	_ "github.com/robfig/cron/v3"
)

// scanner 扫描器
type scanner struct {
	frequency time.Duration // 扫描频率

	mu sync.Mutex
}
