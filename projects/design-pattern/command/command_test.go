package command

import "testing"

func TestCommand(t *testing.T) {
	invoker := &Invoker{}

	invoker.Add(NewCommandA(Receiver{name: "maodou"}))

	invoker.invoker()
}
