package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {

	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Received:", text)

		_, err := conn.Write([]byte(text + "\n"))

		if err != nil {
			fmt.Println("writing error:", err)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Read error:", err)
	}

	fmt.Println("Client disconnected")

}
