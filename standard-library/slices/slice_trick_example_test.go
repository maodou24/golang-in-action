package slice

import (
	"fmt"
	"sync"
	"testing"
)

// slice trick: when the slice is a slice of struct,
// and the struct has a pointer field, when the slice is reordered,
// the pointer field will point to the wrong value
func TestSliceTrickPointerElementReorder(t *testing.T) {
	s := []B{
		{a: A{N: 1}},
		{a: A{N: 2}},
		{a: A{N: 3}},
	}

	for i := range s {
		s[i].aPoint = &s[i].a
	}

	fmt.Println("s before operation is", s)
	// s before operation is
	// s: [{a: 1, aPoint: 1} {a: 2, aPoint: 2} {a: 3, aPoint: 3}]
	// remove the first element
	s = append(s[:0], s[1:]...)
	fmt.Println("s after operation is", s)
	// s after operation is
	// s: [{a: 2, aPoint: 3} {a: 3, aPoint: 3}]

	// when the first element is removed, the index 1 element is moved to index 0, the index 2 element is moved to index 1,
	// so the aPoint of index 0 is the same value as the a of index 1
}

// slice trick: when use element of slice in goroutine,
// the element will be captured by reference, not by value
// this will cause the goroutine to use the last value of the slice,
// not the value at the time of the goroutine creation
// note: the feature is changed in go1.20, now the element will be captured by value, not by reference
func TestSliceIterateClosure(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	var wg sync.WaitGroup
	for _, v := range s {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(v)
		}()
	}
	wg.Wait()

	// before go1.20
	// the output is: 9 9 9 9 9 9 9 9 9 9
	// after go1.20
	// the output is random: 0~9
}
