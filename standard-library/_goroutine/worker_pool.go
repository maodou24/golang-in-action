package _goroutine

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	tasks      chan func() error
	maxWorkers int
	signal     chan struct{}
	errors     chan error
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	return &WorkerPool{
		tasks:      make(chan func() error, 100),
		maxWorkers: maxWorkers,
		signal:     make(chan struct{}, maxWorkers),
		errors:     make(chan error, maxWorkers),
	}
}

func (w *WorkerPool) Submit(task func() error) {
	w.tasks <- task
}

func (w *WorkerPool) Start() {
	var wg sync.WaitGroup

	go func() {
		for err := range w.errors {
			fmt.Println(err) // handle error
		}
	}()

	for i := 0; i < w.maxWorkers; i++ {
		wg.Add(1)
		w.signal <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-w.signal }()

			for task := range w.tasks {
				w.errors <- task()
			}
		}()
	}

	wg.Wait()
}
