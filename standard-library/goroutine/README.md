# goroutine

生产消费模型

## 控制协程的并发数量

### 使用缓冲channel作为信号量

利用缓冲channel阻塞原理实现控制并发数量

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 5) // buffer channel

	for i := range 100 {
		wg.Add(1)
		ch <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-ch }()

			fmt.Println(i)
			time.Sleep(time.Second)
		}()
	}
	wg.Wait()
}
```

## 练习题

### 两个goroutine交替打印AB