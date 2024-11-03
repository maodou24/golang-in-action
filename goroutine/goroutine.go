package goroutine

import (
	"fmt"
	"sync"
)

// 交替打印AB 10次
func PrintAB() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			_ = <-ch1
			fmt.Print("A")
			ch2 <- 1
		}
	}()
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			_ = <-ch2
			fmt.Print("B")
			ch1 <- 1
		}
	}()

	ch1 <- 1

	wg.Wait()
	close(ch1)
	close(ch2)
}
