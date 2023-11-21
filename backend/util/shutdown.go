package util

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func gracefulShutdown(cancel context.CancelFunc) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-interruptChan
		// TODO: wrap this in log.METHOD
		fmt.Printf("%s received.\n", sig.String())
		// TODO: shut down any necessary things i.e. server
		cancel()
	}()
}
