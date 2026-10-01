package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Linear(t *testing.T) {
	stack := NewLinearStack()

	stack.Push(1)
	stack.Push(2)

	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, 0, stack.Pop())
	assert.Equal(t, []int{}, stack.Show())
}
