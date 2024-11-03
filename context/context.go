package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func TaskDone() {
	ctx := context.Background()

	ctx1, cancel := context.WithCancel(ctx)
	n := 5

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()

			for {
				time.Sleep(100 * time.Millisecond)
				select {
				case <-ctx1.Done():
					fmt.Printf("task[%v] cancel\n", i)
					return
				default:
					fmt.Printf("task[%v] doing work\n", i)
					if i == 4 {

					}
				}
			}
		}(i)
	}

	time.Sleep(1 * time.Second)
	// 当其中一个子goroutine任务结束后，需要通知主线程执行cancel操作
	cancel()
	wg.Wait()
}
