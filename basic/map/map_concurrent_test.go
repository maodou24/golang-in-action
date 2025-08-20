package _map

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentReadMap(t *testing.T) {
	ConcurrentReadMap()
}

func TestConcurrentReadWriteMap(t *testing.T) {
	assert.Panics(t, ConcurrentReadWriteMap)
}
