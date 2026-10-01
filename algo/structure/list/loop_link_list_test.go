package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoopListInsert(t *testing.T) {
	head := NewLoopList()

	head.Insert(0, 2)
	head.Insert(0, 3)

	assert.Equal(t, []int{3,2}, head.show())
}
