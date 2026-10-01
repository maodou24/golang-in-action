package list

type Ele struct {
	prev *Ele
	next *Ele

	Value any
}

type doublyList struct {
	len int

	root Ele
}

func NewDoublyList() *doublyList {
	d := &doublyList{}

	d.root.next = &d.root
	d.root.prev = &d.root
	return d
}

func (d *doublyList) Front() *Ele {
	if d.len == 0 {
		return nil
	}

	return d.root.next
}

func (d *doublyList) Back() *Ele {
	if d.len == 0 {
		return nil
	}
	return d.root.prev
}

func (d *doublyList) PushFront(v any) *Ele {
	ele := &Ele{Value: v}

	ele.next = d.root.next
	ele.prev = &d.root

	ele.next.prev = ele
	ele.prev.next = ele

	d.len++

	return ele
}

func (d *doublyList) PushBack(v any) *Ele {
	ele := &Ele{Value: v}

	tail := d.root.prev

	ele.next = tail.next
	ele.prev = tail

	ele.next.prev = ele
	ele.prev.next = ele

	d.len++

	return ele
}

func (d *doublyList) MoveToFront(e *Ele) {
	d.Remove(e)

	e.next = d.root.next
	e.prev = d.root.prev

	e.next.prev = e
	e.prev.next = e
}

func (d *doublyList) MoveToBack(e *Ele) {
	d.Remove(e)

	tail := d.root.prev

	e.next = tail.next
	e.prev = tail.prev

	e.next.prev = e
	e.prev.next = e
}

func (d *doublyList) Remove(e *Ele) {
	e.prev.next = e.next
	e.next.prev = e.prev

	e.next = nil
	e.prev = nil

	d.len--
}

func (d *doublyList) Len() int {
	return d.len
}
