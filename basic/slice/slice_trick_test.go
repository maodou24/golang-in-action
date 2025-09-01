package slice

import (
	"fmt"
	"testing"
)

func TestTrickEleIdxPointer(t *testing.T) {
	s := []B{
		{a: A{N: 1}},
		{a: A{N: 2}},
		{a: A{N: 3}},
	}

	for i := range s {
		s[i].aPoint = &s[i].a
	}

	fmt.Println("s before operation is", s) // s before operation is [{a: 1, aPoint: 1} {a: 2, aPoint: 2} {a: 3, aPoint: 3}]
	// remove the first element
	s = append(s[:0], s[1:]...)
	fmt.Println("s after operation is", s) // s after operation is [{a: 2, aPoint: 3} {a: 3, aPoint: 3}]

	// when the first element is removed, the index 1 element is moved to index 0, the index 2 element is moved to index 1,
	// so the aPoint of index 0 is the same value as the a of index 1
}

func TestIterateSliceInClose(t *testing.T) {
	IterateSliceInClose()
}
