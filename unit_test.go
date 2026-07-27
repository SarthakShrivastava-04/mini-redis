package main

import (
	"strings"
	"testing"
	"time"
)

func resetStore() {
	store = make(map[string]string)
	expire = make(map[string]time.Time)
}

func TestPing(t *testing.T) {
	resetStore()

	result := executeCommand([]string{"PING"})

	if result != "PONG\n" {
		t.Errorf("expected PONG\\n, got %q", result)
	}
}

func TestSet(t *testing.T) {
	resetStore()

	result := executeCommand([]string{"SET", "name", "sarthak"})

	if result != "OK\n" {
		t.Errorf("expected OK, got %q", result)
	}

	if store["name"] != "sarthak" {
		t.Errorf("expected value sarthak")
	}
}

func TestGet(t *testing.T) {
	resetStore()

	store["name"] = "sarthak"

	result := executeCommand([]string{"GET", "name"})

	if result != "sarthak\n" {
		t.Errorf("expected sarthak, got %q", result)
	}
}

func TestDel(t *testing.T) {
	resetStore()

	store["name"] = "sarthak"

	result := executeCommand([]string{"DEL", "name"})

	if result != "DELETED\n" {
		t.Errorf("expected DELETED, got %q", result)
	}

	if _, ok := store["name"]; ok {
		t.Errorf("key still exists")
	}
}

func TestExists(t *testing.T) {
	resetStore()

	store["name"] = "sarthak"

	result := executeCommand([]string{"EXISTS", "name"})

	if result != "1\n" {
		t.Errorf("expected 1, got %q", result)
	}

	result = executeCommand([]string{"EXISTS", "city"})

	if result != "0\n" {
		t.Errorf("expected 0, got %q", result)
	}
}

func TestCount(t *testing.T) {
	resetStore()

	store["a"] = "1"
	store["b"] = "2"
	store["c"] = "3"

	result := executeCommand([]string{"COUNT"})

	if result != "3\n" {
		t.Errorf("expected 3, got %q", result)
	}
}

func TestKeys(t *testing.T) {
	resetStore()

	store["name"] = "sarthak"
	store["city"] = "mumbai"

	result := executeCommand([]string{"KEYS"})

	if !strings.Contains(result, "name") {
		t.Errorf("missing key name")
	}

	if !strings.Contains(result, "city") {
		t.Errorf("missing key city")
	}
}

func TestExpire(t *testing.T) {
	resetStore()

	executeCommand([]string{"SET", "token", "abc"})

	result := executeCommand([]string{"EXPIRE", "token", "1"})

	if result != "OK\n" {
		t.Errorf("expected OK, got %q", result)
	}

	time.Sleep(2 * time.Second)

	result = executeCommand([]string{"GET", "token"})

	if result != "ERR invalid key\n" {
		t.Errorf("expected expired key, got %q", result)
	}
}