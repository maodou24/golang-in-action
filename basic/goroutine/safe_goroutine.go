package goroutine

import (
	"log"
	"sync"
)

func SafeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recover goroutine panic: %v", r)
			}
		}()

		fn()
	}()
}

func SafeGoLoop(fn func()) {
	go func() {
		for {
			func() {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("recover goroutine panic: %v", r)
					}
				}()

				fn()
			}()

			// add some loop condition to avoid infinite loop
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
					log.Printf("recover goroutine panic: %v", r)
				}
			}()

			fn()
		}()
	}
}
