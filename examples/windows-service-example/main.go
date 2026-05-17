package main

import (
	"github.com/kardianos/service"
)

func main() {
	svcConfig := &service.Config{
		Name:        "Test Windows Service",
		DisplayName: "Test Windows Service",
		Description: "A simple Windows service",
	}

	program := &program{}
	svc, err := service.New(program, svcConfig)
	if err != nil {
		panic(err)
	}

	err = svc.Run()
	if err != nil {
		panic(err)
	}
}

type program struct{}

// should not block.
func (p *program) Start(s service.Service) error {
	// Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() {
	// Do work here
}

// should not block.
func (p *program) Stop(s service.Service) error {
	// clean resources before service stops.
	return nil
}