package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	funcs := []func(head *Node) *Node{
		Reverse,
		ReverseByRecursion,
	}

	for _, fc := range funcs {
		data := IntSliceToListNode([]int{1, 2, 3, 4, 5})

		got := fc(data)

		assert.EqualValues(t, IntSliceToListNode([]int{5, 4, 3, 2, 1}), got)
	}
}
