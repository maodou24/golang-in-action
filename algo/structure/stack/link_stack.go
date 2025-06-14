package stack

type Node struct {
	Val  any
	Next *Node
}

type LinkStack struct {
	top *Node
	n   int
}

func NewLinkStack() *LinkStack {
	return &LinkStack{}
}

func (l *LinkStack) Push(a any) {
	top := &Node{Val: a, Next: l.top}

	l.top = top
	l.n++
}

func (l *LinkStack) Pop() any {
	if l.top == nil {
		return nil
	}
	res := l.top.Val
	l.top = l.top.Next
	return res
}
