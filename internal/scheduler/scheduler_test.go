package scheduler

import (
	"github.com/devkilyungi/time-scheduler/internal/errors"
	"github.com/devkilyungi/time-scheduler/internal/task"
	"reflect"
	"testing"
)

func TestScheduler(t *testing.T) {
	taskA := task.Task{
		Name:  "Task A",
		Delay: 3,
	}

	taskB := task.Task{
		Name:  "Task B",
		Delay: 2,
	}
	scheduler := NewScheduler()

	t.Run("adding a task adds it to the scheduler", func(t *testing.T) {
		scheduler.Add(taskA)

		wantLength := 1
		wantTaskName := "Task A"
		wantTaskDelay := 3
		wantTaskStatus := task.Pending

		if len(scheduler.tasks) != wantLength {
			t.Errorf("scheduler.tasks length = %d, want %d", len(scheduler.tasks), wantLength)
		}

		if scheduler.tasks[0].Name != wantTaskName {
			t.Errorf("scheduler.tasks[0].Name = %s, want %s", scheduler.tasks[0].Name, wantTaskName)
		}

		if scheduler.tasks[0].Delay != wantTaskDelay {
			t.Errorf("scheduler.tasks[0].Delay = %v, want %v", scheduler.tasks[0].Delay, wantTaskDelay)
		}

		if scheduler.tasks[0].Status() != wantTaskStatus.String() {
			t.Errorf("scheduler.tasks[1].Status = %v, want %v", scheduler.tasks[1].Status(), wantTaskStatus)
		}
	})

	// TODO: Should be refined to pass a dependency on writer and sleeper
	// TODO: so we can mock the tests on those without waiting for the delay
	t.Run("running all tasks executes all tasks", func(t *testing.T) {
		scheduler.Add(taskA)
		scheduler.Add(taskB)

		scheduler.RunAll()

		wantStatus := []string{"Completed", "Completed"}
		var gotStatus []string

		for _, taskInScheduler := range scheduler.tasks {
			gotStatus = append(gotStatus, taskInScheduler.Status())
		}

		if !reflect.DeepEqual(gotStatus, wantStatus) {
			t.Errorf("scheduler.tasks = %v, want %v", gotStatus, wantStatus)
		}
	})

	t.Run("running pending tasks executes only pending tasks", func(t *testing.T) {
		scheduler.Add(taskA)

		scheduler.RunAll()
		scheduler.Add(taskB)

		wantStatus := []string{"Completed", "Pending"}
		var gotStatus []string

		for _, taskInScheduler := range scheduler.tasks {
			gotStatus = append(gotStatus, taskInScheduler.Status())
		}

		if !reflect.DeepEqual(gotStatus, wantStatus) {
			t.Errorf("scheduler.tasks = %v, want %v", gotStatus, wantStatus)
		}

		scheduler.RunPending()
		wantNewStatus := []string{"Completed", "Completed"}
		var gotNewStatus []string

		for _, taskInScheduler := range scheduler.tasks {
			gotNewStatus = append(gotNewStatus, taskInScheduler.Status())
		}

		if !reflect.DeepEqual(gotNewStatus, wantNewStatus) {
			t.Errorf("scheduler.tasks = %v, want %v", gotNewStatus, wantNewStatus)
		}
	})

	t.Run("deleting a task removes it from the scheduler", func(t *testing.T) {
		scheduler.Add(taskA)
		scheduler.Add(taskB)

		err := scheduler.Delete(taskA.Name)
		if err != nil {
			t.Errorf("scheduler.Delete(taskA.Name) failed: %v", err)
		}

		wantLength := 1

		if len(scheduler.tasks) != wantLength {
			t.Errorf("scheduler.tasks length = %d, want %d", len(scheduler.tasks), wantLength)
		}
	})

	t.Run("deleting a non-existent task returns an error", func(t *testing.T) {
		scheduler.Add(taskA)

		err := scheduler.Delete("Task Z")
		if err == nil {
			t.Errorf("scheduler.Delete(task Z) didn't return an error")
		}

		if err != errors.ErrTaskNotFound {
			t.Errorf("scheduler.Delete(task Z) returned err = %v, want %v", err, errors.ErrTaskNotFound)
		}
	})

	t.Run("rescheduling a tasks reschedules it", func(t *testing.T) {
		scheduler.Add(taskA)
		scheduler.Add(taskB)
		scheduler.RunAll()

		err := scheduler.Reschedule(taskA.Name, 1)
		if err != nil {
			t.Errorf("scheduler.Reschedule(taskA.Name) failed: %v", err)
		}

		wantStatus := []string{"Pending", "Completed"}
		var gotStatus []string

		for _, taskInScheduler := range scheduler.tasks {
			gotStatus = append(gotStatus, taskInScheduler.Status())
		}

		if !reflect.DeepEqual(gotStatus, wantStatus) {
			t.Errorf("scheduler.tasks = %v, want %v", gotStatus, wantStatus)
		}
	})

	t.Run("rescheduling a non-existent task return an error", func(t *testing.T) {
		err := scheduler.Reschedule("Task Z", 1)
		if err == nil {
			t.Errorf("scheduler.Reschedule(taskA.Name) didn't return an error")
		}

		if err != errors.ErrTaskNotFound {
			t.Errorf("scheduler.Reschedule(taskA.Name) returned err = %v, want %v", err, errors.ErrTaskNotFound)
		}
	})
}
