package _map

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// concurrent read map
func TestMapConcurrentRead(t *testing.T) {
	key := 1
	var m = map[int]string{
		key: "one",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_ = m[key]
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_ = m[key]
			}
		}
	}()

	wg.Wait()
}

// concurrent read and write map, will panic
// map not safe for concurrent use: read and write
func TestMapConcurrentReadWrite(t *testing.T) {
	f, err := os.Create("panic.log")
	assert.Nil(t, err)
	err = debug.SetCrashOutput(f, debug.CrashOptions{})
	assert.Nil(t, err)

	key := 1
	var m = map[int]string{
		key: "one",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				m[key] = "write"
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_ = m[key]
			}
		}
	}()

	wg.Wait()
}

// concurrent read and write map by RWMutex, safe for concurrent use
func TestMapConcurrentReadWriteRWLock(t *testing.T) {
	key := 1
	var m = map[int]string{
		key: "one",
	}
	var mu sync.RWMutex

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Go(func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				mu.Lock()
				m[key] = "write"
				mu.Unlock()
			}
		}
	})

	for range 5 {
		wg.Go(func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					mu.RLock()
					_ = m[key]
					mu.RUnlock()
				}
			}
		})
	}

	wg.Wait()
}

func TestMapKeySupportedDataType(t *testing.T) {
	ch1 := make(chan int)

	m := make(map[any]string)
	m[ch1] = "channel key"
	m[1] = "integer key"
	m["1"] = "string key"
	m[struct{ a int }{a: 1}] = "struct key"
	m[true] = "bool key"
	m[[3]int{1, 2, 3}] = "array key"

	s := []int{1, 2, 3}
	m[s] = "slice key" // slice is not comparable, cannot be used as map key

	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

func TestMapKeyNotSupportedDataType(t *testing.T) {
	m := make(map[any]string)

	s := []int{1, 2, 3}
	m[s] = "slice key" // unhashable type cannot be used as map key

	mkey := struct{ m map[int]int }{m: map[int]int{1: 1}}
	m[mkey] = "map key" // unhashable type cannot be used as map key

	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

// delete key during iteration, will not panic,
// but the behavior is undefined.
// It's recommended to collect the keys and delete them after the iteration.
func TestMapDeleteKeyDuringIteration(t *testing.T) {
	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		m[i] = i
	}

	for k := range m {
		if k%2 == 0 {
			delete(m, k)
		}
	}

	fmt.Println(m)
}
