package _defer

import (
	"fmt"
	"testing"
)

func TestDeferFunc(t *testing.T) {
	fmt.Println(DeferFunc1(1)) // 4
	fmt.Println(DeferFunc2(2)) // 4
	fmt.Println(DeferFunc3(1)) // 1
	fmt.Println(DeferFunc4(1)) // 1

	var a = 1
	fmt.Println(*DeferFunc5(&a)) // 4
	fmt.Println(DeferFunc6(1))   // 1
}
