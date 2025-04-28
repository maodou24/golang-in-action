package channel

import (
	"fmt"
	"time"
)

func ReadBufferClosedChan() {
	ch := make(chan int, 2)
	ch <- 1

	close(ch)

	// 缓存通道即使关闭了，也可以读取
	v, ok := <-ch
	fmt.Println(v, ok) // 1, true

	// 数据读取完了，返回无效值
	v1, ok1 := <-ch
	fmt.Println(v1, ok1) // 0, false
}

// 读取channel不阻塞
func ReadChanNotBlock(ch chan int) {
	select {
	case v := <-ch:
		fmt.Println("read value is", v)
	default:
		fmt.Println("no data in chan")
	}
}

// 直到指定时间后，读取channel不阻塞
func ReadChanNotBlockWithTimer(ch chan int, timer *time.Timer) {
	select {
	case v := <-ch:
		fmt.Println("read value is", v)
	case <-timer.C:
		fmt.Println("wait chan data timeout")
	}
}

func CloseChannelMuti(ch chan int) {
	close(ch)

	close(ch) // Close a closed channel will cause panic
}

func SendToCloseChannel(ch chan int) {
	close(ch)

	ch <- 1 // sent to a closed channel will cause panic
}