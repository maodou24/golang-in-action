package defer_example

import (
	"fmt"
	"testing"
)

// defer function execution sequence.
// the last deferred function is executed first
func TestDeferExecFuncSeq(t *testing.T) {
	defer fmt.Print("A")
	defer fmt.Print("B")
	defer fmt.Print("C")

	// output
	// CBA
}

// the chain functions will seem as a single function
func TestDeferFuncList(t *testing.T) {
	s := new(S)
	defer s.add(1).add(2).add(3)
}
