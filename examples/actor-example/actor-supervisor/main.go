package main

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type msg string

type create struct{}

type parentActor struct{}

func (p *parentActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *create:
		props := actor.PropsFromProducer(func() actor.Actor { return &childActor{} })

		pid := context.Spawn(props)
		log.Printf("Created actor: pid=%v\n", pid)
	case *msg:
		for _, pid := range context.Children() {
			context.Send(pid, msg)
		}
	}
}

type childActor struct{}

func (p *childActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		log.Println("actor started")
	case *actor.Stopping:
		log.Println("Stopping, actor is about to shut down")
	case *actor.Stopped:
		log.Println("Stopped, actor and its children are stopped")
	case *actor.Restarting:
		log.Println("Restarting, actor is about to restart")
	case *msg:
		log.Printf("Hello %v\n", *msg)
		panic("Ouch")
	}
}

func main() {
	system := actor.NewActorSystem()

	decider := func(reason interface{}) actor.Directive {
		fmt.Println("handling failure for child")
		return actor.StopDirective
	}

	strategy := actor.NewOneForOneStrategy(10, 1000, decider)

	props := actor.PropsFromProducer(func() actor.Actor { return &parentActor{} },
		actor.WithSupervisor(strategy))

	pid := system.Root.Spawn(props)

	// create new actor1
	system.Root.Send(pid, &create{})

	// create new actor2
	system.Root.Send(pid, &create{})

	m := msg("hello world")
	// send msg
	system.Root.Send(pid, &m)

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}
