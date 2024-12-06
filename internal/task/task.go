package task

import (
	"fmt"
	"github.com/devkilyungi/time-scheduler/internal/dependencies"
	"github.com/devkilyungi/time-scheduler/internal/errors"
	"io"
	"time"
)

type Status int

const (
	Pending Status = iota
	Completed
)

func (s *Status) String() string {
	return [...]string{"Pending", "Completed"}[*s]
}

type Task struct {
	Name   string
	Delay  time.Duration
	status Status
}

func (t *Task) Status() string {
	return t.status.String()
}

func (t *Task) Execute(w io.Writer, s dependencies.Sleeper) error {
	for i := t.Delay; i > 0; i-- {
		_, err := fmt.Fprintf(w, "%d...", i)
		if err != nil {
			return errors.ErrTaskFailedToExecute
		}
		s.Sleep()
	}

	t.status = Completed
	_, err := fmt.Fprintf(w, "\n%s executed!\n", t.Name)
	if err != nil {
		return errors.ErrTaskFailedToExecute
	}

	return nil
}
