package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueuePopPush(t *testing.T) {
	q := NewQueueWithStack()

	q.Push(1)
	q.Push(2)

	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, 2, q.Pop())
}
