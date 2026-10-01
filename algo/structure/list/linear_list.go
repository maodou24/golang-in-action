package list

// List顺序存储结构。golang中数组
type LinearList struct {
	data [100]int
	len  int
}

func NewLinearList() *LinearList {
	return &LinearList{}
}

func (l *LinearList) Get(n int) int {
	if n < 0 || n >= l.len {
		panic("out index")
	}
	return l.data[n]
}

func (l *LinearList) Insert(n int, v int) {
	if n < 0 || n > l.len || l.full() {
		return
	}
	for i := l.len - 1; i >= n; i-- {
		l.data[i+1] = l.data[i]
	}
	l.data[n] = v
	l.len++
}

func (l *LinearList) Remove(n int) int {
	if n < 0 || n >= l.len {
		panic("out index")
	}
	result := l.data[n]
	for i := n; i < l.len; i++ {
		l.data[i] = l.data[i+1]
	}
	l.data[l.len-1] = 0
	l.len--
	return result
}

func (l *LinearList) RemoveValue(v int) int {
	index := -1
	for i := 0; i < l.len; i++ {
		if l.data[i] == v {
			index = i
		}
	}
	return l.Remove(index)
}

func (l *LinearList) Size() int {
	if l != nil {
		return l.len
	}
	return 0
}

func (l *LinearList) Elements() []int {
	result := make([]int, l.len)
	for i := range result {
		result[i] = l.data[i]
	}
	return result
}

func (l *LinearList) full() bool {
	return l.len == 100
}
