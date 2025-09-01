package go1_23

import (
	"fmt"
)

func RangeOverFunc() {
	fibo := func(yield func(x int) bool) {
		f0, f1 := 0, 1
		for yield(f0) {
			f0, f1 = f1, f0+f1
		}
	}

	for x := range fibo {
		if x > 1000 {
			break
		}
		fmt.Printf("%d ", x) // 0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987
	}
}
