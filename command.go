package main

import (
	"fmt"
	"strings"
)

func executeCommand(parts []string) string {
	command := strings.ToUpper(parts[0])

	switch command {

	case "PING":
		return "PONG\n"

	case "SET":
		if len(parts) != 3 {
			return "ERR wrong SET command\n"
		}

		mu.Lock()
		store[parts[1]] = parts[2]
		mu.Unlock()

		return "OK\n"

	case "GET":
		if len(parts) != 2 {
			return "ERR wrong GET command\n"
		}

		mu.RLock()
		value, found := store[parts[1]]
		mu.RUnlock()

		if !found {
			return "ERR invalid key\n"
		}

		return value + "\n"

	case "DEL":
		if len(parts) != 2 {
			return "ERR wrong DEL command\n"
		}

		if _, found := store[parts[1]]; !found {
			return "ERR invalid key\n"
		}

		mu.Lock()
		delete(store, parts[1])
		mu.Unlock()

		return "DELETED\n"

	case "EXISTS":
		if len(parts) != 2 {
			return "ERR wrong EXISTS command\n"
		}

		mu.RLock()
		_, found := store[parts[1]]
		mu.RUnlock()

		if found {
			return "1\n"
		}

		return "0\n"

	case "COUNT":
		if len(parts) != 1 {
			return "ERR wrong COUNT command\n"
		}

		mu.RLock()
		count := len(store)
		mu.RUnlock()

		return fmt.Sprintf("%d\n", count)

	case "KEYS":
		if len(parts) != 1 {
			return "ERR wrong KEYS command\n"
		}

		mu.RLock()

		keys := make([]string, 0, len(store))

		for key := range store {
			keys = append(keys, key)
		}

		mu.RUnlock()

		return strings.Join(keys, " ") + "\n"

	default:
		return "ERR unknown command\n"
	}
}