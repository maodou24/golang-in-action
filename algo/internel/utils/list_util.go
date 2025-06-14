package utils

import "github.com/maodou24/algorithm-go/internel/list"

func NewList(s []int) *list.Node {
	head := &list.Node{}

	node := head
	for _, v := range s {
		node.Next = &list.Node{Val: v}

		node = node.Next
	}
	return head.Next
}
