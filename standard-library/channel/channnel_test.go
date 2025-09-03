package channel

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReadBufferClosedChan(t *testing.T) {
	ReadBufferClosedChan()
}

func TestReadChanNotBlock(t *testing.T) {
	ch := make(chan int, 1)

	ch <- 1
	ReadChanNotBlock(ch)
	// 前面已经消费掉消息
	ReadChanNotBlock(ch)
}

func TestReadChanNotBlockWithTimer(t *testing.T) {
	ch := make(chan int, 1)

	ch <- 1
	timer := time.NewTimer(time.Second)
	ReadChanNotBlockWithTimer(ch, timer)
	// 前面已经消费掉消息
	ReadChanNotBlockWithTimer(ch, timer)
}

func TestCloseChannelMulti(t *testing.T) {
	ch := make(chan int, 1)

	assert.Panics(t, func() {
		CloseChannelMuti(ch)
	})
}

func TestSendToCloseChannel(t *testing.T) {
	ch := make(chan int, 1)

	assert.Panics(t, func() {
		SendToCloseChannel(ch)
	})
}

func TestSafeCloseRudely(t *testing.T) {
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

func TestSafeSendRudely(t *testing.T) {
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
