package probes

import (
	"github.com/shirou/gopsutil/load"
)

type TaskCountsProbe struct {
	Total   int
	Running int
}

func NewTaskCountsProbe() *TaskCountsProbe {
	t := TaskCountsProbe{
		Total:   0,
		Running: 0,
	}
	return &t
}

func (t *TaskCountsProbe) Acquire() error {
	mi, _ := load.Misc() // not implemented in gopsutil
	t.Total = mi.ProcsTotal
	t.Running = mi.ProcsRunning
	return nil
}
