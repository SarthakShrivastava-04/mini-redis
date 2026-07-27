package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleClient(conn net.Conn) {
	defer wg.Done()

	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Received:", text)

		parts := strings.Fields(text)

		if len(parts) == 0 {
			continue
		}

		response := executeCommand(parts)

		_, err := conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Read error:", err)
	}

	fmt.Println("Client disconnected")
}