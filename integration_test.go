package main

import (
	"bufio"
	"net"
	"testing"
	"time"
)

func TestPingIntegration(t *testing.T) {

	go func() {
		listener, _ := net.Listen("tcp", ":6379")
		startServer(listener)
	}()

	time.Sleep(100 * time.Millisecond)

	conn, err := net.Dial("tcp", ":6379")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("PING\n"))
	if err != nil {
		t.Fatal(err)
	}

	reader := bufio.NewReader(conn)

	response, err := reader.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}

	if response != "PONG\n" {
		t.Fatalf("expected PONG, got %q", response)
	}
}