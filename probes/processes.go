package probes

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

type RunningProcess struct {
	pid  int32
	name string
	mem  string
}

func (r *RunningProcess) ToString() [9]string {
	line := [9]string{}
	line[0] = fmt.Sprint(r.pid)
	line[1] = r.name
	line[2] = r.mem
	return line

}

func RunningProcesses() ([]*RunningProcess, error) {
	runnings, err := process.Processes()
	if err != nil {
		return nil, fmt.Errorf("processes: %w", err)
	}

	outputs := []*RunningProcess{}

	for _, pr := range runnings {
		name, _ := pr.Name()
		mem, _ := pr.MemoryInfo()
		rp := RunningProcess{
			pid:  pr.Pid,
			name: name,
			mem:  fmt.Sprint(mem.VMS),
		}
		outputs = append(outputs, &rp)
	}
	return outputs, nil
}
