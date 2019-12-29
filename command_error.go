package intelpower

import (
	"fmt"
)

const ErrCommand = "command execution error"

type CommandError struct {
	inner error
}

func NewCommandError(err error) CommandError {
	return CommandError{
		inner: err,
	}
}

func (err CommandError) Error() string {
	return fmt.Sprintf("%s: %s", ErrCommand, err.inner.Error())
}

func (err CommandError) Unwrap() error {
	return err.inner
}
