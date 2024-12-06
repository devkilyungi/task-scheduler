package errors

type SchedulerError string

func (e SchedulerError) Error() string {
	return string(e)
}

const (
	ErrTaskFailedToExecute = SchedulerError("task execution failed")
)
