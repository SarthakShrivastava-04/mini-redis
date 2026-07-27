package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func gracefulShutdown(listener net.Listener) {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan

	fmt.Println("\nShutting down server...")

	listener.Close()

	fmt.Println("Waiting for clients to finish...")

	wg.Wait()

	fmt.Println("Shutdown complete.")
}
