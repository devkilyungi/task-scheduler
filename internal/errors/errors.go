package errors

import "errors"

var (
	ErrTaskFailedToExecute = errors.New("task execution failed")
	ErrTaskNotFound        = errors.New("task not found")
)
