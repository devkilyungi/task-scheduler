package main

import (
	"fmt"
	"github.com/devkilyungi/time-scheduler/internal/handlers"
	"github.com/devkilyungi/time-scheduler/internal/scheduler"
	"github.com/devkilyungi/time-scheduler/internal/task"
)

func main() {
	sch := scheduler.NewScheduler()

	fmt.Println("\nWelcome to the Task Scheduler!")
	fmt.Println("1. Add a task")
	fmt.Println("2. View tasks")
	fmt.Println("3. Run all tasks")
	fmt.Println("4. Run pending tasks")
	fmt.Println("5. Delete a task")
	fmt.Println("6. Reschedule a task")
	fmt.Println("7. Exit")

	for {
		choice, err := handlers.GetUserChoice()
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		switch choice {
		case 1:
			taskName := handlers.GetTaskName()
			if taskName == "" {
				fmt.Println("Invalid input. Please enter a valid task name.")
				continue
			}

			delayTime, err := handlers.GetDelayTime()
			if err != nil || delayTime < 0 {
				fmt.Println("Invalid input. Please enter a valid delay time.")
				continue
			}

			newTask := task.Task{Name: taskName, Delay: delayTime}
			sch.Add(newTask)
			fmt.Printf("Task %q scheduled with a %d-second delay.\n", newTask.Name, delayTime)
		case 2:
			sch.ViewTasks()
		case 3:
			sch.RunAll()
		case 4:
			sch.RunPending()
		case 5:
			taskName := handlers.GetTaskName()
			err := sch.Delete(taskName)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				fmt.Printf("Task %q deleted.\n", taskName)
			}
		case 6:
			taskName := handlers.GetTaskName()
			newDelay, err := handlers.GetDelayTime()
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid delay time.")
				continue
			}
			err = sch.Reschedule(taskName, newDelay)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				fmt.Printf("Task %q rescheduled to %d seconds.\n", taskName, newDelay)
			}
		case 7:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
