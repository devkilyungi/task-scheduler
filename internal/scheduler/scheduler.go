package scheduler

import (
	"fmt"
	"github.com/devkilyungi/time-scheduler/internal/dependencies"
	"github.com/devkilyungi/time-scheduler/internal/errors"
	"github.com/devkilyungi/time-scheduler/internal/task"
	"io"
)

type Scheduler struct {
	tasks   []task.Task
	writer  io.Writer
	sleeper dependencies.Sleeper
}

func NewScheduler(writer io.Writer, sleeper dependencies.Sleeper) *Scheduler {
	return &Scheduler{
		tasks:   []task.Task{},
		writer:  writer,
		sleeper: sleeper,
	}
}

func (s *Scheduler) Add(t task.Task) {
	s.tasks = append(s.tasks, t)
}

func (s *Scheduler) ViewTasks() {
	if len(s.tasks) == 0 {
		_, _ = fmt.Fprintln(s.writer, "No tasks found.")
		return
	}
	for _, t := range s.tasks {
		_, _ = fmt.Fprintf(s.writer, "- %s: %s\n", t.Name, t.Status())
	}
}

func (s *Scheduler) RunAll() {
	if len(s.tasks) == 0 {
		_, _ = fmt.Fprintln(s.writer, "No tasks found.")
		return
	}

	for i := range s.tasks {
		if s.tasks[i].IsPending() {
			_ = s.tasks[i].Execute(s.writer, s.sleeper)
		} else {
			_, _ = fmt.Fprintf(s.writer, "- %s is already %s!\n", s.tasks[i].Name, s.tasks[i].Status())
		}
	}
}

func (s *Scheduler) RunPending() {
	if len(s.tasks) == 0 {
		_, _ = fmt.Fprintln(s.writer, "No tasks found.")
		return
	}

	for i := range s.tasks {
		if s.tasks[i].IsPending() {
			_ = s.tasks[i].Execute(s.writer, s.sleeper)
		} else {
			_, _ = fmt.Fprintf(s.writer, "- %s is already %s!\n", s.tasks[i].Name, s.tasks[i].Status())
		}
	}
}

func (s *Scheduler) Delete(name string) error {
	if len(s.tasks) == 0 {
		_, _ = fmt.Fprintln(s.writer, "No tasks found.")
		return errors.ErrTaskNotFound
	}

	for i, t := range s.tasks {
		if t.Name == name {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.ErrTaskNotFound
}

func (s *Scheduler) Reschedule(name string, delay int) error {
	if len(s.tasks) == 0 {
		_, _ = fmt.Fprintln(s.writer, "No tasks found.")
		return errors.ErrTaskNotFound
	}

	for i := range s.tasks {
		if s.tasks[i].Name == name {
			s.tasks[i].Delay = delay
			s.tasks[i].Reschedule()
			return nil
		}
	}
	return errors.ErrTaskNotFound
}
