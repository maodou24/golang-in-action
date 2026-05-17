package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinearList(t *testing.T) {
	list := NewLinearList()
	list.Insert(0, 1)
	list.Insert(0, 2)
	list.Insert(1, 3)

	assert.Equal(t, 3, list.Size())
	assert.Equal(t, []int{2, 3, 1}, list.Elements())

	list.Remove(2)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, []int{2, 3}, list.Elements())

	list.Remove(0)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, []int{3}, list.Elements())
}
