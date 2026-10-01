package go1_24

import (
	"fmt"
	"testing"
	"weak"
)

func TestWeakPkg(t *testing.T) {
	a := new(int)
	wa := weak.Make[int](a)
	if wa.Value() != nil {
		fmt.Println(*wa.Value())
	}
}
