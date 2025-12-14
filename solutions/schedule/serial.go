package schedule

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
)

type SerialScheduler struct {
	mu  sync.RWMutex
	ctx context.Context

	queue chan *Task
	wg    sync.WaitGroup

	callback func(task *Task)
}

func NewSerialScheduler(ctx context.Context, maxTask int, callback func(task *Task)) *SerialScheduler {
	s := &SerialScheduler{
		queue:    make(chan *Task, maxTask),
		ctx:      ctx,
		callback: callback,
	}

	s.wg.Go(func() {
		for {
			select {
			case <-s.ctx.Done():
				return
			case task, ok := <-s.queue:
				if !ok {
					return // chan closed
				}

				task.StartedAt = time.Now()
				s.execute(task)
				task.EndedAt = time.Now()
			}
		}
	})

	return s
}

func (s *SerialScheduler) Submit(task *Task) error {
	task.ID = uuid.New().String()
	select {
	case s.queue <- task:
		task.EnqueuedAt = time.Now()
	default:
		return errors.New("task queue reach max capacity")
	}
	return nil
}

func (s *SerialScheduler) execute(task *Task) {
	task.StartedAt = time.Now()
	task.Err = task.Fn()
	task.EndedAt = time.Now()

	s.callback(task)
}
