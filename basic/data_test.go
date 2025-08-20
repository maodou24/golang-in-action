package basic

import (
	"fmt"
	"testing"
)

func TestAllocationNew(t *testing.T) {
	a := new(int)
	*a = 1
	fmt.Println(*a)
}
