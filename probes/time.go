package probes

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"time"
)

func GetUptime() string {
	uptime, _ := host.Uptime()
	t := time.Unix(int64(uptime), 0)
	return fmt.Sprintf(t.Format("15:04:05"))
}
