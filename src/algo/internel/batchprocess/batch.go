package chainid

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

type TaskID string

type Device struct {
	ID           TaskID
	Execute      func() error
	Dependencies []TaskID
	State        DeviceState

	// result
	Err error
}

type DeviceState int

const (
	TaskStatePending DeviceState = iota
	TaskStateRunning
	TaskStateCompleted
	TaskStateFailed
)

type Batch struct {
	devices  map[TaskID]*Device
	mu       sync.RWMutex
	notifier chan TaskID
	count    int
	ctx      context.Context

	queue chan *Device
}

func NewBatch(ctx context.Context) *Batch {
	return &Batch{
		devices:  make(map[TaskID]*Device, 10),
		notifier: make(chan TaskID, 100),
		queue:    make(chan *Device, 100),
		ctx:      ctx,
	}
}

func (w *Batch) Add(d *Device) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.devices[d.ID] = d
}

func (w *Batch) Start() error {
	var wg sync.WaitGroup

	wg.Go(w.waitForExecution)

	wg.Go(func() {
		for i := 0; i < 5; i++ {
			go func() {
				select {
				case <-w.ctx.Done():
					return
				case task := <-w.queue:
					go w.executeTask(task)
				}
			}()
		}

	})

	w.mu.RLock()
	for _, task := range w.devices {
		if len(task.Dependencies) == 0 && task.State == TaskStatePending {
			w.queue <- task
		}
	}
	w.mu.RUnlock()

	wg.Wait()

	return nil
}

func (w *Batch) executeTask(task *Device) {
	w.mu.Lock()
	task.State = TaskStateRunning
	w.mu.Unlock()

	fmt.Printf("开始执行任务: %s\n", task.ID)

	err := task.Execute()

	w.mu.Lock()
	if err != nil {
		task.State = TaskStateFailed
		fmt.Printf("任务失败: %s, 错误: %v\n", task.ID, err)
	} else {
		task.State = TaskStateCompleted
		fmt.Printf("任务完成: %s\n", task.ID)
	}
	w.mu.Unlock()

	// 通知任务完成
	w.notifier <- task.ID
}

func (w *Batch) waitForExecution() {
	for {
		select {
		case <-w.ctx.Done():
			return
		case completedTaskID := <-w.notifier:
			w.mu.RLock()
			for _, task := range w.devices {
				if contains(task.Dependencies, completedTaskID) && w.canStart(task) {
					w.queue <- task
				}
			}
			w.mu.RUnlock()

			w.count++
			if w.count == len(w.devices) {
				close(w.queue)
				return
			}
		}
	}
}

func (w *Batch) canStart(task *Device) bool {
	for _, depID := range task.Dependencies {
		if dep, exists := w.devices[depID]; !exists || dep.State != TaskStateCompleted {
			return false
		}
	}
	return task.State == TaskStatePending
}

func contains(slice []TaskID, item TaskID) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

type TrieNode struct {
	IsEnd    bool
	ChainID  ChainID
	Task     *Device
	Children map[string]*TrieNode
}

func NewTrieNode(devices []*Device) *TrieNode {
	root := &TrieNode{Children: make(map[string]*TrieNode)}

	for _, task := range devices {
		parts := strings.Split(string(task.ID), ".")
		node := root
		for _, part := range parts {
			if node.Children[part] == nil {
				node.Children[part] = &TrieNode{Children: make(map[string]*TrieNode)}
			}
			node = node.Children[part]
		}
		node.IsEnd = true
		node.Task = task
		node.ChainID = ChainID(task.ID)
	}

	return root
}
