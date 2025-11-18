package schedule

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestScheduler_Submit(t *testing.T) {
	s := NewScheduler(context.Background(), 1, 1, map[TaskType]int{}, make(map[TaskType]string))

	_ = s.Submit(&Task{
		Type:   "task",
		UserID: "1",
		Fn: func() error {
			fmt.Println("do task 1")
			return nil
		},
	})

	_ = s.Submit(&Task{
		Type:   "task2",
		UserID: "1",
		Fn: func() error {
			fmt.Println("do task 2")
			return nil
		},
	})

	_ = s.Submit(&Task{
		Type:   "task3",
		UserID: "1",
		Fn: func() error {
			fmt.Println("do task 3")
			return nil
		},
	})

	_ = s.Submit(&Task{
		Type:   "task4",
		UserID: "1",
		Fn: func() error {
			fmt.Println("do task 4")
			return nil
		},
	})

	//for t := range s.queue {
	//	fmt.Println("task type:", t.Type, "user:", t.UserID)
	//}
	time.Sleep(1 * time.Second)
}
