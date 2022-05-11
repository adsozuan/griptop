package services

import (
	"adnotanumber.com/griptop/probes"
	"fmt"
	"time"
)

func Acquire(quit chan bool) {
	for {
		cpu := probes.AcquireCpuUsage()
		mem := probes.NewMemoryUsage()
		mem.Update()
		fmt.Printf("cpu %2.2f\n", cpu)
		fmt.Printf("mem %v\n", mem)

		select {
		case <-quit:
			break
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}
