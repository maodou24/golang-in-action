package stack

const maxSize = 100

type LinearStack struct {
	data [maxSize]int
	top  int
}

func NewLinearStack() Stack {
	return &LinearStack{top: -1}
}

func (l *LinearStack) Push(v int) {
	if l.top >= maxSize-1 {
		return
	}
	l.top++
	l.data[l.top] = v
}

func (l *LinearStack) Pop() int {
	if l.top < 0 {
		return 0
	}

	ans := l.data[l.top]
	l.top--
	return ans
}

func (l *LinearStack) Show() []int {
	return l.data[:l.top+1]
}
