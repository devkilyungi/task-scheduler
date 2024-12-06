package task

import (
	"fmt"
	"github.com/devkilyungi/time-scheduler/internal/dependencies"
	"github.com/devkilyungi/time-scheduler/internal/errors"
	"io"
)

type Status int

const (
	Pending Status = iota
	Completed
)

func (s Status) String() string {
	return [...]string{"Pending", "Completed"}[s]
}

type Task struct {
	Name   string
	Delay  int
	status Status
}

func (t *Task) IsPending() bool {
	return t.status == Pending
}

func (t *Task) Status() string {
	return t.status.String()
}

func (t *Task) Reschedule() {
	t.status = Pending
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
	return err
}
