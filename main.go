package main

import (
	"adnotanumber.com/griptop/services"
	"adnotanumber.com/griptop/ui"
)

func main() {

	quit := make(chan bool)
	sysinfodyn := make(chan services.SystemInfoDyn)
	go services.Acquire(quit, sysinfodyn)

	sysinfostatic := services.GetInfoStatic()
	ui.Run(sysinfodyn, sysinfostatic)
	quit <- true
}
