package slice

import (
	"cmp"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestSlicesSort(t *testing.T) {
	s := []int{3, 5, 1}
	slices.Sort(s)

	assert.True(t, slices.IsSorted(s))
}

func TestSlicesSortFunc(t *testing.T) {
	s := []int{3, 5, 1}
	slices.SortFunc(s, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	assert.True(t, slices.IsSorted(s))
}

func TestSlicesBinarySearch(t *testing.T) {
	s := []int{1, 3, 6, 7, 9, 13, 18}
	idx, ok := slices.BinarySearch(s, 7)
	assert.True(t, ok)
	assert.Equal(t, idx, 3)

	idx, ok = slices.BinarySearch(s, 3)
	assert.True(t, ok)
	assert.Equal(t, idx, 1)

	idx, ok = slices.BinarySearch(s, 4)
	assert.False(t, ok)
	assert.Equal(t, idx, 2) // near target value
}
