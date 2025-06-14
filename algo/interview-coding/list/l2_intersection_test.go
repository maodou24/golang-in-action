package list

import (
	"github.com/maodou24/algorithm-go/internel/list"
	"github.com/maodou24/algorithm-go/internel/utils"
	"testing"
)

func TestIntersection(t *testing.T) {
	tail := utils.NewList([]int{8, 4, 5})

	list1 := &list.Node{
		Val: 4,
		Next: &list.Node{
			Val:  1,
			Next: tail,
		},
	}

	list2 := &list.Node{
		Val: 5,
		Next: &list.Node{
			Val: 6,
			Next: &list.Node{
				Val:  1,
				Next: tail,
			},
		},
	}

	result := intersection(list1, list2)
	if result != tail {
		t.Fatal("not intersection")
	}
}
