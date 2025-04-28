package _defer

import "fmt"

// output: CBA
func DeferExecSeq() {
	defer fmt.Print("A")
	defer fmt.Print("B")
	defer fmt.Print("C")
}
