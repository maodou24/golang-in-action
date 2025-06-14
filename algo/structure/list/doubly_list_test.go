package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyListLen(t *testing.T) {
	list := NewDoublyList()

	list.PushFront(1)
	list.PushFront(2)
	list.PushBack(3)
	list.PushFront(4)
	list.PushBack(5)

	assert.Equal(t, 5, list.Len())
}

func TestDoublyListPushFrontAndPushBack(t *testing.T) {
	list := NewDoublyList()

	list.PushFront(1)

	assert.Equal(t, 1, list.Len())
	assert.Equal(t, 1, list.Back().Value)
	assert.Equal(t, 1, list.Front().Value)


	list.PushFront(2)
	assert.Equal(t, 2, list.Len())
	assert.Equal(t, 1, list.Back().Value)
	assert.Equal(t, 2, list.Front().Value)

	list.PushBack(3)
	assert.Equal(t, 3, list.Len())
	assert.Equal(t, 3, list.Back().Value)
	assert.Equal(t, 2, list.Front().Value)
}

func TestDoublyListRemove(t *testing.T) {
	list := NewDoublyList()

	list.PushFront(1)
	list.PushFront(2)
	list.PushBack(3)
	ele4 := list.PushFront(4)
	list.PushBack(5)
	
	list.Remove(ele4)
	except := []int{2, 1, 3, 5}
	i := 0
	for node := list.root.next; node != &list.root; node = node.next {
		assert.Equal(t, except[i], node.Value)
		i++
	}
}