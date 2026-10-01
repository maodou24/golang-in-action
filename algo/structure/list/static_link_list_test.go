package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticLinkListInsert(t *testing.T) {
	head := NewStaticLinkList()

	head.Insert(1, 2)
	head.Insert(1, 3)
	head.Insert(2, 1)
	head.Insert(1, 4)

	assert.Equal(t, []int{4, 3, 1, 2}, head.show())
}

func TestStaticLinkListInsertSequence(t *testing.T) {
	head := NewStaticLinkList()

	head.Insert(1, 1)
	head.Insert(2, 2)
	head.Insert(3, 3)

	assert.Equal(t, []int{1, 2, 3}, head.show())
}

func newTestStaticLinkList123() *staticLinkList {
	head := NewStaticLinkList()

	head.Insert(1, 1)
	head.Insert(2, 2)
	head.Insert(3, 3)
	return head
}

func TestStaticLinkListDelete(t *testing.T) {
	head := newTestStaticLinkList123()

	assert.Equal(t, 1, head.Remove(1))
	assert.Equal(t, []int{2, 3}, head.show())

	assert.Equal(t, 3, head.Remove(2))
	assert.Equal(t, []int{2}, head.show())
}

func TestStaticLinkListDeleteInsert(t *testing.T) {
	head := newTestStaticLinkList123()

	assert.Equal(t, 1, head.Remove(1))
	assert.Equal(t, []int{2, 3}, head.show())

	assert.Equal(t, 3, head.Remove(2))
	assert.Equal(t, []int{2}, head.show())

	head.Insert(2, 4)
	head.Insert(1, 1)
	assert.Equal(t, []int{1, 2, 4}, head.show())
}
