package _net

import (
	"fmt"
	"net"
	"testing"
)

func TestDemoSocketConnect(t *testing.T) {
	// listen
	host := "127.0.0.1:8080"
	go EchoSocketServer(host)

	conn, err := net.Dial("tcp4", host)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Write([]byte("hello"))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("client: %s\n", string(buf[:n]))
}
