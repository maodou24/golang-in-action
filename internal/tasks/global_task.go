package tasks

import (
	"context"
	"fmt"
	"sync"
)

type GlobalTask struct {
	sync.Mutex
	running bool
	fn      func(ctx context.Context) error
	cancel  context.CancelFunc
}

var globalTask = &GlobalTask{}

func StartGlobalTask(ctx context.Context, fn func(ctx context.Context) error) error {
	globalTask.Lock()
	defer globalTask.Unlock()

	if globalTask.running {
		return fmt.Errorf("global task is already running")
	}
	globalTask.running = true
	globalTask.fn = fn

	cancelCtx, cancel := context.WithCancel(ctx)
	globalTask.cancel = cancel

	go globalTask.run(cancelCtx)

	return nil
}

func StopGlobalTask() {
	globalTask.Lock()
	defer globalTask.Unlock()

	if !globalTask.running {
		return
	}
	globalTask.running = false

	if globalTask.cancel != nil {
		globalTask.cancel()
	}
	globalTask.cancel = nil
}

func (g *GlobalTask) run(ctx context.Context) {
	defer g.stop()

	g.fn(ctx)
}

func (g *GlobalTask) stop() {
	g.Lock()
	defer g.Unlock()

	if !g.running {
		return
	}
	g.running = false

	if g.cancel != nil {
		g.cancel()
	}
	g.cancel = nil
}
