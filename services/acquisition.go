package services

import (
	"adnotanumber.com/griptop/probes"
	"time"
)

type SystemInfoDyn struct {
	CpuUsage         float64
	MemUsagePercent  float64
	TotalTaskCount   int64
	RunningTaskCount int64
}

func Acquire(quit chan bool, sysinfodyn chan SystemInfoDyn) {

	for {
		cpu := probes.AcquireCpuUsage()
		mem := probes.NewMemoryUsage()
		mem.Update()

		sysinfocurr := SystemInfoDyn{
			CpuUsage:         cpu[0],
			MemUsagePercent:  mem.UsedMemoryPercent,
			TotalTaskCount:   0,
			RunningTaskCount: 0,
		}

		select {
		case <-quit:
			break
		default:
			sysinfodyn <- sysinfocurr
			time.Sleep(500 * time.Millisecond)
		}
	}
}
