package services

import (
	"context"
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
	Processes        []*probes.RunningProcess
}

type SystemInfoStatic struct {
	MemSize string
	Proc    string
}

func Acquire(ctx context.Context, sysinfodyn chan SystemInfoDyn) error {

	for {
		cpu, err := probes.AcquireCpuUsage()
		if err != nil {
			return fmt.Errorf("cpu: %w", err)
		}
		mem := probes.NewMemoryUsage()
		tasks := probes.NewTaskCountsProbe()

		err = mem.Acquire()
		if err != nil {
			return fmt.Errorf("mem: %w", err)
		}
		err = tasks.Acquire()
		if err != nil {
			return fmt.Errorf("tasks: %w", err)
		}
		processes, err := probes.RunningProcesses()
		if err != nil {
			return fmt.Errorf("processes: %w", err)
		}

		sysinfocurr := SystemInfoDyn{
			CpuUsage:         cpu[0],
			MemUsagePercent:  mem.UsedMemoryPercent,
			TotalTaskCount:   tasks.Total,
			RunningTaskCount: tasks.Running,
			Uptime:           probes.GetUptime(),
			Processes:        processes,
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
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
