package services

import (
	"adnotanumber.com/griptop/probes"
	"time"
)

const (
	rate time.Duration = 500 * time.Millisecond
)

type SystemInfoDyn struct {
	CpuUsage         float64
	MemUsagePercent  float64
	TotalTaskCount   int
	RunningTaskCount int
}

type SystemInfoStatic struct {
	MemSize string
	Proc    string
}

func Acquire(quit chan bool, sysinfodyn chan SystemInfoDyn) {

	for {
		cpu := probes.AcquireCpuUsage()
		mem := probes.NewMemoryUsage()
		tasks := probes.NewTaskCountsProbe()
		mem.Acquire()
		tasks.Acquire()

		sysinfocurr := SystemInfoDyn{
			CpuUsage:         cpu[0],
			MemUsagePercent:  mem.UsedMemoryPercent,
			TotalTaskCount:   tasks.Total,
			RunningTaskCount: tasks.Running,
		}

		select {
		case <-quit:
			break
		default:
			sysinfodyn <- sysinfocurr
			time.Sleep(rate)
		}
	}
}

func GetInfoStatic() SystemInfoStatic {
	sysinfostatic := SystemInfoStatic{
		MemSize: "32 GB",
		Proc:    "CPU bubl",
	}

	return sysinfostatic

}
