package _goroutine

import (
	"fmt"
	"sync"
	"time"
)

// 交替打印AB 10次
func PrintAB() {
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			_ = <-ch1
			fmt.Print("A")
			ch2 <- struct{}{}
		}
	}()
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			_ = <-ch2
			fmt.Print("B")
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}

	wg.Wait()
	close(ch1)
	close(ch2)
}

// panic not catch, the program dies
func GoWithoutRecover() {
	go func() {
		panic("test panic")
	}()

	go func() {
		for {
			fmt.Println("hello world")
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 10)
}
