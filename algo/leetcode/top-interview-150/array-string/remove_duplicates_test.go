package array_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type data struct {
		nums   []int
		val    int
		except int
	}

	testData := []data{
		{
			nums:   []int{1, 1, 2},
			except: 2,
		},
		{
			nums:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			except: 5,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, removeDuplicates(d.nums))
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	type data struct {
		nums   []int
		val    int
		except int
	}

	testData := []data{
		{
			nums:   []int{1, 1, 1, 2, 2, 3},
			except: 5,
		},
		{
			nums:   []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			except: 7,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, removeDuplicates2(d.nums))
	}
}
