package probes

import (
	"fmt"

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
	mi, err := load.Misc()
	if err != nil {
		return fmt.Errorf("task info: %w", err)
	}
	t.Total = mi.ProcsTotal
	t.Running = mi.ProcsRunning
	return nil
}
