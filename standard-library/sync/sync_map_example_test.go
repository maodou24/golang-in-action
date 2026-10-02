package sync

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncMap(t *testing.T) {
	var m sync.Map

	m.Store("key1", "value1")
	m.Store("key2", "value2")
	// {key1: value1, key2: value2}

	value, ok := m.Load("key1")
	if !ok || value != "value1" {
		t.Errorf("expected value1, got %s", value)
	}

	value2, ok2 := m.Load("key2")
	if !ok2 || value2 != "value2" {
		t.Errorf("expected value2, got %s", value2)
	}

	// get all keys and values
	var keys []string
	var values []string
	m.Range(func(k, v any) bool {
		if str, ok := k.(string); ok {
			keys = append(keys, str)
		}
		if str, ok := v.(string); ok {
			values = append(values, str)
		}
		return true
	})
	fmt.Printf("keys: %v\n", keys)     // [key1 key2]
	fmt.Printf("values: %v\n", values) // [value1 value2]

	m.Store("key1", "value1-1")
	// {key1: value1-1, key2: value2}

	m.Delete("key1")
	// {key2: value2}
}

// sync.Map is safe for concurrent use
func TestSyncMapConcurrentReadWrite(t *testing.T) {
	var m sync.Map

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
				m.Store("key1", time.Now().String())
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
				m.Load("key1")
			}
		}
	}()

	wg.Wait()
}
