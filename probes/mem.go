package probes

import (
	"github.com/shirou/gopsutil/mem"
)

type MemoryUsage struct {
	TotalMemory       uint64
	UsedMemoryPercent float64
}

func NewMemoryUsage() MemoryUsage {
	memoryUsage := MemoryUsage{
		TotalMemory:       1,
		UsedMemoryPercent: 2,
	}
	return memoryUsage
}

func (m *MemoryUsage) Acquire() {
	v, _ := mem.VirtualMemory()
	m.TotalMemory = v.Total
	m.UsedMemoryPercent = v.UsedPercent
}
