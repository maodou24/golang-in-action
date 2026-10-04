package memento

type Memento struct {
	state string
}

func (m *Memento) GetState() string {
	return m.state
}

type Originator struct {
	memento *Memento

	state string
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) SaveToMemento() *Memento {
	return &Memento{state: o.state}
}

func (o *Originator) RecoverFromMemento(m *Memento) {
	o.state = m.GetState()
}

type Caretaker struct {
	memento []*Memento
}

func (c *Caretaker) Add(m *Memento) {
	c.memento = append(c.memento, m)
}

func (c *Caretaker) Get(idx int) *Memento {
	return c.memento[idx]
}
