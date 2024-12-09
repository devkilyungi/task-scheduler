package scheduler

import (
	"bytes"
	"github.com/devkilyungi/time-scheduler/internal/task"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyWriter struct {
	Buffer *bytes.Buffer
}

func NewSpyWriter() *SpyWriter {
	return &SpyWriter{Buffer: new(bytes.Buffer)}
}

func (s *SpyWriter) Write(p []byte) (n int, err error) {
	return s.Buffer.Write(p)
}

func (s *SpyWriter) String() string {
	return s.Buffer.String()
}

func TestScheduler(t *testing.T) {
	spySleeper := &SpySleeper{}
	spyWriter := NewSpyWriter()
	scheduler := NewScheduler(spyWriter, spySleeper)

	t.Run("adding a task adds it to the scheduler", func(t *testing.T) {
		taskA := task.Task{Name: "Task A", Delay: 3}
		scheduler.Add(taskA)

		if len(scheduler.tasks) != 1 {
			t.Errorf("expected %d tasks, got %d", 1, len(scheduler.tasks))
		}
	})

	t.Run("viewing tasks writes to the writer", func(t *testing.T) {
		taskA := task.Task{Name: "Task A", Delay: 2}
		scheduler.Add(taskA)

		scheduler.ViewTasks()

		got := spyWriter.String()
		want := "- Task A: Pending\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("running all tasks executes all tasks", func(t *testing.T) {
		taskA := task.Task{Name: "Task A", Delay: 2}
		taskB := task.Task{Name: "Task B", Delay: 3}

		scheduler.Add(taskA)
		scheduler.Add(taskB)
		scheduler.RunAll()

		want := taskA.Delay + taskB.Delay

		if spySleeper.Calls != want {
			t.Errorf("expected sleeper to be called %d time, but got %d", want, spySleeper.Calls)
		}
	})

	t.Run("running pending tasks executes only pending tasks", func(t *testing.T) {
		taskA := task.Task{Name: "Task A", Delay: 2}
		scheduler.Add(taskA)
		scheduler.RunAll()
		spySleeper.Calls = 0 // Reset spy calls for Task A

		taskB := task.Task{Name: "Task B", Delay: 3}
		scheduler.Add(taskB)
		scheduler.RunPending()

		want := taskB.Delay
		if spySleeper.Calls != want {
			t.Errorf("expected sleeper to be called %d time, but got %d", want, spySleeper.Calls)
		}
	})

	t.Run("deleting a task removes it from the scheduler", func(t *testing.T) {
		taskA := task.Task{Name: "Task A", Delay: 2}
		scheduler.Add(taskA)
		err := scheduler.Delete(taskA.Name)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(scheduler.tasks) != 0 {
			t.Errorf("expected 0 tasks, but got %d", len(scheduler.tasks))
		}
	})

	t.Run("deleting a non-existent task returns an error", func(t *testing.T) {
		err := scheduler.Delete("NonExistentTask")

		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})

	t.Run("rescheduling a task reschedules it", func(t *testing.T) {
		taskA := task.Task{Name: "Task A", Delay: 2}
		scheduler.Add(taskA)
		scheduler.RunAll()
		err := scheduler.Reschedule(taskA.Name, 3)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !taskA.IsPending() {
			t.Errorf("expected task to be pending but was not")
		}
	})

	t.Run("rescheduling a non-existent task returns an error", func(t *testing.T) {
		err := scheduler.Reschedule("NonExistentTask", 3)

		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})
}
