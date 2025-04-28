package slice

import "fmt"

type A struct {
	N int
}

func (a A) String() string {
	return fmt.Sprintf("%d", a.N)
}

type B struct {
	a A
	aPoint *A
}

func (b B) String() string {
	return fmt.Sprintf("{a: %v, aPoint: %v}", b.a, b.aPoint)
}