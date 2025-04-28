package _map

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrentReadMap(t *testing.T) {
	ConcurrentReadMap()
}

func TestConcurrentReadWriteMap(t *testing.T) {
	ConcurrentReadWriteMap()
}

func TestSyncMapUsage(t *testing.T) {
	var m sync.Map

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			m.Store(1, "one")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if v, ok := m.Load(1); ok {
				fmt.Println(v)
			}
		}
	}()

	wg.Wait()
}
