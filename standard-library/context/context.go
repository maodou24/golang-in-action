package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 多个协程处理任务，只要其中一个处理完成，结束所有任务
func ParllalTaskDone() {
	ctx := context.Background()

	ctx1, cancel := context.WithCancel(ctx)
	defer cancel()
	n := 5

	exit := make(chan int)

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()

			var n int
			for {
				time.Sleep(100 * time.Millisecond)
				select {
				case <-ctx1.Done():
					fmt.Printf("task[%v] cancel\n", i)
					return
				default:
					fmt.Printf("task[%v] doing work\n", i)
					if i == 4 && n == 5 {
						// 模拟第一个协程任务完成
						exit <- 0
					}
					n++
				}
			}
		}(i)
	}

	for range exit {
		close(exit)
	}

	wg.Wait()
	fmt.Println("all task done")
}

// 协程任务设置定时器，超过定时器设置时间没有处理完成，退出任务
func TaskSetTimer() {
	timerCtx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-timerCtx.Done():
				fmt.Println("The task execution exceeded the time limit")
				return // exit task
			default: 
				fmt.Println("task doing")
			}

			time.Sleep(time.Second / 5)
		}
	}()

	wg.Wait()
}