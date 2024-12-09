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

func TestTaskExecution(t *testing.T) {
	taskA := &Task{
		Name:   "Task A",
		Delay:  4,
		status: Pending,
	}
	spyExecution := &SpyExecution{}

	t.Run("task is created with correct details", func(t *testing.T) {
		wantName := "Task A"
		wantDelay := 4
		wantStatus := Pending

		if taskA.Name != wantName {
			t.Errorf("taskA.Name is %s; want %s", taskA.Name, wantName)
		}

		if taskA.Delay != wantDelay {
			t.Errorf("taskA.Delay is %d; want %d", taskA.Delay, wantDelay)
		}

		if taskA.status != wantStatus {
			t.Errorf("taskA.status is %s; want %s", taskA.status, wantStatus)
		}
	})

	t.Run("task executes in the correct order", func(t *testing.T) {
		err := taskA.Execute(spyExecution, spyExecution)
		if err != nil {
			t.Fatalf("task execution failed: %v", err)
		}

		want := []string{write, sleep, write, sleep, write, sleep, write, sleep, write}
		if !reflect.DeepEqual(want, spyExecution.Calls) {
			t.Fatalf("got %v, want %v", spyExecution.Calls, want)
		}
	})

	t.Run("task changes status to Completed on completion", func(t *testing.T) {
		err := taskA.Execute(spyExecution, spyExecution)
		if err != nil {
			t.Fatalf("task execution failed: %v", err)
		}

		want := Completed
		if taskA.status != want {
			t.Fatalf("got %v, want %v", taskA.status, want)
		}
	})

	t.Run("task countdown prints correct output", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyExecution{}
		err := taskA.Execute(buffer, spySleeper)
		if err != nil {
			t.Fatalf("task execution failed: %v", err)
		}

		got := buffer.String()
		want := `4...3...2...1...
Task A executed!
`

		if got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}
