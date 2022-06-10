package probes

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

type RunningProcess struct {
	pid      int32
	user     string
	priority int32
	cpu      string
	mem      string
	thread   int32
	disk     string
	time     string
	name     string
}

func (r *RunningProcess) ToString() [9]string {
	line := [9]string{}
	line[0] = fmt.Sprint(r.pid)
	line[1] = r.user
	line[2] = fmt.Sprint(r.priority)
	line[3] = r.cpu
	line[4] = r.mem
	line[5] = fmt.Sprint(r.thread)
	line[6] = r.disk
	line[7] = r.time
	line[8] = r.name
	return line

}

func RunningProcesses() ([]*RunningProcess, error) {
	runnings, err := process.Processes()
	if err != nil {
		return nil, fmt.Errorf("processes: %w", err)
	}

	outputs := []*RunningProcess{}

	for _, pr := range runnings {
		if pr.Pid != 0 {
			name, _ := pr.Name()
			mem, _ := pr.MemoryInfoEx()
			usrname, _ := pr.Username()
			prio, err := pr.CPUAffinity()
			if err != nil {
				return nil, err
			}
			cpu, _ := pr.CPUPercent()
			thrd, _ := pr.NumThreads()

			rp := RunningProcess{
				pid:      pr.Pid,
				user:     usrname,
				priority: prio[0],
				cpu:      fmt.Sprintf("%f%%", cpu),
				mem:      fmt.Sprint(mem),
				thread:   thrd,
				time:     "",
				name:     name,
			}
			outputs = append(outputs, &rp)
		}
	}
	return outputs, nil
}
