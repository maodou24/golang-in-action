package _map

import (
	"fmt"
	"sync"
)

func ConcurrentReadMap() {
	var m = map[int]string{
		1: "one",
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			fmt.Println(m[1])
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for {
			fmt.Println(m[1])
		}
		wg.Done()
	}()

	wg.Wait()
}

func ConcurrentReadWriteMap() {
	var m = map[int]string{
		1: "one",
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			m[1] = "one"
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for {
			fmt.Println(m[1])
		}
		wg.Done()
	}()

	wg.Wait()
}
