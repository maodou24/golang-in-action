package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkListInsert(t *testing.T) {
	head := NewLinkList()

	head.Insert(0, 2)
	head.Insert(0, 3)
	head.Insert(2, 1)

	assert.Equal(t, &linkList{
		size: 3,
		Head: &Node{
			Val: 0,
			Next: &Node{
				Val: 3,
				Next: &Node{
					Val: 2,
					Next: &Node{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}, head)
}

func TestLinkListPushFront(t *testing.T) {
	head := NewLinkList()

	head.PushFront(1)
	head.PushFront(2)
	head.PushFront(3)

	assert.Equal(t, head.Size(), 3)
	assert.Equal(t, []int{3, 2, 1}, head.show())
}

func newTestDataLinkList123() *linkList {
	head := NewLinkList()
	head.PushBack(1)
	head.PushBack(2)
	head.PushBack(3)
	return head
}

func TestLinkListPopFront(t *testing.T) {
	head := newTestDataLinkList123()

	assert.Equal(t, 1, head.PopFront())
	assert.Equal(t, 2, head.PopFront())
	assert.Equal(t, 3, head.PopFront())
	assert.Panics(t, func() {
		head.PopFront()
	})
}

func TestLinkListPopBack(t *testing.T) {
	head := newTestDataLinkList123()

	assert.Equal(t, 3, head.PopBack())
	assert.Equal(t, 2, head.PopBack())
	assert.Equal(t, 1, head.PopBack())
	assert.True(t, head.IsEmpty())
	assert.Panics(t, func() {
		head.PopFront()
	})
}

func TestLinkListRemove(t *testing.T) {
	head := newTestDataLinkList123()

	assert.Equal(t, 2, head.Remove(1))
	assert.Equal(t, 3, head.Remove(1))

	assert.Equal(t, []int{1}, head.show())
}

func TestLinkListValueAtIndex(t *testing.T) {
	head := newTestDataLinkList123()

	assert.Equal(t, 2, head.ValueAtIndex(1))
	assert.Equal(t, 3, head.ValueAtIndex(2))
}

func TestLinkListRemoveValue(t *testing.T) {
	head := newTestDataLinkList123()

	head.RemoveValue(2)
	assert.Equal(t, []int{1, 3}, head.show())

	head.RemoveValue(1)
	assert.Equal(t, []int{3}, head.show())
}

func TestLinkListReverse(t *testing.T) {
	head := newTestDataLinkList123()

	assert.Equal(t, []int{3, 2, 1}, head.Reverse().show())
}

func TestLinkListBack(t *testing.T) {
	head := newTestDataLinkList123()

	assert.Equal(t, head.Back(1), 3)
	assert.Equal(t, head.Back(3), 1)
}
