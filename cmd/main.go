package main

import (
	"fmt"
	"github.com/devkilyungi/time-scheduler/internal/handlers"
	"github.com/devkilyungi/time-scheduler/internal/scheduler"
	"github.com/devkilyungi/time-scheduler/internal/task"
	"time"
)

func main() {
	sch := scheduler.NewScheduler()

	fmt.Println("Welcome to the Task Scheduler!")
	fmt.Println("\nChoose an option:")
	fmt.Println("1. Add a task")
	fmt.Println("2. View tasks")
	fmt.Println("3. Run all tasks")
	fmt.Println("4. Exit")

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
				fmt.Println("Invalid input. Please enter a valid t name.")
				continue
			}

			delayTime, err := handlers.GetDelayTime()
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid delay time.")
				continue
			}

			if delayTime < 0 {
				fmt.Println("Invalid input. Please enter a valid delay time.")
				continue
			}

			newTask := task.Task{
				Name:  taskName,
				Delay: time.Duration(delayTime),
			}

			sch.Add(newTask)
			fmt.Printf("Task %q scheduled with a %d-second delay.\n", newTask.Name, newTask.Delay)
		case 2:
			fmt.Println("Scheduled tasks:")
			tasks := sch.ViewTasks()
			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
				continue
			}

			for _, t := range tasks {
				fmt.Printf("- %s: %s\n", t.Name, t.Status())
			}
		case 3:
			fmt.Println("Executing tasks:")
			tasks := sch.ViewTasks()
			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
				continue
			}

			sch.RunAll()
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
