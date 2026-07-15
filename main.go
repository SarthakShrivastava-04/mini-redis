package main

import (
	"fmt"
	"os"
)

func main() {

	if err := startServer(); err != nil {
		fmt.Println("Server failed", err)
		os.Exit(1)
	}
}
