package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func sanitizeInput(s string) string {
	var result []rune

	for _, r := range s {
		switch r {
		case '\b', 127: // Backspace or DEL
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		default:
			result = append(result, r)
		}
	}

	return string(result)
}

func handleClient(conn net.Conn) {
	defer wg.Done()

	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := sanitizeInput(scanner.Text())
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
