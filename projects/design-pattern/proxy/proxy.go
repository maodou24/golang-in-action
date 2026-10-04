package proxy

import "fmt"

type Subject interface {
	Do()
}

type RealSubject struct{}

func (RealSubject) Do() {
	fmt.Println("RealSubject")
}

func NewProxy() Proxy {
	return Proxy{subject: &RealSubject{}}
}

type Proxy struct {
	subject *RealSubject
}

func (p Proxy) Do() {
	fmt.Println("proxy start")
	p.subject.Do()
}
