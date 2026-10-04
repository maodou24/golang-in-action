package prototype

type Cloneable interface {
	Clone() Cloneable
}

type prototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *prototypeManager {
	return &prototypeManager{prototypes: make(map[string]Cloneable)}
}

func (p *prototypeManager) Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p *prototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}
