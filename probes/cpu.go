package probes

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type CpuUsage struct {
	totalMemory       uint64
	usedMemoryPercent float64
}

func AcquireCpuUsage() ([]float64, error) {
	cpuPercent, err := cpu.Percent(500*time.Millisecond, false)
	if err != nil {
		return nil, fmt.Errorf("cpu info: %w", err)
	}

	return cpuPercent, nil
}

func GetCpuModelName() string {
	cpuinfo, _ := cpu.Info()
	return cpuinfo[0].ModelName
}
