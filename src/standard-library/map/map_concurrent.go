package _map

import (
	"fmt"
	"sync"
	"time"
)

func ConcurrentReadMap() {
	var m = map[int]string{
		1: "one",
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			fmt.Println(m[1])
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			fmt.Println(m[1])
		}
	}()

	wg.Wait()
}

// one goroutine read, another goroutine write
func ConcurrentReadWriteMap() {
	var m = map[int]string{
		1: "one",
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			m[1] = "one"
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			_ = m[1]
		}
	}()

	wg.Wait()
}

// map支持并发读，读多写小的场景，适合用读写锁
func RWMapWithRWLock() {
	var m = map[int]string{
		1: "one",
	}

	var rwMutex sync.RWMutex

	go func() {
		ticker := time.NewTicker(time.Millisecond)
		for range ticker.C {
			rwMutex.Lock()
			m[1] = "write"
			rwMutex.Unlock()
		}
	}()

	for i := 0; i < 5; i++ {
		go func() {
			ticker := time.NewTicker(time.Microsecond * 300)
			for range ticker.C {
				rwMutex.RLock()
				_ = m[1] // read
				rwMutex.RUnlock()
			}
		}()
	}
}
