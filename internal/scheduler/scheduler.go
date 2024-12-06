package scheduler

import (
	"github.com/devkilyungi/time-scheduler/internal/dependencies"
	"github.com/devkilyungi/time-scheduler/internal/task"
	"os"
	"time"
)

type Scheduler struct {
	tasks []task.Task
}

func NewScheduler() *Scheduler {
	return &Scheduler{tasks: []task.Task{}}
}

func (s *Scheduler) Add(task task.Task) {
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) ViewTasks() []task.Task {
	return s.tasks
}

func (s *Scheduler) RunAll() {
	for _, t := range s.tasks {
		sleeper := dependencies.NewConfigurableSleeper(1*time.Second, time.Sleep)
		_ = t.Execute(os.Stdout, sleeper)
	}
}
