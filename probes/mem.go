package probes

import (
	"github.com/shirou/gopsutil/mem"
)

type MemoryUsage struct {
	totalMemory       uint64
	usedMemoryPercent float64
}

func NewMemoryUsage() MemoryUsage {
	memoryUsage := MemoryUsage{
		totalMemory:       1,
		usedMemoryPercent: 2,
	}
	return memoryUsage
}

func (m *MemoryUsage) Update() {
	v, _ := mem.VirtualMemory()
	m.totalMemory = v.Total
	m.usedMemoryPercent = v.UsedPercent
}
