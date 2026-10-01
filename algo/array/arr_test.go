package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	type data struct {
		arr    []int
		except int
	}

	testData := []data{
		{
			arr:    []int{1, 4, 2},
			except: 4,
		},
		{
			arr:    []int{-1, 5, 3, 19, 9, -10},
			except: 19,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, Max(d.arr))
	}
}

func TestSum(t *testing.T) {
	type data struct {
		arr    []int
		except int
	}

	testData := []data{
		{
			arr:    []int{1, 4, 2, 3},
			except: 10,
		},
		{
			arr:    []int{-1, 5, 3, 19, 9, -10},
			except: 25,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, Sum(d.arr))
	}
}

func TestAverage(t *testing.T) {
	type data struct {
		arr    []int
		except int
	}

	testData := []data{
		{
			arr:    []int{1, 6, 2, 3},
			except: 3,
		},
		{
			arr:    []int{2, 5, 3, 9, 9, 2},
			except: 5,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, Average(d.arr))
	}
}