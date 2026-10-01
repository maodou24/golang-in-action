package queue

const queueSize = 10

type queue struct {
	data [queueSize]any

	front int
	rear  int
}

func NewQueue() *queue {
	return &queue{}
}

func (q *queue) EnQueue(a any) {
	if (q.rear+1)%10 == q.front {
		return
	}
	q.data[q.rear] = a
	q.rear = (q.rear + 1) % queueSize
}

func (q *queue) DeQueue() any {
	if q.rear == q.front {
		return nil
	}
	res := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % queueSize

	return res
}

type node struct {
	Val  any
	Next *node
}

type linkQueue struct {
	front *node
	rear  *node
}

func NewLinkQueue() *linkQueue {
	head := &node{}
	return &linkQueue{front: head, rear: head}
}

func (q *linkQueue) EnQueue(a any) {
	newNode := &node{Val: a}
	q.rear.Next = newNode

	q.rear = newNode
}

func (q *linkQueue) DeQueue() any {
	if q.front == q.rear {
		return nil
	}

	res := q.front.Next
	q.front.Next = res.Next

	if q.rear == res {
		q.rear = q.front
	}

	return res.Val
}
