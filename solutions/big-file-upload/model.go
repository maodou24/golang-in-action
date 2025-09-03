package main

import (
	"errors"
	"os"
)

type Task struct {
	Filename  string `json:"filename"`
	FileSize  int64  `json:"fileSize"`
	ChunkSize int64  `json:"chunkSize"`

	TaskId   string `json:"taskId"`
	Uploaded []int  `json:"uploaded"` // uploaded chunks
}

func (t *Task) Init() error {
	if len(t.Filename) == 0 {
		return errors.New("filename is empty")
	}

	task, err := repo.Find(t.Filename)
	if err != nil {
		return err
	}

	if task == nil {
		if err := repo.Save(t); err != nil {
			return err
		}

		if err := os.MkdirAll("./upload/"+t.TaskId, os.ModePerm); err != nil {
			return err
		}
	} else {
		t.Uploaded = task.Uploaded
		t.TaskId = task.TaskId
		t.FileSize = task.FileSize
		t.ChunkSize = task.ChunkSize
	}

	return nil
}

type ChunkMeta struct {
	TaskId  string `json:"taskId"`
	ChunkId int    `json:"chunkId"`
}
