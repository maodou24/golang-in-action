package go1_25

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWaitGroupGo(t *testing.T) {
	WaitGroupGo()
}

func TestReflectTypeAssert(t *testing.T) {
	var a int = 1
	value := reflect.ValueOf(a)
	origA, ok := reflect.TypeAssert[int](value)
	fmt.Println(ok)    // true
	fmt.Println(origA) // 1
}
