package _defer

import "fmt"

type S []int

func (s *S) add(n int) *S {
	fmt.Println(s, n)
	*s = append(*s, n)
	return s
}

func DeferFuncList() {
	s := new(S)
	defer s.add(1).add(2).add(3)
}
