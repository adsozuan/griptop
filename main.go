package main

import (
	"context"
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

	cancellingCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sysinfodyn := make(chan services.SystemInfoDyn)
	go services.Acquire(cancellingCtx, sysinfodyn)

	sysinfostatic, err := services.GetInfoStatic()
	if err != nil {
		return fmt.Errorf("info error: %w", err)
	}

	err = ui.Run(sysinfodyn, sysinfostatic)
	if err != nil {
		return fmt.Errorf("UI error: %w", err)
	}

	return nil
}
