package schedule

import "time"

type TaskType string
type Task struct {
	Type   TaskType
	UserID string

	ID         string
	Fn         func() error
	EnqueuedAt time.Time
	StartedAt  time.Time
	EndedAt    time.Time
	State      int
	Err        error // 执行错误
}
