package go1_22

import (
	"fmt"
)

func LoopRangeOverInteger() {
	for v := range 10 {
		fmt.Print(v) // 012345...10
	}
	fmt.Println()
}
