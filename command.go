package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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

		mu.Lock()
		defer mu.Unlock()

		if expiryTime, isThere := expire[parts[1]]; isThere {
			if time.Now().After(expiryTime) {
				delete(store, parts[1])
				delete(expire, parts[1])
				return "ERR invalid key\n"
			}
		}

		value, found := store[parts[1]]

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
		defer mu.RUnlock()

		keys := make([]string, 0, len(store))

		for key := range store {
			keys = append(keys, key)
		}

		return strings.Join(keys, " ") + "\n"

	case "EXPIRE":
		if len(parts) != 3 {
			return "ERR wrong EXPIRE command\n"
		}

		seconds, err := strconv.Atoi(parts[2])
		if err != nil {
			return "ERR invalid expiry time\n"
		}

		mu.Lock()
		defer mu.Unlock()

		if _, found := store[parts[1]]; !found {
			return "ERR invalid key\n"
		}

		expire[parts[1]] = time.Now().Add(time.Duration(seconds) * time.Second)

		return "OK\n"

	default:
		return "ERR unknown command\n"
	}
}
