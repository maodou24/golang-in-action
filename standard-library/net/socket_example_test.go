package net

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestNetSocketClientServerEcho(t *testing.T) {
	host := "127.0.0.1:8080"

	go ListenSocketEchoServer(context.Background(), host)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			conn, err := net.Dial("tcp4", host)
			if err != nil {
				t.Error(err)
				return
			}
			defer conn.Close()

			if _, err := conn.Write([]byte("hello")); err != nil {
				t.Error(err)
				return
			}
			fmt.Printf("client%v send message: hello\n", i)

			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				t.Error(err)
				return
			}
			fmt.Printf("client%v receive server reply: %s\n", i, string(buf[:n]))
		}(i)
	}

	wg.Wait()
}
