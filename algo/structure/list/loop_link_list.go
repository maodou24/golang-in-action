package list

func NewLoopList() *linkList {
	head := &Node{}
	head.Next = head
	return &linkList{Head: head}
}
