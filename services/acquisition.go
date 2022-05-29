package services

import (
	"fmt"
	"strconv"
	"time"

	"adnotanumber.com/griptop/probes"
)

const (
	rate time.Duration = 500 * time.Millisecond
)

type SystemInfoDyn struct {
	CpuUsage         float64
	MemUsagePercent  float64
	TotalTaskCount   int
	RunningTaskCount int
	Uptime           string
}

type SystemInfoStatic struct {
	MemSize string
	Proc    string
}

func Acquire(quit chan bool, sysinfodyn chan SystemInfoDyn) {

	for {
		cpu, err := probes.AcquireCpuUsage()
		mem := probes.NewMemoryUsage()
		tasks := probes.NewTaskCountsProbe()
		err = mem.Acquire()
		err = tasks.Acquire()
		if err != nil {
			fmt.Printf("acquisition: %w", err)
		}

		sysinfocurr := SystemInfoDyn{
			CpuUsage:         cpu[0],
			MemUsagePercent:  mem.UsedMemoryPercent,
			TotalTaskCount:   tasks.Total,
			RunningTaskCount: tasks.Running,
			Uptime:           probes.GetUptime(),
		}

		select {
		case <-quit:
			return
		default:
			sysinfodyn <- sysinfocurr
			time.Sleep(rate)
		}
	}
}

func GetInfoStatic() (SystemInfoStatic, error) {
	mem := probes.NewMemoryUsage()
	err := mem.Acquire()
	if err != nil {
		return SystemInfoStatic{}, fmt.Errorf("static info: %w", err)
	}

	cpumodel := probes.GetCpuModelName()

	sysinfostatic := SystemInfoStatic{
		MemSize: strconv.FormatUint(mem.TotalMemory, 10),
		Proc:    cpumodel,
	}

	return sysinfostatic, nil

}
