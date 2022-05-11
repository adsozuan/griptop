package main

import (
	"adnotanumber.com/griptop/probes"
	"adnotanumber.com/griptop/services"
	"adnotanumber.com/griptop/ui"
	"fmt"
	"time"
)

func main() {
	memUsage := probes.NewMemoryUsage()
	memUsage.Update()
	s := fmt.Sprintf("%v", memUsage)
	fmt.Println(s)

	quit := make(chan bool)
	go services.Acquire(quit)
	ui.Run()
	time.Sleep(3 * time.Second)
	quit <- true
}
