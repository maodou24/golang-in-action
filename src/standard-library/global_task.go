package standard_library

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type GlobalTask struct {
	sync.Mutex
	running bool
	cancel  func()
}

var globalTask = &GlobalTask{}

func StartGlobalTask(ctx context.Context, timeout int) error {
	globalTask.Lock()
	defer globalTask.Unlock()

	if globalTask.running {
		return fmt.Errorf("global task is already running")
	}
	globalTask.running = true
	timeCtx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	globalTask.cancel = cancel

	go runGlobalTask(timeCtx)

	return nil
}

func runGlobalTask(ctx context.Context) {
	defer StopGlobalTask()
	select {
	case <-ctx.Done():
		fmt.Println("global task timeout, stop task")
		return
	default:
		fmt.Println("global task doing work")
		time.Sleep(1 * time.Second)
	}
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
