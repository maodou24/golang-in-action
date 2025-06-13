package main

import (
	"github.com/rs/xid"
)

type TaskRepository struct {
	tasks map[string]*Task
}

var repo = TaskRepository{tasks: make(map[string]*Task, 10)}

func (r *TaskRepository) Save(task *Task) error {
	task.TaskId = xid.New().String()
	r.tasks[task.TaskId] = task
	return nil
}

func (r *TaskRepository) Find(taskId string) (*Task, error) {
	return r.tasks[taskId], nil
}
