package errors

type SchedulerError struct {
	Code    string
	Message string
}

func (e *SchedulerError) Error() string {
	return e.Message
}

var (
	ErrTaskNotFound        = &SchedulerError{Code: "TASK_NOT_FOUND", Message: "task not found"}
	ErrTaskFailedToExecute = &SchedulerError{Code: "TASK_EXECUTION_FAILED", Message: "task execution failed"}
)
