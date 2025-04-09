package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

var oldTerminalState *term.State

func setRawMode(raw bool) {
	var err error

	if raw {
		oldTerminalState, err = term.MakeRaw(int(os.Stdin.Fd()))

		if err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}
	} else {
		term.Restore(int(os.Stdin.Fd()), oldTerminalState)
	}
}

func runInterruptNotifyMonitor() context.CancelFunc {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	go func() {
		<-ctx.Done()
		fmt.Println("\nExiting...")
		term.Restore(int(os.Stdin.Fd()), oldTerminalState)
		os.Exit(0)
	}()

	return stop
}
