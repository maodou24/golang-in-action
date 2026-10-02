package slice

import (
	"fmt"
	"testing"
)

// array is passed by value, so the original array will not be modified.
func TestArrayUpdateInFunction(t *testing.T) {
	arr := [4]int{1, 2, 3, 4}
	func(arr [4]int) {
		arr[0] = 55
		fmt.Println(arr)
		// [55 2 3 4]
	}(arr)

	fmt.Println(arr)
	// [1 2 3 4]
	if arr != [4]int{1, 2, 3, 4} {
		t.Error("ModifyArrVal modify arr")
	}
}

// Array can be converted to slice by slicing operation.
func TestArrayConvertToSlice(t *testing.T) {
	arr := [4]int{1, 2, 3, 4}

	s := arr[1:]

	printSliceInfo(s)
	// [2 3 4]
	// len: 3, cap: 3
	if len(s) != 3 {
		t.Fatal("convert fail")
	}
}
