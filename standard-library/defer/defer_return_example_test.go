package defer_example

import (
	"fmt"
	"testing"
)

// defer function execution and return value assignment priority
func TestDeferFuncExecAndReturnPriority(t *testing.T) {
	fn1 := func(i int) (t int) {
		t = i

		defer func() {
			t = t + 3
		}()

		return t
	}
	fmt.Printf("fn1(1) = %d\n", fn1(1)) // 4

	fn2 := func(i int) (t int) {
		t = i // 第1步修改变量t的值为i(2)

		defer func() {
			t = t + 3 // 第3步修改变量t的值为t+3
		}()

		return 1 // 第2步修改变量t的值为1
	}
	fmt.Printf("fn2(2) = %d\n", fn2(2)) // 4
	fn3 := func(i int) int {
		t := i

		defer func() {
			t = t + 3
		}()

		return t
	}
	fmt.Printf("fn3(1) = %d\n", fn3(1)) // 1
	fn4 := func(i int) int {
		t := i

		defer func(t int) {
			t = t + 3
		}(t)

		return t
	}
	fmt.Printf("fn4(1) = %d\n", fn4(1)) // 1

	a := 1
	fn5 := func(i *int) *int {
		defer func() {
			*i = *i + 3
		}()

		return i
	}
	fmt.Printf("fn5(1) = %d\n", *fn5(&a)) // 4

	fn6 := func(i int) int {
		defer func(i int) {
			i = i + 3
		}(i)

		return i
	}
	fmt.Printf("fn6(1) = %d\n", fn6(1)) // 1
}
