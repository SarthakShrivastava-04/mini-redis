package main

import (
	"fmt"
	"net"
)

func startServer() error {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		return err
	}

	fmt.Println("Server listening on :6379")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		fmt.Println("Client connected")

		go handleClient(conn)
	}
}