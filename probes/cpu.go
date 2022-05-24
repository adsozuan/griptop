package probes

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type CpuUsage struct {
	totalMemory       uint64
	usedMemoryPercent float64
}

func AcquireCpuUsage() []float64 {
	cpuPercent, _ := cpu.Percent(500*time.Millisecond, false)

	return cpuPercent
}

func GetCpuModelName() string {
	cpuinfo, _ := cpu.Info()
	return cpuinfo[0].ModelName
}
