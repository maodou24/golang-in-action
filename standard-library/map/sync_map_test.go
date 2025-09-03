package _map

import (
	"sync"
	"testing"
	"time"
)

func TestSyncMapUsageFunc(t *testing.T) {
	syncMap := &syncMapUsage{}

	syncMap.Put("key1", "value1")
	syncMap.Put("key2", "value2")

	value, ok := syncMap.Get("key1")
	if !ok || value != "value1" {
		t.Errorf("expected value1, got %s", value)
	}

	value2, ok2 := syncMap.Get("key2")
	if !ok2 || value2 != "value2" {
		t.Errorf("expected value2, got %s", value2)
	}

	keys := syncMap.Keys()
	values := syncMap.Values()
	if keys[0] != "key1" || keys[1] != "key2" || values[0] != "value1" || values[1] != "value2" {
		t.Errorf("expected keys: [key1, key2], values: [value1, value2]")
	}

	syncMap.Delete("key1")
	deleted := syncMap.DeleteValue("key2", "value2")
	if !deleted {
		t.Errorf("expected deleted is true, got false")
	}
	syncMap.Clear()
}

func TestSyncMapUsageConcurrent(t *testing.T) {
	syncMap := &syncMapUsage{}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			syncMap.Put("key1", time.Now().String())
			time.Sleep(time.Second / 1000000)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			syncMap.Get("key1")
			time.Sleep(time.Second / 1000000)
		}
	}()

	wg.Wait()
}
