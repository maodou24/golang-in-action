package channel

import (
	"fmt"
	"sync"
)

// don't close a channel from the receiver side and don't close a channel
// if the channel has multiple concurrent senders

// one sender, and multiple receivers
func CloseChannelDemo1() {
	ch := make(chan int, 10)

	var wg sync.WaitGroup
	// sender
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}

		close(ch)
	}()

	const receivers = 4
	for i := 0; i < receivers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range ch {
				fmt.Println(n)
			}
		}()
	}

	wg.Wait()
}
