package main

import (
	"fmt"
	"os"
)

func main() {
    
    go cleanupExpiredKeys()

	if err := startServer(); err != nil {
		fmt.Println("Server failed", err)
		os.Exit(1)
	}
}
