package defer_example

import "fmt"

type S []int

func (s *S) add(n int) *S {
	fmt.Println(s, n)
	*s = append(*s, n)
	return s
}
