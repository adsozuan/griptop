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

			name, err := pr.Name()
			if err != nil {
				return nil, err
			}

			usrname, err := pr.Username()
			if err != nil {
				// return nil, fmt.Errorf("username: %w", err)
				usrname = "default"
			}

			prio, err := pr.CPUAffinity()
			if err != nil {
				// return nil, err
				prio = []int32{0}
			}

			cpu, err := pr.CPUPercent()
			if err != nil {
				// return nil, err
				cpu = 0.0
			}

			thrd, err := pr.NumThreads()
			if err != nil {
				return nil, err
			}

			// mem, err := pr.MemoryInfoEx()
			// if err != nil {
			// 	return nil, fmt.Errorf("meminfo %w", err)
			// }

			rp := RunningProcess{
				pid:      pr.Pid,
				user:     usrname,
				priority: prio[0],
				cpu:      fmt.Sprintf("%f%%", cpu),
				mem:      fmt.Sprint(""),
				thread:   thrd,
				time:     "",
				name:     name,
			}
			outputs = append(outputs, &rp)
		}
	}
	return outputs, nil
}
