package main

import (
	"context"
	"fmt"
	"os"

	"adnotanumber.com/griptop/services"
	"adnotanumber.com/griptop/ui"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {

	g, ctx := errgroup.WithContext(context.Background())
	_, cancel := context.WithCancel(ctx)

	defer cancel()

	sysinfodyn := make(chan services.SystemInfoDyn)
	sysinfostatic, err := services.GetInfoStatic()
	if err != nil {
		return fmt.Errorf("info error: %w", err)
	}

	g.Go(func() error {
		if err := ui.Run(ctx, sysinfodyn, sysinfostatic); err != nil {
			cancel()
			return fmt.Errorf("UI error: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		if err := services.Acquire(ctx, sysinfodyn); err != nil {
			cancel()
			return fmt.Errorf("acquisition error: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
