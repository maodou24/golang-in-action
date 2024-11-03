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
	fmt.Println(v, ok)

	// 数据读取完了，返回无效值
	v1, ok1 := <-ch
	fmt.Println(v1, ok1)
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
