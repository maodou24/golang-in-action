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