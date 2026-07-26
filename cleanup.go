package main

import (
	"time"
)

func cleanupExpiredKeys() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {

		mu.Lock()

		for key, expiryTime := range expire {
			if time.Now().After(expiryTime) {
				delete(expire, key)
				delete(store, key)
			}
		}

		mu.Unlock()
	}
}
