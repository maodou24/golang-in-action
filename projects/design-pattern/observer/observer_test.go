package observer

import "testing"

func TestObserver(t *testing.T) {
	subject := ConcreteSubject{}

	subject.RegisterObserver(ConcreteObserver{name: "observer1"})
	subject.RegisterObserver(ConcreteObserver{name: "observer2"})
	subject.RegisterObserver(ConcreteObserver{name: "observer3"})

	subject.UpdateContext("ctx1")
}
