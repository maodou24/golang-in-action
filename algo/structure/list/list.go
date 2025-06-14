package list

type Node struct {
	Val  int
	Next *Node
}

type DoublyNode struct {
	Val  int
	Pre  *DoublyNode
	Next *DoublyNode
}

func Reverse(head *Node) *Node {
	cur := head
	var pre *Node

	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
	return pre
}

func ReverseByRecursion(head *Node) *Node {
	var recur func(cur, pre *Node) *Node

	recur = func(cur, pre *Node) *Node {
		if cur == nil {
			return pre
		}

		ans := recur(cur.Next, cur)
		cur.Next = pre
		return ans
	}

	return recur(head, nil)
}
