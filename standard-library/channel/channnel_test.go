package channel

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChannelReadBufferClosedChan(t *testing.T) {
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

func TestChannelReadNotBlock(t *testing.T) {
	ch := make(chan int, 1)

	ch <- 1

	<-ch // 消费

	select {
	case v := <-ch:
		fmt.Println("read value is", v)
	default:
		fmt.Println("no data in chan")
	}
}

func TestChannelReadChanNotBlockWithTimer(t *testing.T) {
	ch := make(chan int, 1)

	ch <- 1

	<-ch // 消费

	// 阻塞1s
	timer := time.NewTimer(time.Second)
	select {
	case v := <-ch:
		fmt.Println("read value is", v)
	case <-timer.C:
		fmt.Println("wait chan data timeout")
	}
}

// close a closed channel will panic
func TestChannelMultiClose(t *testing.T) {
	ch := make(chan int, 1)

	assert.Panics(t, func() {
		close(ch)

		close(ch) // Close a closed channel will cause panic
	})
}

// send to a closed channel will panic
func TestChannelSendToClosed(t *testing.T) {
	ch := make(chan int, 1)

	assert.Panics(t, func() {
		close(ch)

		ch <- 1 // sent to a closed channel will cause panic
	})
}

func TestChannelSafeCloseRudely(t *testing.T) {
	ch := make(chan any, 1)
	close(ch)

	assert.NotPanics(t, func() {
		closed := SafeCloseRudely(ch)
		assert.False(t, closed)
	})

	ch2 := make(chan any, 1)
	assert.NotPanics(t, func() {
		closed := SafeCloseRudely(ch2)
		assert.True(t, closed)
		closed = SafeCloseRudely(ch2)
		assert.False(t, closed)
	})
}

func TestChannelSafeSendRudely(t *testing.T) {
	ch := make(chan any, 1)
	close(ch)

	assert.NotPanics(t, func() {
		closed := SafeSendRudely(ch, 1)
		assert.True(t, closed)
	})

	ch2 := make(chan any, 1)
	assert.NotPanics(t, func() {
		closed := SafeSendRudely(ch2, 1)
		assert.False(t, closed)
	})
}

func TestSafeCloser(t *testing.T) {
	ch := make(chan any, 1)

	closer := NewSafeCloser(ch)

	assert.NotPanics(t, func() {
		closer.Close()
		closer.Close()
	})
}
