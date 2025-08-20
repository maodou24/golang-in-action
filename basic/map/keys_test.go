package _map

import (
	"fmt"
	"testing"
)

func TestChannelAsMapKey(t *testing.T) {
	ChannelAsMapKey()
}

// go 1.22以上，行为清晰
func TestDeleteKeyDuringIteration(t *testing.T) {
	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		m[i] = i
	}

	for k := range m {
		if k%2 == 0 {
			delete(m, k)
		}
	}

	fmt.Println(m)
}
