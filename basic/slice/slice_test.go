package slice

import (
	"fmt"
	"testing"
)

func TestAddEleWithReCapability(t *testing.T) {
	s := []int{1, 2, 3}

	newS := AddEle(s, 4)

	fmt.Println(s) // [1 2 3]
	if len(newS) == len(s) {
		t.Fatal("AddEle can't add ele with re-capability")
	}
}

func TestAddEleWithoutReCapability(t *testing.T) {
	s := make([]int, 0, 3)

	s = AddEle(s, 1)
	s = AddEle(s, 2)
	s = AddEle(s, 3)

	fmt.Println(s) // [1 2 3]
	if len(s) != 3 {
		t.Fatal("AddEle can add ele without re-capability")
	}
}

func TestModifySliceEle(t *testing.T) {
	s := []int{1, 2, 3}

	ModifySliceEle(s)

	fmt.Println(s) // [55 2 3]
	if s[0] != 55 {
		t.Fatal("ModifySliceEle first ele failed")
	}
}

func TestArrayExceptFirstToSlice(t *testing.T) {
	arr := [4]int{1, 2, 3, 4}

	s := ArrayExceptFirstToSlice(arr)

	fmt.Println(s) // [55 2 3]
	if len(s) != 3 {
		t.Fatal("convert fail")
	}
}

func TestOperateSlice(t *testing.T) {
	OperateSlice()
}

func TestAppendSlice(t *testing.T) {
	s := make([]int, 0, 2)
	AppendSlice(s, 1)
	if len(s) != 0 {
		t.Fatal("AppendSlice can't affect slice")
	}
}

func TestModifySliceWithReCap(t *testing.T) {
	s := []int{1, 2, 3}
	ModifySliceWithReCap(s)

	if s[0] == 55 {
		t.Fatal("ModifySliceWithReCap failed")
	}
}

func TestFuncAppendSliceWithReCap(t *testing.T) {
	s := []int{1, 2, 3}
	fmt.Println("-- func out before append --")
	fmt.Printf("%v\n", s)
	fmt.Printf("%v, len: %v, cap: %v\n", s, len(s), cap(s))
	fmt.Printf("point: %p, %p\n", s, &s)
	AppendSlice(s, 1)

	fmt.Println("-- func out after append --")
	fmt.Printf("%v, len: %v, cap: %v\n", s, len(s), cap(s))
	fmt.Printf("%v\n", s)
	fmt.Printf("point: %p, %p\n", s, &s)
}

func TestFuncAppendSliceWithoutReCap(t *testing.T) {
	s := make([]int, 0, 2)
	fmt.Println("-- func out before append --")
	fmt.Printf("%v, len: %v, cap: %v\n", s, len(s), cap(s))
	fmt.Printf("point: %p, %p\n", s, &s)
	AppendSlice(s, 1)

	fmt.Println("-- func out after append --")
	fmt.Printf("%v, len: %v, cap: %v\n", s, len(s), cap(s))
	fmt.Printf("point: %p, %p\n", s, &s)
}
