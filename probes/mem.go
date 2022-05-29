package probes

import (
	"fmt"

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

func (m *MemoryUsage) Acquire() error {
	v, err := mem.VirtualMemory()
	if err != nil {
		return fmt.Errorf("Mem: %w", err)
	}

	m.TotalMemory = v.Total
	m.UsedMemoryPercent = v.UsedPercent
	return nil
}
