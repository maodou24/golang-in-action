package _goroutine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSafeGoroutine(t *testing.T) {
	assert.NotPanics(t, func() {
		SafeGo(func() {
			panic("test panic")
		})
	})
}

func TestSafeGoroutineLoop(t *testing.T) {
	assert.NotPanics(t, func() {
		SafeGoLoop(func() {
			panic("test panic")
		})
	})
}

func TestSafeGoFuncs(t *testing.T) {
	assert.NotPanics(t, func() {
		SafeGoFuncs(
			func() { fmt.Println("hello world") },
			func() { panic("test panic") })
	})
}
