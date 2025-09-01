package _panic

import (
	"fmt"
	"testing"
)

func TestPanicNilPointer(t *testing.T) {
	defer func() {
		err := recover()
		fmt.Printf("recover goroutine panic: %v\n", err)
	}()
	panic(nil)
}
