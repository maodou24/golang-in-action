package tasks

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGlobalTaskStart(t *testing.T) {
	done := make(chan struct{})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	err := StartGlobalTask(ctx, func(ctx context.Context) error {
		defer close(done)

		tk := time.NewTicker(time.Second)
		defer tk.Stop()

		for {
			select {
			case <-ctx.Done():
				return nil
			case <-tk.C:
				fmt.Println("Task is running...")
			}
		}
	})
	if err != nil {
		t.Fatal(err)
	}

	<-done
	cancel()
}

func TestGlobalTaskStop(t *testing.T) {
	done := make(chan struct{})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	err := StartGlobalTask(ctx, func(ctx context.Context) error {
		defer close(done)

		tk := time.NewTicker(time.Second)
		defer tk.Stop()

		for {
			select {
			case <-ctx.Done():
				return nil
			case <-tk.C:
				fmt.Println("Task is running...")
			}
		}
	})
	if err != nil {
		t.Fatal(err)
	}
	StopGlobalTask()

	<-done
	cancel()
}
