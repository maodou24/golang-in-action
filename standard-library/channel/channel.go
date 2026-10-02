package channel

import (
	"sync"
)

func SafeCloseRudely(ch chan any) (closeSuccess bool) {
	defer func() {
		if recover() != nil {
			closeSuccess = false
		}
	}()
	close(ch)
	return true // chan close success
}

func SafeSendRudely(ch chan any, value any) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- value
	return false // send chan success, channel is not closed
}

type SafeCloser struct {
	ch chan any
	sync.Once
}

func NewSafeCloser(ch chan any) *SafeCloser {
	return &SafeCloser{ch: ch}
}

func (s *SafeCloser) Close() {
	s.Do(func() {
		close(s.ch)
	})
}

type SafeCloserMutex struct {
	ch    chan any
	close bool
	sync.Mutex
}

func NewSafeCloserMutex(ch chan any) *SafeCloserMutex {
	return &SafeCloserMutex{ch: ch}
}

func (s *SafeCloserMutex) Close() {
	s.Lock()
	defer s.Unlock()
	if s.close {
		return
	}
	close(s.ch)
	s.close = true
}
