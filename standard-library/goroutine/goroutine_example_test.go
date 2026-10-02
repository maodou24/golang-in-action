package _goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 两个goroutine交替打印AB
func TestGoroutineAlternatePrintAB(t *testing.T) {
	achan := make(chan string, 1)
	bchan := make(chan string, 1)

	achan <- "A"
	var wg sync.WaitGroup

	wg.Go(func() {
		for range 10 {
			fmt.Print(<-achan)
			bchan <- "B"
		}
	})

	wg.Go(func() {
		for range 10 {
			fmt.Print(<-bchan)
			achan <- " A"
		}
	})

	wg.Wait()
}

func TestGoroutineProducerConsumer(t *testing.T) {
	square := make(chan int, 10)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		// producer
		t := time.NewTicker(time.Second)
		var i int = 1
		for range t.C {
			square <- i
			if i >= 10 {
				break
			}
			i++
		}
		close(square)
	}()

	go func() {
		defer wg.Done()
		// consumer
		for v := range square {
			result := v * v
			fmt.Println(result)
		}
	}()

	wg.Wait()
}

// goroutine的panic未被捕获后，会一直向上游传递，直到程序崩溃
// panic not catch, the program dies
func TestGoroutineCrashNotCapturePanic(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		panic("test panic")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			fmt.Println("hello world")
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
