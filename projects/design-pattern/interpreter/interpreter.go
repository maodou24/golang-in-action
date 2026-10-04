package interpreter

type Expression interface {
	interpret() int
}

type AddExpression struct {
	epx1, epx2 Expression
}

func (a *AddExpression) interpret() int {
	return a.epx1.interpret() + a.epx2.interpret()
}

type SubExpression struct {
	epx1, epx2 Expression
}

func (s *SubExpression) interpret() int {
	return s.epx1.interpret() - s.epx2.interpret()
}

type Value struct {
	n int
}

func (v Value) interpret() int {
	return v.n
}
