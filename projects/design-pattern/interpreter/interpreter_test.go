package interpreter

import "testing"

func TestInterpreterAdd(t *testing.T) {
	expression := AddExpression{
		epx1: Value{n: 1},
		epx2: Value{n: 2},
	}

	if expression.interpret() != 3 {
		t.Fatal()
	}
}

func TestInterpreterSub(t *testing.T) {
	expression := SubExpression{
		epx1: Value{n: 3},
		epx2: Value{n: 2},
	}

	if expression.interpret() != 1 {
		t.Fatal()
	}
}
