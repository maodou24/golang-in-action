package sorting

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorts(t *testing.T) {
	input := [][]int{
		{1, 22, 3, 15, 9, 8},
	}

	except := [][]int{
		{1, 3, 8, 9, 15, 22},
	}

	testFunc := []func([]int){
		insert, selectSort, shell, quickSort,
	}

	for i, sort := range testFunc {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			for i := range input {
				temp := make([]int, len(input[i]))
				copy(temp, input[i])
				sort(temp)
				assert.Equal(t, except[i], temp)
			}
		})
	}
}
