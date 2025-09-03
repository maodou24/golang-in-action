package slice

import (
	"fmt"
	"time"
)

type A struct {
	N int
}

func (a A) String() string {
	return fmt.Sprintf("%d", a.N)
}

type B struct {
	a      A
	aPoint *A
}

func (b B) String() string {
	return fmt.Sprintf("{a: %v, aPoint: %v}", b.a, b.aPoint)
}

func IterateSliceInClose() {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	for _, v := range s {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(1 * time.Second)
}
