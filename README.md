# Task Scheduler

A simple command-line Task Scheduler application written in Go. This project demonstrates concepts like dependency injection, mocking, and testing while providing a practical utility for scheduling and executing tasks with delays.

---

## Features

- **Add Tasks**: Create a task with a name and execution delay (in seconds).
- **View Tasks**: Display all tasks with their current status (`Pending` or `Completed`).
- **Run All Tasks**: Execute all tasks in the order they were added.
- **Run Pending Tasks**: Execute only tasks with the `Pending` status.
- **Delete Tasks**: Remove a task by its name.
- **Reschedule Tasks**: Update the delay of a task by its name.
- **Testing**: Unit tests and benchmarks ensure reliability.

---

## Getting Started

### Prerequisites
- Go version 1.19 or higher installed on your machine.

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/devkilyungi/task-scheduler.git
   cd task-scheduler
