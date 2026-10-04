package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	list := List{
		data: []int{1, 2, 3, 4, 5},
	}

	iterator := list.GetIterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
