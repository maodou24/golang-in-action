package slice

import (
	"cmp"
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlicesSort(t *testing.T) {
	s := []int{3, 5, 1}
	slices.Sort(s)
	// s: [1 3 5]

	assert.True(t, slices.IsSorted(s))
}

func TestSlicesSortStableFunc(t *testing.T) {
	type item struct {
		key int
		id  string
	}

	input := []item{
		{key: 2, id: "A"},
		{key: 1, id: "B"},
		{key: 2, id: "C"},
		{key: 1, id: "E"},
		{key: 2, id: "E"},
		{key: 2, id: "F"},
		{key: 2, id: "G"},
	}
	compare := func(a, b item) int {
		return cmp.Compare(a.key, b.key)
	}

	stable := slices.Clone(input)
	unstable := slices.Clone(input)
	// 相同 key 的元素在稳定排序后仍保持输入中的相对顺序。
	slices.SortStableFunc(stable, compare)
	// 普通排序只保证按 key 排序；相同 key 的相对顺序不作保证。
	slices.SortFunc(unstable, compare)

	fmt.Printf("SortStableFunc: %v\n", stable)
	fmt.Printf("SortFunc:       %v\n", unstable)
}

func TestSlicesSortFunc(t *testing.T) {
	s := []int{3, 5, 1}
	slices.SortFunc(s, func(a, b int) int {
		return cmp.Compare(a, b)
	})
	// s: [1 3 5]
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
