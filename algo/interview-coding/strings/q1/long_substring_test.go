package q1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongSubstring(t *testing.T) {
	type args struct {
		str string
		k   int

		out int
	}

	testData := []args{
		{str: "abcck", k: 1, out: 2},
		{str: "abcck", k: 2, out: 3},
		{str: "lrbbmqbhcdarz", k: 2, out: 3},
		{str: "efsarcbynec", k: 5, out: 6},
	}

	for _, data := range testData {
		assert.Equal(t, data.out, longSubstring(data.str, data.k))
		assert.Equal(t, data.out, longSubstring2(data.str, data.k))
	}
}
