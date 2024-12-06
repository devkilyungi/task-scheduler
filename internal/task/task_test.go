package task

import (
	"bytes"
	"reflect"
	"testing"
)

const write = "write"
const sleep = "sleep"

type SpyExecution struct {
	Calls []string
}

func (s *SpyExecution) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyExecution) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func TestTask_Execute(t *testing.T) {
	taskA := &Task{
		Name:   "Task A",
		Delay:  3,
		status: pending,
	}

	t.Run("prints and sleeps in the right order", func(t *testing.T) {
		spySleeper := &SpyExecution{}
		err := taskA.Execute(spySleeper, spySleeper)
		if err != nil {
			t.Fatalf("task execution failed: %v", err)
		}

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeper.Calls) {
			t.Errorf("got %v, want %v", spySleeper.Calls, want)
		}
	})

	t.Run("prints delay countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpyExecution{}
		err := taskA.Execute(buffer, sleeper)
		if err != nil {
			t.Fatalf("task execution failed: %v", err)
		}
		
		got := buffer.String()
		want := `3...2...1...
Task A executed!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if int(taskA.Delay) != len(sleeper.Calls) {
			t.Errorf("got %d calls, want %d", len(sleeper.Calls), taskA.Delay)
		}
	})
}
