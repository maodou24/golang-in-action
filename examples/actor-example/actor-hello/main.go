package main

import (
	"github.com/asynkron/protoactor-go/actor"
	"os"
	"os/signal"
	"syscall"
)

type hello struct {
	Who string
}
type helloActor struct {
}

func (r *helloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *hello:
		context.Logger().Info("Hello ", "who", msg.Who)
	}
}

func main() {
	system := actor.NewActorSystem()

	props := actor.PropsFromProducer(func() actor.Actor { return &helloActor{} })

	pid := system.Root.Spawn(props)

	system.Root.Send(pid, &hello{Who: "world"})

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}
