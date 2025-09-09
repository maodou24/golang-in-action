package _goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	workers := NewWorkerPool(5)

	for i := range 100 {
		workers.Submit(func() error {
			fmt.Println("worker", i)
			time.Sleep(1 * time.Second)
			return nil
		})
	}

	workers.Start()
}
