package main

import (
	"fmt"
	"os"
    "net"
)

func main() {
    go cleanupExpiredKeys()

	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		os.Exit(1)
	}

	go gracefulShutdown(listener)

	if err := startServer(listener); err != nil {
        fmt.Println("Server stopped:", err)
	}
}
