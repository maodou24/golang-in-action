package slice

import "fmt"

func AddEle(s []int, n int) []int {
	return append(s, n)
}

func ModifySliceEle(s []int) {
	s[0] = 55
}

func ArrayExceptFirstToSlice(s [4]int) []int {
	return s[1:]
}

func OperateSlice() {
	s := []int{1, 2, 3, 4, 5, 6}

	s1 := s[1:4]
	s2 := s[2:5]

	s1 = append(s1, 7)
	s2 = append(s2, 8)
	fmt.Println(s) // [1 2 3 4 7 8]

	s1[2] = 0

	fmt.Println(s)  // [1 2 3 0 7 8]
	fmt.Println(s1) //   [2 3 0 7]
	fmt.Println(s2) //     [3 0 7 8]
}
