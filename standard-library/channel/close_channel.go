package channel

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// don't close a channel from the receiver side and don't close a channel
// if the channel has multiple concurrent senders

// case 1: one sender, one receiver
// case 2: one sender, multiple receivers
// case 3: multiple senders, one receiver
// case 4: multiple senders, multiple receivers

// one sender, one or multiple receivers, colse channel by sender
func CloseChannelCase1AndCase2() {
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

// case3: multiple senders and one receiver
func CloseChannelCase3() {
	ch := make(chan int, 10)

	exit := make(chan struct{})

	var wg sync.WaitGroup
	// multiple senders
	const senders = 4
	for i := 0; i < senders; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				select {
				case <-exit:
					return
				default:
					ch <- i
					time.Sleep(time.Second) // every second generate one
				}
			}
		}()
	}

	// one receiver
	wg.Add(1)
	go func() {
		defer wg.Done()

		var count int
		const MAX = 10
		for n := range ch {
			if count > MAX {
				fmt.Println("notify sender")
				close(exit)
				return
			}
			count++
			fmt.Println(count, n) // consume
		}
	}()

	// ch is not invoke close function, it will close by gc, no goroutine refer this channel
	wg.Wait()
}

// case4: multiple senders, multiple receivers
func CloseChannelCase4() {
	const receivers = 5
	const senders = 5
	const MaxSendCount = 100

	ch := make(chan int, 10)

	exit := make(chan struct{})

	var wg sync.WaitGroup
	exitBy := make(chan string, 1)
	go func() {
		g := <-exitBy
		fmt.Println(g)
		close(exit)
	}()

	for i := 0; i < senders; i++ {
		wg.Add(1)

		var count int32
		go func(id int) {
			defer wg.Done()

			if atomic.LoadInt32(&count) > MaxSendCount {
				select {
				case exitBy <- fmt.Sprintf("sender %v exit", id):
				default:
				}
				return
			}
			for {
				select {
				case <-exit:
					fmt.Printf("sender %v exit\n", id)
					return
				default:
					ch <- i
					atomic.AddInt32(&count, 1)
					time.Sleep(time.Second) // every second generate one
				}
			}
		}(i)
	}

	// receivers
	for i := 0; i < receivers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			var count int
			const MAX = 10

			for {
				select {
				case <-exit:
					return
				case n := <-ch:
					if count > MAX {
						select {
						case exitBy <- fmt.Sprintf("receiver %v exit", id):
						default:
						}
						return
					}
					count++
					fmt.Printf("receiver #%d receive %d\n", id, n)
				}
			}

		}(i)
	}

	wg.Wait()
}
