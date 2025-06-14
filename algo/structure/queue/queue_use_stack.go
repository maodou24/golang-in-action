package queue

import "github.com/maodou24/algorithm-go/structure/stack"

func NewQueueWithStack() *queueWithStack {
	return &queueWithStack{A: stack.NewLinearStack(), B: stack.NewLinearStack()}
}

// 两个栈实现队列
type queueWithStack struct {
	A stack.Stack
	B stack.Stack
}

func (q *queueWithStack) Push(n int) {
	q.A.Push(n)
}

func (q *queueWithStack) Pop() int {
	if len(q.B.Show()) != 0 {
		return q.B.Pop()
	}

	if len(q.A.Show()) == 0 {
		return -1
	}

	for len(q.A.Show()) != 0 {
		q.B.Push(q.A.Pop())
	}
	return q.B.Pop()
}
