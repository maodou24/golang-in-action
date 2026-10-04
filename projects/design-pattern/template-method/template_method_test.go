package template_method

import "testing"

func TestTemplateMethod(t *testing.T) {
	template := NewTemplateMethod(&TemplateImpl{})
	template.DoSomething()
}
