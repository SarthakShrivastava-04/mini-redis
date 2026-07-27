package main

import (
	"fmt"
	"net"
	"errors"
)

func startServer(listener net.Listener) error {

	fmt.Println("Server listening on :6379")

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}

			fmt.Println("Accept error:", err)
			continue
		}

		fmt.Println("Client connected")
        
	    wg.Add(1)
		go handleClient(conn)
	}
}