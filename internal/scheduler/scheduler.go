package scheduler

import (
	"fmt"
	"github.com/devkilyungi/time-scheduler/internal/dependencies"
	"github.com/devkilyungi/time-scheduler/internal/errors"
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

func (s *Scheduler) Add(t task.Task) {
	s.tasks = append(s.tasks, t)
}

func (s *Scheduler) ViewTasks() {
	if len(s.tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, t := range s.tasks {
		fmt.Printf("- %s: %s\n", t.Name, t.Status())
	}
}

func (s *Scheduler) RunAll() {
	for i := range s.tasks {
		if s.tasks[i].IsPending() {
			_ = s.tasks[i].Execute(os.Stdout, dependencies.NewConfigurableSleeper(1*time.Second, time.Sleep))
		} else {
			fmt.Printf("- %s is already %s!\n", s.tasks[i].Name, s.tasks[i].Status())
		}
	}
}

func (s *Scheduler) RunPending() {
	for i := range s.tasks {
		if s.tasks[i].IsPending() {
			_ = s.tasks[i].Execute(os.Stdout, dependencies.NewConfigurableSleeper(1*time.Second, time.Sleep))
		} else {
			fmt.Printf("- %s is already %s!\n", s.tasks[i].Name, s.tasks[i].Status())
		}
	}
}

func (s *Scheduler) Delete(name string) error {
	for i, t := range s.tasks {
		if t.Name == name {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.ErrTaskNotFound
}

func (s *Scheduler) Reschedule(name string, delay int) error {
	for i := range s.tasks {
		if s.tasks[i].Name == name {
			s.tasks[i].Delay = int(delay)
			s.tasks[i].Reschedule()
			return nil
		}
	}
	return errors.ErrTaskNotFound
}
