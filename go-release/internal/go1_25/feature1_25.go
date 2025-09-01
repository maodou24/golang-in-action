package go1_25

import (
	"fmt"
	"sync"
)

func WaitGroupGo() {
	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Go(func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	})

	wg.Go(func() {
		for v := range ch {
			fmt.Println(v * v)
		}
	})

	wg.Wait()
}

func IterateSliceInClose() {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	for v := range s {
		defer func() {
			fmt.Println(v)
		}()
	}
}
