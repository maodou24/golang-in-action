package _map

import (
	"os"
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentReadMap(t *testing.T) {
	ConcurrentReadMap()
}

func TestConcurrentReadWriteMap(t *testing.T) {
	f, err := os.Create("panic.log")
	assert.Nil(t, err)
	err = debug.SetCrashOutput(f, debug.CrashOptions{})
	assert.Nil(t, err)
	assert.Panics(t, ConcurrentReadWriteMap)
}
