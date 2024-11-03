package _defer

import "testing"

func TestDeferExecSeq1(t *testing.T) {
	DeferExecSeq()

	// output
	// CBA
}
