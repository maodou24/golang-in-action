package iterator

type Iterator interface {
	Next() any
	HasNext() bool
}

type List struct {
	data []int
	i    int
}

func (l *List) Next() any {
	r := l.data[l.i]
	l.i++
	return r
}

func (l *List) HasNext() bool {
	return l.i < len(l.data)
}

func (l *List) GetIterator() Iterator {
	return l
}
