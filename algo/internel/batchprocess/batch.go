package chainid

import (
	"fmt"
	"strings"
	"sync"
)

type TaskID string

type Task struct {
	ID           TaskID
	Execute      func() error
	Dependencies []TaskID
	State        TaskState

	// result
	Err error
}

type TaskState int

const (
	TaskStatePending TaskState = iota
	TaskStateRunning
	TaskStateCompleted
	TaskStateFailed
)

type Batch struct {
	tasks    map[TaskID]*Task
	mu       sync.RWMutex
	notifier chan TaskID
	count    int

	queue chan *Task
}

func NewWorkflow() *Batch {
	return &Batch{
		tasks:    make(map[TaskID]*Task, 10),
		notifier: make(chan TaskID, 100),
		queue:    make(chan *Task, 100),
	}
}

func (w *Batch) AddTask(task *Task) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.tasks[task.ID] = task
}

func (w *Batch) Start() error {
	var wg sync.WaitGroup

	wg.Go(w.waitForExecution)

	wg.Go(func() {
		for i := 0; i < 5; i++ {
			go func() {
				for task := range w.queue {
					go w.executeTask(task)
				}
			}()
		}

	})

	w.mu.RLock()
	for _, task := range w.tasks {
		if len(task.Dependencies) == 0 && task.State == TaskStatePending {
			w.queue <- task
		}
	}
	w.mu.RUnlock()

	wg.Wait()

	return nil
}

func (w *Batch) executeTask(task *Task) {
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
	for completedTaskID := range w.notifier {
		w.mu.RLock()
		for _, task := range w.tasks {
			if contains(task.Dependencies, completedTaskID) && w.canStart(task) {
				w.queue <- task
			}
		}
		w.mu.RUnlock()

		w.count++
		if w.count == len(w.tasks) {
			close(w.queue)
			return
		}
	}
}

func (w *Batch) canStart(task *Task) bool {
	for _, depID := range task.Dependencies {
		if dep, exists := w.tasks[depID]; !exists || dep.State != TaskStateCompleted {
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
	Task     *Task
	Children map[string]*TrieNode
}

func NewTrieNode(tasks []*Task) *TrieNode {
	root := &TrieNode{Children: make(map[string]*TrieNode)}

	for _, task := range tasks {
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
