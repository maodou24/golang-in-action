package slice

import "fmt"

func printSliceInfo[T any](s []T) {
	fmt.Printf("%v, len: %v, cap: %v\n", s, len(s), cap(s))
	fmt.Printf("point: %p, %p\n", s, &s)
}
