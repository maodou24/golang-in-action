package factory_method

type OperatorFactory interface {
	Create() Operator
}

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type value struct {
	a, b int
}

type PlusFactory struct {
	value
}

func (v *value) SetA(a int) {
	v.a = a
}

func (v *value) SetB(b int) {
	v.b = b
}

func (p PlusFactory) Result() int {
	return p.a + p.b
}

func (p PlusFactory) Create() Operator {
	return &PlusFactory{}
}

type SubFactory struct {
	value
}

func (p SubFactory) Result() int {
	return p.a - p.b
}

func (p SubFactory) Create() Operator {
	return &SubFactory{}
}

func Calculate(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}
