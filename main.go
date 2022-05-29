package main

import (
	"fmt"
	"os"

	"adnotanumber.com/griptop/services"
	"adnotanumber.com/griptop/ui"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

}

func run() error {
	quit := make(chan bool)
	sysinfodyn := make(chan services.SystemInfoDyn)
	go services.Acquire(quit, sysinfodyn)

	sysinfostatic, err := services.GetInfoStatic()
	if err != nil {
		return fmt.Errorf("info error: %w", err)
	}

	err = ui.Run(sysinfodyn, sysinfostatic)
	if err != nil {
		return fmt.Errorf("UI error: %w", err)
	}

	quit <- true

	return nil
}
