package slice

import (
	"testing"
)

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
