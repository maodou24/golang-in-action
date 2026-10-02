package _goroutine

import (
	"fmt"
	"sync"
)

func SafeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recover goroutine panic: %v\n", r)
			}
		}()

		fn()
	}()
}

func SafeGoLoop(fn func()) {
	go func() {
		var n int
		for {
			func() {
				defer func() {
					if r := recover(); r != nil {
						fmt.Printf("recover goroutine panic: %v\n", r)
					}
				}()

				fn()
			}()

			// add some loop condition to avoid infinite loop
			if n > 5 {
				break
			}
			n++
		}
	}()
}

func SafeGoFuncs(funcs ...func()) {
	var wg sync.WaitGroup
	for _, fn := range funcs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("recover goroutine panic: %v\n", r)
				}
			}()

			fn()
		}()
	}
}
