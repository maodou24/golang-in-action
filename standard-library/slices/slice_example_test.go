package slice

import (
	"fmt"
	"testing"
)

// append slice overflow capacity, the slice will be re-capability,
// and the new slice will be returned
func TestSliceAppendEleReCapability(t *testing.T) {
	s := []int{1, 2, 3}

	printSliceInfo(s)
	// len: 3, cap: 3

	newS := append(s, 4) // append will re-capability the slice

	printSliceInfo(newS)
	// len: 4, cap: 6
}

// append slice under capacity, lower level array not changed
func TestSliceAppendEleInitCapability(t *testing.T) {
	s := make([]int, 0, 3)

	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)

	printSliceInfo(s)
	// len: 3, cap: 3
}

func TestSliceUpdateIndexEleInFunction(t *testing.T) {
	s := []int{1, 2, 3}

	// change ele in function, the slice will be changed
	func(s []int) {
		s[0] = 55
	}(s)

	fmt.Println(s) // [55 2 3]
	if s[0] != 55 {
		t.Fatal("TestSliceUpdateIndexEleInFunction first ele failed")
	}
}

func TestSliceBaseOperate(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}

	printSliceInfo(s)

	s1 := s[1:4] // 切片操作 re-slicing
	s2 := s[2:5] // 切片操作 re-slicing

	s1 = append(s1, 7)
	s2 = append(s2, 8)

	printSliceInfo(s1)
	printSliceInfo(s2)

	s1[2] = 0

	printSliceInfo(s)
	printSliceInfo(s1)
	printSliceInfo(s2)
}

// 函数中对 slice 进行 append 操作时
// 在slice 的容量足够时，需要将 append 的结果重新赋值给 slice 变量，否则 slice 变量不会发生变化
func TestSliceAppendSliceInFunc(t *testing.T) {
	s := make([]int, 0, 2)
	printSliceInfo(s)
	// s len: 0, cap: 2
	func(s []int) {
		s = append(s, 1)
		printSliceInfo(s)
	}(s)

	// s slice info is not changed, but the lower level array is changed
	printSliceInfo(s)
}

func TestSliceUpdateIndexEleAfterReCapInFunction(t *testing.T) {
	s := []int{1, 2, 3}

	printSliceInfo(s)
	func(s []int) { // recap slice and change the first element
		s = append(s, 10)
		s[0] = 55

		printSliceInfo(s)
	}(s)

	printSliceInfo(s)

	if s[0] == 55 {
		t.Fatal("TestSliceUpdateIndexEleAfterReCapInFunction failed")
	}
}
