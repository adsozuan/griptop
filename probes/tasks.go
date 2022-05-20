package probes

import "github.com/shirou/gopsutil/load"

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

func (t *TaskCountsProbe) Acquire() {
	mi, _ := load.Misc()
	t.Total = mi.ProcsTotal
	t.Running = mi.ProcsRunning
}
