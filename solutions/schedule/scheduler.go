package schedule

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/sync/semaphore"
	"sync"
	"time"
)

type Scheduler struct {
	globalSem       *semaphore.Weighted
	groupSems       map[TaskType]*semaphore.Weighted
	groupMax        map[TaskType]int
	userSems        map[string]*semaphore.Weighted
	userMax         int
	taskTypeToGroup map[TaskType]string

	mu  sync.RWMutex
	ctx context.Context

	queue chan *Task
	wg    sync.WaitGroup
}

var taskTypeToGroup = map[string]string{
	"task1": "combine_sem1",
}

func NewScheduler(ctx context.Context, max int, perUser int, group map[TaskType]int, typeToGroup map[TaskType]string) *Scheduler {
	s := &Scheduler{
		globalSem:       semaphore.NewWeighted(int64(max)),
		groupSems:       make(map[TaskType]*semaphore.Weighted),
		groupMax:        group,
		userMax:         perUser,
		userSems:        make(map[string]*semaphore.Weighted),
		taskTypeToGroup: typeToGroup,
		queue:           make(chan *Task, 20),
		ctx:             ctx,
	}

	for i := 0; i < max*2; i++ {
		s.wg.Go(s.work)
	}

	return s
}

func (s *Scheduler) Submit(task *Task) error {
	task.ID = uuid.New().String()
	task.EnqueuedAt = time.Now()
	s.queue <- task
	return nil
}

func (s *Scheduler) work() {
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
}

func (s *Scheduler) execute(task *Task) {
	globalSem := s.globalSem
	groupSem := s.getGroupSem(task.Type)
	userSem := s.getUserSem(task.UserID)

	if err := globalSem.Acquire(s.ctx, 1); err != nil {
		return
	}
	defer globalSem.Release(1)

	if err := groupSem.Acquire(s.ctx, 1); err != nil {
		return
	}
	defer groupSem.Release(1)

	if err := userSem.Acquire(s.ctx, 1); err != nil {
		return
	}
	defer userSem.Release(1)

	task.StartedAt = time.Now()
	task.Err = task.Fn()
	task.EndedAt = time.Now()
}

func (s *Scheduler) getGroupSem(task TaskType) *semaphore.Weighted {
	group := TaskType(s.getQuotaGroup(task))
	s.mu.RLock()
	sem, ok := s.groupSems[group]
	if ok {
		s.mu.RUnlock()
		return sem
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()
	sem, ok = s.groupSems[group]
	if ok {
		return sem
	}

	n := max(s.groupMax[group], 1)
	sem = semaphore.NewWeighted(int64(n))
	s.groupSems[group] = sem
	return sem
}

func (s *Scheduler) getUserSem(userId string) *semaphore.Weighted {
	s.mu.RLock()
	sem, ok := s.userSems[userId]
	if ok {
		s.mu.RUnlock()
		return sem
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()
	sem, ok = s.userSems[userId]
	if ok {
		return sem
	}

	n := max(s.userMax, 1)
	sem = semaphore.NewWeighted(int64(n))
	s.userSems[userId] = sem
	return sem
}

func (s *Scheduler) getQuotaGroup(taskType TaskType) string {
	group, ok := s.taskTypeToGroup[taskType]
	if ok {
		return group
	}

	return string(taskType)
}
