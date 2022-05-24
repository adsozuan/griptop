package services

import (
	"adnotanumber.com/griptop/probes"
	"strconv"
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
	mem := probes.NewMemoryUsage()
	mem.Acquire()
	cpumodel := probes.GetCpuModelName()

	sysinfostatic := SystemInfoStatic{
		MemSize: strconv.FormatUint(mem.TotalMemory, 10),
		Proc:    cpumodel,
	}

	return sysinfostatic

}
