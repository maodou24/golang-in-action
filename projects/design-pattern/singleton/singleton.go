package singleton

import (
	"sync"
)

type Singleton interface {
	DoSomething() string
}

type singleton struct{}

func (s *singleton) DoSomething() string {
	return "Hello World!"
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
