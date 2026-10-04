package template_method

import "fmt"

type Template interface {
	first()
	second()
	third()
}

type TemplateMethod struct {
	Template
}

func NewTemplateMethod(t Template) *TemplateMethod {
	return &TemplateMethod{t}
}

func (t *TemplateMethod) DoSomething() {
	t.first()

	t.second()

	t.third()
}

type TemplateImpl struct {
}

func (t *TemplateImpl) first() {
	fmt.Println("first")
}

func (t *TemplateImpl) second() {
	fmt.Println("second")
}

func (t *TemplateImpl) third() {
	fmt.Println("third")
}
