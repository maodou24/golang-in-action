package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPrintAB(t *testing.T) {
	PrintAB()
}

func TestProducerConsumer(t *testing.T) {
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

func TestGoWithoutRecover(t *testing.T) {
	GoWithoutRecover()
}
